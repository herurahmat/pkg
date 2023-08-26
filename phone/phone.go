package phone

type Phone interface {
	Parse(phone string) (PhoneFormater, error)
}
