package domain

type Service interface {
	Init(userID int64, credentials map[string]string) error
	Finalize()

	CredentialsKeys() []string
	Ident() string
}
