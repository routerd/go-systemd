package encoding

// KeyComments implements a container to allow
// attaching comments to keys.
type KeyComments struct {
	comments map[string]string
}

func (c *KeyComments) GetKeyComment(key string) string {
	return c.comments[key]
}

func (c *KeyComments) AddKeyComment(key, comment string) {
	if c.comments == nil {
		c.comments = map[string]string{}
	}
	c.comments[key] = comment
}

func (c *KeyComments) RemoveKeyComment(key string) {
	if c.comments != nil {
		delete(c.comments, key)
	}
}

// SectionList is storing arbitrary sections.
// When embedded into a struct that is given to Unmarshal,
// sections that cannot be assigned to a field are
// appended to this SectionList
type SectionList []Section

func (l *SectionList) AddSection(s Section) {
	*l = append(*l, s)
}

// KeyList is storing arbitrary keys.
// When embedded into a struct that is given to Unmarshal,
// keys that cannot be assigned to a field are
// appended to this KeyList
type KeyList []Key

func (l *KeyList) AddKey(k Key) {
	*l = append(*l, k)
}
