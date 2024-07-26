package services

type HashService interface {
	Has(password string) (string, error)
	Compare(hashedPassword, password string) error
}
