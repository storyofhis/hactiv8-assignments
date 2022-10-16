package service

import (
	"context"
	"fmt"
	"latihan-testing/mocks"
	"latihan-testing/models/entity"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// add two cases here
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := &entity.User{
			Email:    "testone@mail.com",
			Username: "testone",
			Name:     "testone",
			Age:      18,
		}

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewService(&USConfig{
			UserRepository: mockUserRepository,
		})
		mockUserRepository.On("FindByID", mock.Anything, uid).Return(mockUserResp, nil)

		ctx := context.TODO()
		u, err := us.Get(ctx, uid)

		assert.NoError(t, err)
		assert.Equal(t, u, mockUserResp)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewService(&USConfig{
			UserRepository: mockUserRepository,
		})

		mockUserRepository.On("FindByID", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down the call chain"))

		ctx := context.TODO()
		u, err := us.Get(ctx, uid)

		assert.Nil(t, u)
		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}
