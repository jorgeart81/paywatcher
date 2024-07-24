package services

type HashService interface {
	Has(password string) ([]byte, error)
	Compare(hashedPassword, password string) error
}
