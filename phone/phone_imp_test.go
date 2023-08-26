package phone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fileNameExists = "prefix.json"
)

func TestXxx(t *testing.T) {

	t.Run("error config if not use metada", func(t *testing.T) {
		filenameMetadata = fileNameExists
		config := PhoneConfig{
			UseMetadata:   false,
			PrefixNumber:  "",
			MinimumNumber: 0,
			MaximumNumber: 0,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("6281234567890")

		assert.Error(t, err)
		assert.Equal(t, parse, PhoneFormater{})
	})

	t.Run("error config if use meta data", func(t *testing.T) {
		filenameMetadata = fileNameExists
		config := PhoneConfig{
			UseMetadata:   true,
			PrefixNumber:  "",
			MinimumNumber: 0,
			MaximumNumber: 0,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("6281234567890")

		assert.Error(t, err)
		assert.Equal(t, parse, PhoneFormater{})
	})

	t.Run("error config if use meta data cannot read metadata", func(t *testing.T) {
		filenameMetadata = "fake.json"
		config := PhoneConfig{
			UseMetadata:   true,
			PrefixNumber:  "",
			MinimumNumber: 10,
			MaximumNumber: 12,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("6281234567890")

		assert.Error(t, err)
		assert.Equal(t, parse, PhoneFormater{})
	})

	t.Run("error get prefix from meta data when data not found", func(t *testing.T) {
		filenameMetadata = fileNameExists
		config := PhoneConfig{
			UseMetadata:   true,
			PrefixNumber:  "",
			MinimumNumber: 10,
			MaximumNumber: 12,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("26281234567890")

		assert.Error(t, err)
		assert.Equal(t, parse, PhoneFormater{})
	})

	t.Run("success set prefix but failed sanitize length phone number invalid", func(t *testing.T) {
		filenameMetadata = fileNameExists
		config := PhoneConfig{
			UseMetadata:   false,
			PrefixNumber:  "62",
			MinimumNumber: 10,
			MaximumNumber: 12,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("2")
		parse.PrefixNumber = ""
		assert.Error(t, err)
		assert.Equal(t, parse, PhoneFormater{})
	})

	t.Run("success set prefix but failed sanitize prefix not match", func(t *testing.T) {
		filenameMetadata = fileNameExists
		config := PhoneConfig{
			UseMetadata:   false,
			PrefixNumber:  "62",
			MinimumNumber: 10,
			MaximumNumber: 12,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("89232898323")
		parse.PrefixNumber = ""
		assert.Error(t, err)
		assert.Equal(t, parse, PhoneFormater{})
	})

	t.Run("success set prefix but failed sanitize prefix not int", func(t *testing.T) {
		filenameMetadata = fileNameExists
		config := PhoneConfig{
			UseMetadata:   false,
			PrefixNumber:  "62",
			MinimumNumber: 10,
			MaximumNumber: 12,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("62ABCDASDASDKJH")
		assert.Error(t, err)
		assert.Equal(t, parse, PhoneFormater{})
	})

	t.Run("success parse phone", func(t *testing.T) {
		filenameMetadata = fileNameExists
		config := PhoneConfig{
			UseMetadata:   false,
			PrefixNumber:  "62",
			MinimumNumber: 10,
			MaximumNumber: 12,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("628123232389")
		assert.Nil(t, err)
		assert.NotEqual(t, parse, PhoneFormater{})
	})

	t.Run("success parse phone plus", func(t *testing.T) {
		filenameMetadata = fileNameExists
		config := PhoneConfig{
			UseMetadata:   false,
			PrefixNumber:  "62",
			MinimumNumber: 10,
			MaximumNumber: 12,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("+628123232389")
		assert.Nil(t, err)
		assert.NotEqual(t, parse, PhoneFormater{})
	})

	t.Run("success parse with metadata", func(t *testing.T) {
		filenameMetadata = fileNameExists
		config := PhoneConfig{
			UseMetadata:   true,
			PrefixNumber:  "",
			MinimumNumber: 10,
			MaximumNumber: 12,
		}

		actual := NewPhone(config)
		parse, err := actual.Parse("+628123232389")
		assert.Nil(t, err)
		assert.NotEqual(t, parse, PhoneFormater{})
	})

}
