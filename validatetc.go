package validatetc

import (
	"strconv"
)

func Validate(tcnumber string) bool {

	runes := []rune(tcnumber)

	odd, even, sum, rebuild := 0, 0, 0, ""

	for i := 0; i < len(runes)-2; i++ {

		a, _ := strconv.Atoi(string(runes[i]))

		if string(runes[0]) == "0" {
			return false
		}

		if (i+1)%2 == 0 {
			odd += a
		} else {
			even += a
		}

		rebuild += string(runes[i])

		sum += a
	}

	ten := (even*7 - odd) % 10

	indexTen, _ := strconv.Atoi(string(runes[9]))

	eleven := (sum + indexTen) % 10

	build := string(rebuild) + strconv.Itoa(ten) + strconv.Itoa(eleven)

	if build == tcnumber {
		return true
	}

	return false
}
