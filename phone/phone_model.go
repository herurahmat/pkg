package phone

var filenameMetadata = "prefix.json"

type PhoneConfig struct {
	UseMetadata   bool
	PrefixNumber  string
	MinimumNumber int
	MaximumNumber int
}

type PhoneFormater struct {
	OriginalPhoneNumber string
	PrefixNumber        string
	Number              uint64
	National            string
	NationalCode        string
}
