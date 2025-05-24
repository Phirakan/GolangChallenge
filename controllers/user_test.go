package controllers

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"

	"golang-test/models"
)

// Simple Mock Database - เก็บข้อมูลใน memory
type MockDB struct {
	users map[string]*models.User
}

func NewMockDB() *MockDB {
	return &MockDB{
		users: make(map[string]*models.User),
	}
}

func (m *MockDB) CreateUser(user *models.User) error {
	for _, existingUser := range m.users {
		if existingUser.Email == user.Email {
			return errors.New("duplicate email")
		}
	}
	
	user.ID = bson.NewObjectID()
	user.CreatedAt = time.Now()
	m.users[user.ID.Hex()] = user
	return nil
}

func (m *MockDB) GetUserByID(id string) (*models.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MockDB) GetAllUsers() []models.User {
	users := make([]models.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, *user)
	}
	return users
}



// Create user struct to Test Response
func TestUserToResponse(t *testing.T) {
	user := models.User{
		ID:       bson.NewObjectID(),
		Name:     "Test User",
		Email:    "test@gmail.com", 
		Password: "12345678",
		CreatedAt: time.Now(),
	}

	response := user.ToResponse()

	assert.Equal(t, user.ID, response.ID)
	assert.Equal(t, user.Name, response.Name)
	assert.Equal(t, user.Email, response.Email)
	assert.Equal(t, user.CreatedAt, response.CreatedAt)
}

// Create User Test
func TestCreateUser(t *testing.T) {

	mockDB := NewMockDB()
	
	user := &models.User{
		Name:     "mosutech",
		Email:    "mosutech@gmail.com",
		Password: "12345678",
	}
	
	err := mockDB.CreateUser(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.NotZero(t, user.CreatedAt)
	
	CheckUser, err := mockDB.GetUserByID(user.ID.Hex())
	assert.NoError(t, err)
	assert.Equal(t, user.Name, CheckUser.Name)
	assert.Equal(t, user.Email, CheckUser.Email)
}


func TestCreateUser_DuplicateEmail(t *testing.T) {

	mockDB := NewMockDB()
	
	user1 := &models.User{
		Name:     "mosutech",
		Email:    "mosutech@gmail.com",
		Password: "12345678",
	}
	
	err := mockDB.CreateUser(user1)
	assert.NoError(t, err)
	
	user2 := &models.User{
		Name:     "mosuuuutech2",
		Email:    "mosutech@gmail.com", 
		Password: "12345678",
	}
	
	err = mockDB.CreateUser(user2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate email")
	assert.Zero(t, user2.ID) 
}

// Get User by id Test
func TestGetUser(t *testing.T) {

	mockDB := NewMockDB()
	
	user := &models.User{
		Name:     "Test User",
		Email:    "Test@gmail.com",
		Password: "12345678",
	}
	
	err := mockDB.CreateUser(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	
	CheckUser, err := mockDB.GetUserByID(user.ID.Hex())
	assert.NoError(t, err)
	assert.Equal(t, user.Name, CheckUser.Name)
	assert.Equal(t, user.Email, CheckUser.Email)

	_, err = mockDB.GetUserByID("No User id")
	assert.Error(t, err)
}

// Get All Users Test
func TestGetAllUsers(t *testing.T) {

	mockDB := NewMockDB()
	
	users := mockDB.GetAllUsers()
	assert.Empty(t, users)
	
	user1 := &models.User{Name: "mosu", Email: "mosu1@gmail.com", Password: "mosu1"}
	user2 := &models.User{Name: "mosuu", Email: "mosu2@gmail.com", Password: "mosu2"}
	
	mockDB.CreateUser(user1)
	mockDB.CreateUser(user2)
	
	users = mockDB.GetAllUsers()
	assert.Len(t, users, 2)
}


// Email Validation Test
func TestValidateEmail(t *testing.T) {

	testCases := []struct {
		name     string
		email    string
		expected bool
	}{
		{"Valid email", "test@gmail.com", true},
		{"Valid email with subdomain", "user@mail.example.com", true},
		{"Invalid email - no @", "testgmail.com", false},
		{"Invalid email - no domain", "test@", false},
		{"Invalid email - no username", "@gmail.com", false},
		{"Empty email", "", false},
		{"Just @", "@", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			atIndex := strings.Index(tc.email, "@")
			isValid := strings.Count(tc.email, "@") == 1 &&
					  atIndex > 0 &&
					  atIndex < len(tc.email)-1
			
			assert.Equal(t, tc.expected, isValid, "Email: %s", tc.email)
		})
	}
}