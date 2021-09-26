package strand

var complements = map[string]string{
	"C": "G",
	"G": "C",
	"T": "A",
	"A": "U",
}

func ToRNA(dna string) string {
	rna := ""
	for _, char := range dna {
		rna += complements[string([]rune{char})]
	}
	return rna
}
