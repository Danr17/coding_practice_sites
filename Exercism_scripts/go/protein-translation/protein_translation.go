package protein

import "errors"

var (
	ErrStop        = errors.New("STOP word has been found")
	ErrInvalidBase = errors.New("the base is invalid")
)

//FromCodon maps condons to protein name
func FromCodon(input string) (string, error) {

	condons := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	if val, ok := condons[input]; ok {
		if val == "STOP" {
			return "", ErrStop
		}
		return val, nil
	}
	return "", ErrInvalidBase
}

//FromRNA returns the proteins
func FromRNA(input string) ([]string, error) {

	var proteins []string

	for s := 0; s < len(input); s += 3 {
		protein, err := FromCodon(input[s : s+3])
		if err != nil {
			if err == ErrStop {
				return proteins, nil
			}
			return proteins, err
		}

		proteins = append(proteins, protein)

	}

	return proteins, nil
}
