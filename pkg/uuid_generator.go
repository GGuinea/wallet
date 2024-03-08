package pkg

import "github.com/google/uuid"

type UUIDGenerator interface {
	Generate() string
}

type UUIDGeneratorImpl struct{}

func NewUUIDGenerator() *UUIDGeneratorImpl {
	return &UUIDGeneratorImpl{}
}

func (u *UUIDGeneratorImpl) Generate() string {
	return uuid.NewString()
}
