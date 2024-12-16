package gonanp

import (
	"fmt"

	"math/rand/v2"
)

func GenerateNanp() string {
	return fmt.Sprintf("%s%s%s", generateAreaCode(), generateOfficeCode(), generateLineNumber())
}

func generateAreaCode() string {
	return fmt.Sprintf("%d", areaCodes[rand.IntN(len(areaCodes))])
}

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

func generateLineNumber() string {
	for {
		num := rand.IntN(9999-2222+1) + 2222
		if num < 10000 && !(num/1000 == 2 && (num%1000)/100 == 1) {
			return fmt.Sprintf("%04d", num)
		}
	}
}
