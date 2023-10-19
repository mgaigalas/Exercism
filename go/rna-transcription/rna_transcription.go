// Package strand contains tools for converting DNA to RNA
package strand

// ToRNA converts DNA string to RNA string
func ToRNA(dna string) string {
	res := make([]rune, len(dna))
	for idx, val := range dna {
		res[idx] = toRnaNucleotide(val)
	}
	return string(res)
}

func toRnaNucleotide(r rune) rune {
	switch r {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		panic(string(r) + " nucleotide does not exist")
	}
}
