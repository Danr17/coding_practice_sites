//Package bob remarks
package bob

import (
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

// Hey responds to remarks
func Hey(remark string) string {

	remark = strings.Trim(remark, " \t\n\r")

	uRemark := strings.ToUpper(remark)

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case !strings.ContainsAny(uRemark, "ABCDEFGHIJKLMJOPQRSTUVWXYZ?"):
		return "Whatever."
	case (strings.Compare(remark, uRemark) == 0) && (strings.HasSuffix(remark, "?") && strings.ContainsAny(uRemark, "ABCDEFGHIJKLMJOPQRSTUVWXYZ")):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case strings.Compare(remark, uRemark) == 0:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
	return ""
}
