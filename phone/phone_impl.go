package phone

import (
	"errors"
)

type phone struct {
	config PhoneConfig
}

func (p *phone) Parse(phoneNumber string) (PhoneFormater, error) {

	if !p.config.UseMetadata && (p.config.PrefixNumber == "" || p.config.MinimumNumber == 0 || p.config.MaximumNumber == 0) {
		return PhoneFormater{}, errors.New("unable to parse phone. Please specify configuration has been set")
	}

	if p.config.UseMetadata && (p.config.MinimumNumber == 0 || p.config.MaximumNumber == 0) {
		return PhoneFormater{}, errors.New("minimum number and maximum number must be specified")
	}

	formatter := PhoneFormater{}

	if p.config.UseMetadata {
		formatterData, errPhone := getPrefixPhone(phoneNumber)
		if errPhone != nil {
			return PhoneFormater{}, errPhone
		}
		formatter = formatterData
	} else {
		formatter = PhoneFormater{
			PrefixNumber: p.config.PrefixNumber,
			National:     "",
			NationalCode: "",
		}
	}

	sanitize, errSanitize := p.sanitizePhone(formatter, phoneNumber)

	if errSanitize != nil {
		return PhoneFormater{}, errSanitize
	}

	return sanitize, nil

}

func NewPhone(
	config PhoneConfig,
) Phone {
	return &phone{
		config: config,
	}
}
