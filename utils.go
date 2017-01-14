package main

func getLines(bytes []byte) []string {
	lines := []string{}
	line := []byte{}
	cnt := len(bytes)
	for i := 0; i < cnt; i++ {
		b := bytes[i]
		line = append(line, b)
		if b == '\r' {
			if i+1 < cnt && bytes[i+1] == '\n' {
				i++
				line = append(line, bytes[i])
			}
			lines = append(lines, string(line))
			line = []byte{}
		}
		if b == '\n' {
			lines = append(lines, string(line))
			line = []byte{}
		}
	}
	return lines
}
