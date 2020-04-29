def proteins(strand):
    condons = {
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

    RNAs = [condons[strand[i:i+3]] for i in range(0, len(strand), 3)]

    validRNAs = []
    for s in RNAs:
        if s == "" or s == "STOP":
            return validRNAs
        validRNAs.append(s)
    return validRNAs
 

    