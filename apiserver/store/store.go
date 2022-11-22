package store

type Factory interface {
	Secrets() SecretStore
	Users() UserStore
	Policies() PolicyStore
}

var client Factory

func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}
