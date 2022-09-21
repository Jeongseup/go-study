package user

import (
	"golang-unit-tests/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryStub struct {
	mock.Mock
}

func (r *UserRepositoryStub) Add(user entity.User) error {
	args := r.Called(user)

	return args.Error(0)
}

type BadWordsRepositoryStub struct {
	mock.Mock
}

func (r *BadWordsRepositoryStub) FindAll() ([]string, error) {
	args := r.Called()

	return args.Get(0).([]string), args.Error(1)
}

func TestShouldSuccessfullyRegister(t *testing.T) {

	user := entity.User{
		Name:        "Jeongseup",
		Email:       "jeongseup@example.com",
		Description: "Software Developer",
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)

	BadWordsRepository := &BadWordsRepositoryStub{}
	BadWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)

	userService := NewUserService(userRepository, BadWordsRepository)

	err := userService.Register(user)
	userRepository.AssertCalled(t, "Add", user)
	assert.Nil(t, err)
}

func TestShouldNotRegisterateTheUserWhenBadWordIsFound(t *testing.T) {

	user := entity.User{
		Name:        "Jeongseup",
		Email:       "jeongseup@example.com",
		Description: "Software potato Developer",
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)

	BadWordsRepository := &BadWordsRepositoryStub{}
	BadWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)

	userService := NewUserService(userRepository, BadWordsRepository)

	err := userService.Register(user)
	userRepository.AssertNotCalled(t, "Add", user)
	assert.Error(t, err)
}
