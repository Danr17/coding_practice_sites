package secret

var actions map[int]string = map[int]string{
	0: "wink",
	1: "double blink",
	2: "close your eyes",
	3: "jump",
}

//Handshake given a decimal number, convert it to the appropriate sequence of events for a secret handshake.
func Handshake(code uint) []string {

	result := []string{}

	if code&(1<<0) != 0 {
		result = append(result, "wink")
	}
	if code&(1<<1) != 0 {
		result = append(result, "double blink")
	}
	if code&(1<<2) != 0 {
		result = append(result, "close your eyes")
	}
	if code&(1<<3) != 0 {
		result = append(result, "jump")
	}
	if code&(1<<4) != 0 {
		result = reverse(result)
	}
	return result
}

func reverse(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
