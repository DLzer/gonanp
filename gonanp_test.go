package gonanp_test

import (
	"regexp"
	"testing"

	"github.com/DLzer/gonanp"
	"github.com/smallstep/assert"
)

func TestGenerateNanp(t *testing.T) {
	t.Parallel()
	// Regex Asserts NANP
	// Where NANP numbers are ten digits in length
	// and they are in the format NXX-NXX-XXXX
	// where N is any digit 2-9 and X is any digit 0-9
	reg := regexp.MustCompile(`^\(?([2-9][0-8][0-9])\)?[-.●]?([2-9][0-9]{2})[-.●]?([0-9]{4})$`)

	for i := 0; i < 100; i++ {
		assert.True(t, reg.MatchString(gonanp.GenerateNanp()))
	}

	invalidNanpNumber := "0000820123"
	assert.False(t, reg.MatchString(invalidNanpNumber))
}

func TestValidateNanp(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		nanp := gonanp.GenerateNanp()
		assert.True(t, gonanp.ValidateNanp(nanp))
	}

	invalidNanpNumber := "0144797077"
	assert.False(t, gonanp.ValidateNanp(invalidNanpNumber))
}
