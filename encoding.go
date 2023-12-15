package main

func MsdosEncode(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		ch := s[i]
		
		if ch == '\n' {
			result += "\r\n"
		} else if ch != '\r' {
			result += string(ch)
		}
	}

	return result
}
