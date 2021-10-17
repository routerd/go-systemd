package systemd

import (
	"routerd.net/go-systemd/internal/encoding"
)

type SectionList = encoding.SectionList

type KeyList = encoding.KeyList

type KeyComments = encoding.KeyComments

type InvalidUnmarshalError = encoding.InvalidUnmarshalError

var (
	Marshal   = encoding.Marshal
	Unmarshal = encoding.Unmarshal
)

// Generic stuff

type File = encoding.File

type Section = encoding.Section

type Key = encoding.Key

var (
	Decode = encoding.Decode
	Encode = encoding.Encode
)

// utils

var (
	StringPtr = encoding.StringPtr
	BoolPtr   = encoding.BoolPtr
)
