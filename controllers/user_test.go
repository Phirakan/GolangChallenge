package controllers

import (
	"testing"
	"strings"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	
	"golang-test/models"
)

func TestUserToResponse_Unit(t *testing.T) {
	user := models.User{
		ID:       bson.NewObjectID(),
		Name:     "Test User",
		Email:    "test@example.com", 
		Password: "hashed_password",
	}

	response := user.ToResponse()

	assert.Equal(t, user.ID, response.ID)
	assert.Equal(t, user.Name, response.Name)
	assert.Equal(t, user.Email, response.Email)

}

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		expected bool
	}{
		{"Valid email", "test@gmail.com", true},
		{"Invalid email - no @", "testgmail.com", false},
		{"Invalid email - no domain", "test@", false},
		{"Empty email", "", false},
	}

for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
        atIndex := strings.Index(tc.email, "@")
        isValid := strings.Count(tc.email, "@") == 1 &&
                  atIndex > 0 &&
                  atIndex < len(tc.email)-1
        
        assert.Equal(t, tc.expected, isValid)
    })
}
}