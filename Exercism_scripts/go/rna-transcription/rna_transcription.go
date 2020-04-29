package strand

//ToRNA transform DNA to RNA
func ToRNA(dna string) string {
	DNAtoRNA := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	rna := ""
	for _, s := range dna {
		rna = rna + DNAtoRNA[string(s)]
	}
	return rna
}
