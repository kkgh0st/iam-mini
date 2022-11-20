package secret

import metav1 "github.com/marmotedu/component-base/pkg/meta/v1"

type Secret struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	metav1.ObjectMeta `       json:"metadata,omitempty"` //
	Username          string                             `json:"username"           gorm:"column:username"  validate:"omitempty"`
	//nolint: tagliatelle
	SecretID  string `json:"secretID"           gorm:"column:secretID"  validate:"omitempty"`
	SecretKey string `json:"secretKey"          gorm:"column:secretKey" validate:"omitempty"`

	// Required: true
	Expires     int64  `json:"expires"     gorm:"column:expires"     validate:"omitempty"`
	Description string `json:"description" gorm:"column:description" validate:"description"`
}
