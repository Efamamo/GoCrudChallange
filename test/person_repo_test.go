package repo_test

import (
	"testing"

	"github.com/Efamamo/GoCrudChallange/application/people/command"
	ierr "github.com/Efamamo/GoCrudChallange/domain/error"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/Efamamo/GoCrudChallange/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// PersonCommandTestSuite is the base test suite for person command handlers.
type PersonCommandTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockPersonRepo
}

// SetupTest initializes the mock repository before each test.
func (suite *PersonCommandTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewMockPersonRepo()
}

// TearDownTest cleans up after each test.
func (suite *PersonCommandTestSuite) TearDownTest() {
	suite.mockRepo = nil
}

// TestCreatePersonHandler_Success tests the successful creation of a person.
func (suite *PersonCommandTestSuite) TestCreatePersonHandler_Success() {
	handler := command.NewCreatePersonHandler(suite.mockRepo)

	cmd := &command.CreatePersonCommand{
		Name:    "John Doe",
		Age:     30,
		Hobbies: []string{"Reading", "Running"},
	}

	result, err := handler.Handle(cmd)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), cmd.Name, result.Name())
	assert.Equal(suite.T(), cmd.Age, result.Age())
	assert.Equal(suite.T(), cmd.Hobbies, result.Hobbies())
}

// TestCreatePersonHandler_Failure_InvalidPerson tests the failure case for creating a person with invalid data.
func (suite *PersonCommandTestSuite) TestCreatePersonHandler_Failure_InvalidPerson() {
	handler := command.NewCreatePersonHandler(suite.mockRepo)

	// Custom behavior to return an error for invalid data
	suite.mockRepo.SaveFunc = func(p *model.Person) ierr.IErr {
		return ierr.NewValidation("person can't be empty")
	}

	cmd := &command.CreatePersonCommand{
		Name:    "",
		Age:     0,
		Hobbies: nil,
	}

	result, err := handler.Handle(cmd)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

// TestUpdatePersonHandler_Success tests the successful update of a person's details.
func (suite *PersonCommandTestSuite) TestUpdatePersonHandler_Success() {
	handler := command.NewUpdatePersonHandler(suite.mockRepo)

	existingPerson, err := model.CreatePerson(
		&model.PersonConfig{
			Name:    "Existing User",
			Age:     25,
			Hobbies: []string{"Swimming"},
		},
	)
	suite.mockRepo.Save(existingPerson)

	cmd := &command.UpdatePersonCommand{
		ID:      existingPerson.Id(),
		Name:    "Updated Name",
		Age:     30,
		Hobbies: []string{"Updated Hobby"},
	}

	result, err := handler.Handle(cmd)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), cmd.Name, result.Name())
	assert.Equal(suite.T(), cmd.Age, result.Age())
	assert.Equal(suite.T(), cmd.Hobbies, result.Hobbies())
}

// TestUpdatePersonHandler_Failure_NotFound tests the failure case for updating a non-existent person.
func (suite *PersonCommandTestSuite) TestUpdatePersonHandler_Failure_NotFound() {
	handler := command.NewUpdatePersonHandler(suite.mockRepo)

	cmd := &command.UpdatePersonCommand{
		ID:      uuid.New(),
		Name:    "Non-Existent User",
		Age:     30,
		Hobbies: []string{"Hiking"},
	}

	result, err := handler.Handle(cmd)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

// TestDeletePersonHandler_Success tests the successful deletion of a person by ID.
func (suite *PersonCommandTestSuite) TestDeletePersonHandler_Success() {
	handler := command.NewDeletePersonHandler(suite.mockRepo)

	existingPerson, err := model.CreatePerson(
		&model.PersonConfig{
			Name:    "User To Delete",
			Age:     25,
			Hobbies: []string{"Climbing"},
		},
	)

	suite.mockRepo.Save(existingPerson)

	success, err := handler.Handle(existingPerson.Id())
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), success)
}

// TestDeletePersonHandler_Failure_NotFound tests the failure case for deleting a non-existent person.
func (suite *PersonCommandTestSuite) TestDeletePersonHandler_Failure_NotFound() {
	handler := command.NewDeletePersonHandler(suite.mockRepo)

	nonExistentID := uuid.New()

	success, err := handler.Handle(nonExistentID)
	assert.Error(suite.T(), err)
	assert.False(suite.T(), success)
}

// TestPersonCommandTestSuite runs the test suite for person command handlers.
func TestPersonCommandTestSuite(t *testing.T) {
	suite.Run(t, new(PersonCommandTestSuite))
}
