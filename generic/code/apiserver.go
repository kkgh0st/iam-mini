package code

const (
	// ErrUserNotFound - 404: User not found.
	ErrUserNotFound int = iota + 110001

	// ErrUserAlreadyExist - 400: User already exist.
	ErrUserAlreadyExist
)

// iam-apiserver: policy errors.
const (
	// ErrPolicyNotFound - 404: Policy not found.
	ErrPolicyNotFound int = iota + 110201
)

const (
	// ErrEncrypt - 400: Secret reach the max count.
	ErrReachMaxCount int = iota + 110101

	//  ErrSecretNotFound - 404: Secret not found.
	ErrSecretNotFound
)
