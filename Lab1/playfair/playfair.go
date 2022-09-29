package playfair

func Encrypt() {

}

func FindInTable(t [5][5]int, f int) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {

			if t[i][j] == f {
				return i, j
			}
		}
	}
	return 0, 0
}
