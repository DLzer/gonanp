// Package gonanp is a tiny package for generating North American Numbering Plan compatible phone numbers.
/*
NANP numbers are ten digits in length
The format for NANP are in the format NXX-NXX-XXXX, where N is any digit 2-9 and X is any digit 0-9.
The first three digits are called the numbering plan area (NPA) code, often called simply the area code.
The second three digits are called the central office code or prefix.
The final four digits are called the line number.
*/
package gonanp

import (
	"fmt"
	"math/rand/v2"
	"regexp"
)

// GenerateNanp returns a string containing a valid NANP phone number
func GenerateNanp() string {
	return fmt.Sprintf("%s%s%s", generateAreaCode(), generateOfficeCode(), generateLineNumber())
}

// ValidateNanp accepts a NANP phone number string and returns a bool for validity
func ValidateNanp(n string) bool {
	reg := regexp.MustCompile(`^\(?([2-9][0-8][0-9])\)?[-.●]?([2-9][0-9]{2})[-.●]?([0-9]{4})$`)
	return reg.MatchString(n)
}

// generateAreaCode returns a string containing the NANP area code
func generateAreaCode() string {
	return fmt.Sprintf("%d", areaCodes[rand.IntN(len(areaCodes))])
}

// generateOfficeCode returns a string containing the NANP office code
func generateOfficeCode() string {
	for {
		num := rand.IntN(999-222+1) + 222
		if num < 1000 &&
			!((num%100)/10 == 1 && num%10 == 1) &&
			!(num/100 == 5 && (num%100)/10 == 5 && num%10 == 5) {
			return fmt.Sprintf("%03d", num)
		}
	}
}

// generateLineNumber returns a string containing the NANP line number
func generateLineNumber() string {
	for {
		num := rand.IntN(9999-2222+1) + 2222
		if num < 10000 && !(num/1000 == 2 && (num%1000)/100 == 1) {
			return fmt.Sprintf("%04d", num)
		}
	}
}
