package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {

	sV1 := "Hello I'm pepa pig! Oink Oink G"
	replacer := strings.NewReplacer("G", "Oink")
	sV2 := replacer.Replace(sV1)
	pl(sV2)

	// length
	pl("To find the length:", len(sV2))

	// contains
	pl("To find if a string contains something: ", strings.Contains(sV2, "pig"))
	sV3 := "pepa"

	pl("another example", strings.Contains(sV2, sV3))

	// index
	pl("index of pepa", strings.Index(sV2, "pepa"))

	// replace
	// Replace( string_to_work_on , string_to_replace , string_to_replace_with, number_of_occurences_to_replace -> (-1 = all), (1 = one), etc )
	pl("Replace:", strings.Replace(sV2, "o", "0", -1))

	//  trim spaces
	sV4 := "heyooo heyoo heyoo"

	sV4 = strings.TrimSpace(sV4)

	// split string on the basis of a delimiter
	pl("Split :", strings.Split(sV4, " "))

	//  to lower case and upper case
	pl("Lower Case:", strings.ToLower("ABCDEFG"))
	pl("Lower Case:", strings.ToUpper("abcdefg"))
	pl("Has prefix : ", strings.HasPrefix("yestheory", "yes"))
	pl("Has suffix: ", strings.HasSuffix("yesthoery", "ry"))

}
