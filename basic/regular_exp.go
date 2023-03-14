package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {

	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr)

	pl(match)
	// ----------------------------------------------------------------------------------------------------------------------------

	reStr2 := "The Fat Rat, cat mat bat pat"
	r, _ := regexp.Compile("([cmbp]at)")
	pl("Match String : ", r.MatchString(reStr2))
	pl("Find String : ", r.FindString(reStr2))
	pl("Index : ", r.FindStringIndex(reStr2))
	pl("All string : ", r.FindAllString(reStr2, -1))
	pl("First 2 string : ", r.FindAllString(reStr2, 2))
	pl("All submatch index : ", r.FindAllStringSubmatchIndex(reStr, -1))
	pl("ReplaceAllString : ", r.ReplaceAllString(reStr, "Dog"))

}
