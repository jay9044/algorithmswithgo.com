package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	res := ""
	// starting at last word - eg. "cat" index is actually two but length is 3
	for i := len(word) - 1; i >= 0; i-- {
		res += string(word[i])
	}
	return res
}

// word := "jamal"
// 	var res string

// 	for _, r := range word {
// 		res = string(r) + res
// 	}
// 	fmt.Println(res)
