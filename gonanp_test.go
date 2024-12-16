package gonanp_test

import (
	"regexp"
	"testing"

	"github.com/DLzer/gonanp"
	"github.com/smallstep/assert"
)

func TestGenerateNanp(t *testing.T) {
	reg := regexp.MustCompile(`^\(?([2-9][0-8][0-9])\)?[-.●]?([2-9][0-9]{2})[-.●]?([0-9]{4})$`)

	for i := 0; i < 100; i++ {
		assert.True(t, reg.MatchString(gonanp.GenerateNanp()))
	}
}
