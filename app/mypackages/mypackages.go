package stuff

import (
	"errors"
	"strconv"
)

var Name string = "Packages are working fine"

func IntArrToStr(arr []int) []string {
	var strArr []string
	for _, num := range arr {
		strArr = append(strArr, strconv.Itoa(num))
	}
	return strArr
}

// ----------------------------------------------------------------------------------------------------------------------------
//  protecting data and using getter and setter methods

type Date struct {
	day   int
	month int
	year  int
}

// notice that the date variables are in lowercase and hence inaccesible outside this go file
// whereas these functions are in uppercase and will be accesible for call outside this go file

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid date range")
	}
	d.day = day
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month range")
	}
	d.month = month
	return nil
}
func (d *Date) SetYear(year int) error {
	if year < 0 || year > 2023 {
		return errors.New("invalid year range")
	}
	d.year = year
	return nil
}

func (d *Date) GetDay() int {
	return d.day
}
func (d *Date) GetMonth() int {
	return d.month
}
func (d *Date) GetYear() int {
	return d.year
}

// ----------------------------------------------------------------------------------------------------------------------------
