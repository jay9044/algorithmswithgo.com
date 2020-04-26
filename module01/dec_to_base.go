package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	charset:= "0123456789ABCDEF"
	var binary string
	for dec > 0 {
		rem:= dec % base
		binary = string(charset[rem]) + binary
		dec = dec / base
	}
	return binary
}
