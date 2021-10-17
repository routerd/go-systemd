package encoding

// Returns a pointer to the given string.
func StringPtr(str string) *string {
	return &str
}

// Returns a pointer to the given bool.
func BoolPtr(b bool) *bool {
	return &b
}

func StrToBool(b string) *bool {
	switch b {
	case "1", "yes", "true", "on":
		return BoolPtr(true)

	case "0", "no", "false", "off":
		return BoolPtr(false)
	}
	return nil
}

func BoolToStr(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func filterEmpty(in []string) (out []string) {
	for _, s := range in {
		if len(s) == 0 {
			continue
		}
		out = append(out, s)
	}
	return
}
