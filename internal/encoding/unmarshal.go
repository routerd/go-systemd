package encoding

import (
	"reflect"
	"strings"
)

// Unmarshal parses the systemd unit data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	file, err := Decode(data)
	if err != nil {
		return err
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	return unmarshalSections(file, rv)
}

func unmarshalSections(file *File, rv reflect.Value) error {
	// must be a pointer
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{rv.Type()}
	}

	knownSections := map[string]struct{}{}
	tv := rv.Elem().Type()
	for i := 0; i < rv.Elem().NumField(); i++ {
		field := rv.Elem().Field(i)
		fieldConfig := configForField(tv.Field(i))
		knownSections[fieldConfig.Name] = struct{}{}
		sections := file.SectionsByName(fieldConfig.Name)
		if len(sections) == 0 {
			// no section with this name
			continue
		}

		for _, section := range sections {
			var newObject reflect.Value
			switch field.Type().Kind() {
			case reflect.Struct:
				newObject = reflect.New(field.Type()).Elem()

			case reflect.Ptr, reflect.Slice:
				newObject = reflect.New(field.Type().Elem())
			}

			newObjectPtr := newObject
			if newObject.Kind() != reflect.Ptr {
				newObjectPtr = newObject.Addr()
			}

			if err := unmarshalKeys(&section, newObjectPtr); err != nil {
				return err
			}
			commentField := newObjectPtr.Elem().FieldByName("Comment")
			if commentField.IsValid() {
				commentField.Set(reflect.ValueOf(section.Comment))
			}

			if field.Type().Kind() != reflect.Slice {
				field.Set(newObject)
				continue
			}
			field.Set(reflect.Append(field, newObjectPtr.Elem()))
		}
	}

	// Add Sections that don't fit into any other place
	// if there is a AddSection function implemented.
	addSection := rv.MethodByName("AddSection")
	if !addSection.IsValid() {
		return nil
	}
	for _, section := range file.Sections {
		if _, ok := knownSections[section.Name]; ok {
			continue
		}
		addSection.Call([]reflect.Value{reflect.ValueOf(section)})
	}
	return nil
}

func unmarshalKeys(section *Section, rv reflect.Value) error {
	// must be a pointer
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{rv.Type()}
	}

	knownKeys := map[string]struct{}{}
	tv := rv.Elem().Type()
	for i := 0; i < rv.Elem().NumField(); i++ {
		field := rv.Elem().Field(i)

		fieldConfig := configForField(tv.Field(i))
		knownKeys[fieldConfig.Name] = struct{}{}
		keys := section.KeysByName(fieldConfig.Name)
		if len(keys) == 0 {
			// no key with this name
			continue
		}

		var comment string
		switch field.Type().Kind() {
		case reflect.String:
			key := keys[len(keys)-1]
			field.SetString(key.Value)
			comment = key.Comment

		case reflect.Ptr:
			switch field.Type().Elem().Kind() {
			case reflect.Bool:
				key := keys[len(keys)-1]
				if key.Value == "" && fieldConfig.Omitempty {
					// skip empty
					continue
				}

				b := StrToBool(key.Value)
				if b == nil {
					// TODO: warning?
					continue
				}

				boolRV := reflect.New(field.Type().Elem())
				boolRV.Elem().SetBool(*b)
				field.Set(boolRV)

			case reflect.String:
				key := keys[len(keys)-1]
				if key.Value == "" && fieldConfig.Omitempty {
					// skip empty
					continue
				}

				stringRV := reflect.New(field.Type().Elem())
				stringRV.Elem().SetString(key.Value)
				field.Set(stringRV)

			default:
				continue
			}

		case reflect.Slice:
			if field.Type().Elem().Kind() != reflect.String {
				// wrong key type
				continue
			}

			var values []string
			for _, key := range keys {
				if comment != "" {
					comment += "\n"
				}
				comment += key.Comment

				if len(key.Value) == 0 {
					// A key with no value reset's all previously read values.
					values = nil
					// no need to save key comments that don't apply
					comment = ""
					continue
				}

				if fieldConfig.WSlist {
					values = append(values,
						filterEmpty(strings.Split(key.Value, " "))...)
					continue
				}
				values = append(values, key.Value)
			}
			field.Set(reflect.ValueOf(values))
		}

		// comment handling
		if comment == "" {
			continue
		}
		addComment := rv.MethodByName("AddKeyComment")
		if !addComment.IsValid() {
			continue
		}
		addComment.Call([]reflect.Value{
			reflect.ValueOf(fieldConfig.Name),
			reflect.ValueOf(comment),
		})
	}

	// Add Keys that don't fit into any other place
	// if there is a AddKey function implemented.
	addKey := rv.MethodByName("AddKey")
	if !addKey.IsValid() {
		return nil
	}
	for _, key := range section.Keys {
		if _, ok := knownKeys[key.Name]; ok {
			continue
		}
		addKey.Call([]reflect.Value{reflect.ValueOf(key)})
	}
	return nil
}

const fieldTagName = "systemd"

type fieldConfig struct {
	Name      string
	Omitempty bool
	// white space list
	WSlist bool
}

func configForField(structField reflect.StructField) (c fieldConfig) {
	c.Name = structField.Name

	if tag := structField.Tag.Get(fieldTagName); tag != "" {
		idx := strings.Index(tag, ",")
		if idx == -1 {
			c.Name = tag
			return
		}

		name := tag[:idx]
		if name != "" {
			c.Name = tag[:idx]
		}
		c.Omitempty = strings.Contains(tag[idx:], "omitempty")
		c.WSlist = strings.Contains(tag[idx:], "wslist")
	}
	return
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "json: Unmarshal(nil " + e.Type.String() + ")"
}
