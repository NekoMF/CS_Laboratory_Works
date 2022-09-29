package playfair

import "strings"

func Encrypt(key string, text string) string {
	text = cleanText(text)
	table := generateTable(key)

	var result strings.Builder
	for i := 0; i < len(text); i += 2 {
		c1 := text[i]
		c2 := text[i+1]

		i1, j1 := findInTable(table, c1)
		i2, j2 := findInTable(table, c2)

		if i1 == i2 {
			result.WriteByte(table[i1][(j1+1)%5])
			result.WriteByte(table[i2][(j2+1)%5])
		} else if j1 == j2 {
			result.WriteByte(table[(i1+1)%5][j1])
			result.WriteByte(table[(i2+1)%5][j2])
		} else {
			result.WriteByte(table[i1][j2])
			result.WriteByte(table[i2][j1])
		}
	}

	return result.String()
}

func Decrypt(key string, text string) string {
	table := generateTable(key)

	var result strings.Builder
	for i := 0; i < len(text); i += 2 {
		c1 := text[i]
		c2 := text[i+1]

		i1, j1 := findInTable(table, c1)
		i2, j2 := findInTable(table, c2)

		if i1 == i2 {
			result.WriteByte(table[i1][(j1-1+5)%5])
			result.WriteByte(table[i2][(j2-1+5)%5])
		} else if j1 == j2 {
			result.WriteByte(table[(i1-1+5)%5][j1])
			result.WriteByte(table[(i2-1+5)%5][j2])
		} else {
			result.WriteByte(table[i1][j2])
			result.WriteByte(table[i2][j1])
		}
	}

	return result.String()
}

func cleanText(text string) string {
	text = strings.ToLower(text)
	var cleanText strings.Builder
	prevByte := byte('\000')

	for i := 0; i < len(text); i++ {
		nextByte := text[i]
		if nextByte < 'a' || nextByte > 'z' {
			continue
		}
		if nextByte == 'j' {
			nextByte = 'i'
		}
		if nextByte != prevByte {
			cleanText.WriteByte(nextByte)
		} else {
			cleanText.WriteByte('x')
			cleanText.WriteByte(nextByte)
		}
		prevByte = nextByte
	}

	l := cleanText.Len()
	if l%2 == 1 {
		if cleanText.String()[l-1] != 'x' {
			cleanText.WriteByte('x')
		} else {
			cleanText.WriteByte('z')
		}
	}
	return cleanText.String()
}

func findInTable(table [5][5]byte, c byte) (int, int) {
	for i := range table {
		for j := range table[i] {
			if table[i][j] == c {
				return i, j
			}
		}
	}
	return -1, -1
}

func generateTable(key string) [5][5]byte {
	var used [26]bool
	used[9] = true
	var table [5][5]byte
	alphabet := strings.ToLower(key) + "abcdefghiklmnopqrstuvwxyz"
	for i, j, k := 0, 0, 0; k < len(alphabet); k++ {
		c := alphabet[k]
		if c < 'a' || c > 'z' {
			continue
		}
		d := int(c - 97)
		if !used[d] {
			table[i][j] = c
			used[d] = true
			j++
			if j == 5 {
				i++
				if i == 5 {
					break
				}
				j = 0
			}
		}
	}

	return table
}
