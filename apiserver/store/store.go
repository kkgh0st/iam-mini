package store

type Factory interface {
	Secrets() SecretStore
}
