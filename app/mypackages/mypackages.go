package stuff

import "strconv"

var Name string = "Packages are working fine"

func IntArrToStr(arr []int) []string {
	var strArr []string
	for _, num := range arr {
		strArr = append(strArr, strconv.Itoa(num))
	}
	return strArr
}
