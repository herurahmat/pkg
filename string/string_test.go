package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	assert.Equal(t, "⇨ HTTP server started on \033[32m[::]:8080\033[0m\n", Server("HTTP", "8080"))
}

func TestSlug(t *testing.T) {
	t.Run("success generate", func(t *testing.T) {
		slug, err := Slug("?This is a test-title with question mark ? and& \" exclaimation mark !")
		assert.Nil(t, err)
		assert.Equal(t, slug, "this-is-a-test-title-with-question-mark-and-exclaimation-mark")
	})
}

func TestContains(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		slug := Contains("HUAWEI XXX", []string{"HUAWEI", "ADVAN"})
		assert.Equal(t, true, slug)
	})
}

func TestLowerAlphaUnderscore(t *testing.T) {
	assert.Equal(t, "test_objects", LowerAlphaUnderscore("test_objects"))
	assert.Equal(t, "test_objects", LowerAlphaUnderscore("123test_-!@#objects"))
	assert.Equal(t, "test_objects", LowerAlphaUnderscore("TEST_OBJECTS"))
}

func BenchmarkLowerAlphaUnderscore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LowerAlphaUnderscore(strings.Repeat("abcdefghijklmnopqrstuvwxyz_", 10000))
	}
}
