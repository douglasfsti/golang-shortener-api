package shortner

import (
	Serializer "github.com/douglasfsti/golang-shortener-api/internal/serializer"
	Validation "github.com/go-ozzo/ozzo-validation/v3"
	Is "github.com/go-ozzo/ozzo-validation/v3/is"
	"github.com/sony/sonyflake"
	"io"
	"time"
)

type UseCases interface {
	NewRedirect(url string) (*Redirect, error)
	NewRedirectIoReader(data io.Reader) (*Redirect, error)
	Encode(redirect *Redirect) ([]byte, error)
	Validate(redirect *Redirect) error
}

func NewUseCases(codeGenerator *sonyflake.Sonyflake, serializer Serializer.Serializer) UseCases {
	return &usecases{
		CodeGenerator: codeGenerator,
		Serializer:    serializer,
	}
}

type usecases struct {
	CodeGenerator *sonyflake.Sonyflake
	Serializer    Serializer.Serializer
}

func (u *usecases) NewRedirect(url string) (*Redirect, error) {
	id, err := u.CodeGenerator.NextID()
	if err != nil {
		return nil, err
	}

	redirect := &Redirect{
		Code:      id,
		URL:       url,
		CreatedAt: time.Now().UTC(),
	}

	return redirect, u.Validate(redirect)
}

func (u *usecases) NewRedirectIoReader(data io.Reader) (*Redirect, error) {
	var redirect Redirect
	if err := u.Serializer.DecodeIoReader(data, &redirect); err != nil {
		return nil, err
	}

	return u.NewRedirect(redirect.URL)
}

func (u *usecases) Encode(redirect *Redirect) ([]byte, error) {
	return u.Serializer.Encode(redirect)
}

func (u *usecases) Validate(redirect *Redirect) error {
	return Validation.ValidateStruct(redirect, []*Validation.FieldRules{
		Validation.Field(&redirect.URL, Is.URL),
	}...)
}
