package encoding

import (
	"io"
	"strings"
)

// Encode takes the runtime representsation of a systemd configuration file and writes out a normal systemd file.
func Encode(out io.Writer, file *File) error {
	return writeFile(out, file)
}

func writeFile(out io.Writer, file *File) error {
	for i, section := range file.Sections {
		if i != 0 {
			if _, err := out.Write([]byte("\n")); err != nil {
				return err
			}
		}
		if err := writeSection(out, &section); err != nil {
			return err
		}
	}
	return nil
}

func writeSection(out io.Writer, section *Section) error {
	if err := writeComment(out, section.Comment); err != nil {
		return err
	}
	if _, err := out.Write([]byte("[" + section.Name + "]\n")); err != nil {
		return err
	}
	for i, key := range section.Keys {
		if i != 0 && key.Comment != "" {
			if _, err := out.Write([]byte("\n")); err != nil {
				return err
			}
		}
		if err := writeKey(out, &key); err != nil {
			return err
		}
	}
	return nil
}

func writeKey(out io.Writer, key *Key) error {
	if err := writeComment(out, key.Comment); err != nil {
		return err
	}
	if _, err := out.Write([]byte(key.Name + "=" + key.Value + "\n")); err != nil {
		return err
	}
	return nil
}

func writeComment(out io.Writer, comment string) error {
	if comment == "" {
		return nil
	}
	commentBlock := ""
	for _, line := range strings.Split(comment, "\n") {
		commentBlock += "# " + line + "\n"
	}
	_, err := out.Write([]byte(commentBlock))
	if err != nil {
		return err
	}
	return nil
}
