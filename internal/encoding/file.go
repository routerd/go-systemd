package encoding

type File struct {
	Sections []Section
}

func (f *File) SectionsByName(name string) (out []Section) {
	for _, section := range f.Sections {
		if section.Name != name {
			continue
		}

		out = append(out, section)
	}
	return
}

type Section struct {
	Name    string
	Comment string
	Keys    []Key
}

func (s *Section) KeysByName(name string) (out []Key) {
	for _, key := range s.Keys {
		if key.Name != name {
			continue
		}

		out = append(out, key)
	}
	return
}

type Key struct {
	Name    string
	Value   string
	Comment string
}
