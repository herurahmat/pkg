package phone

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type phoneJsons []phoneJson
type phoneJson struct {
	Name     string `json:"name"`
	DialCode string `json:"dial_code"`
	Code     string `json:"code"`
}

func getPrefixPhone(phone string) (PhoneFormater, error) {

	phoneNumber := removePhonePlus(phone)

	twoDigit := phoneNumber[:2]
	threeDigit := phoneNumber[:3]

	file, errFile := os.ReadFile(filenameMetadata)

	if errFile != nil {
		return PhoneFormater{}, errFile
	}

	data := phoneJsons{}
	_ = json.Unmarshal([]byte(file), &data)

	for _, v := range data {
		if v.DialCode == twoDigit || v.Code == threeDigit {
			return PhoneFormater{
				OriginalPhoneNumber: phone,
				PrefixNumber:        v.DialCode,
				National:            v.Name,
				NationalCode:        v.Code,
			}, nil
		}
	}

	return PhoneFormater{}, errors.New("data not found")

}

func (p *phone) sanitizePhone(phoneFormatter PhoneFormater, phoneNumber string) (PhoneFormater, error) {

	lengthPrefix := len(phoneFormatter.PrefixNumber)

	phoneNumber = removePhonePlus(phoneNumber)

	if len(phoneNumber) < lengthPrefix {
		return PhoneFormater{}, errors.New("phone number must greater than prefix number")
	}

	if phoneNumber[:lengthPrefix] != phoneFormatter.PrefixNumber {
		return PhoneFormater{}, errors.New("prefix number not equal to phone number")
	}

	phoneNumberWithOutPrefix := phoneNumber[lengthPrefix:]

	number, err := strconv.ParseUint(phoneNumberWithOutPrefix, 10, 64)

	if err != nil {
		return PhoneFormater{}, err
	}

	phoneFormatter.Number = number

	return phoneFormatter, nil
}

func removePhonePlus(phoneNumber string) string {
	plusCharacter := phoneNumber[:1]

	if plusCharacter == "+" {
		phoneNumber = phoneNumber[1:]
	}

	return phoneNumber
}
