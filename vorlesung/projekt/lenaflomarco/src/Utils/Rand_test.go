package Utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

//TestCorrectLength tests if generated random string has the correct length
func TestCorrectLength(t *testing.T) {
	for i := 1; i < 10; i++ {
		salt := RandString(i)
		assert.Equal(t, i, strings.Count(salt, "") - 1, "Salt is not Generating " + string(i) + " Characters")
	}
}

