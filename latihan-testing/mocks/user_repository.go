package mocks

import (
	"context"
	"latihan-testing/models/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*entity.User, error) {
	ret := m.Called(ctx, uid)

	var r0 *entity.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entity.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
