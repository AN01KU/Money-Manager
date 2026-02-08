package tools

import (
	"time"

	"github.com/google/uuid"
)

type mockDB struct {
	users  map[string]*User
	groups map[string]*Group
}

func (m *mockDB) SetupDatabase() error {
	m.users = map[string]*User{
		"550e8400-e29b-41d4-a716-446655440001": {
			Id:           uuid.MustParse("550e8400-e29b-41d4-a716-446655440001"),
			Email:        "john@example.com",
			PasswordHash: "$2a$10$examplehash1",
			CreatedAt:    time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
			Username:     "johndoe",
			AuthToken:    "token-john-12345",
		},
		"550e8400-e29b-41d4-a716-446655440002": {
			Id:           uuid.MustParse("550e8400-e29b-41d4-a716-446655440002"),
			Email:        "jane@example.com",
			PasswordHash: "$2a$10$examplehash2",
			CreatedAt:    time.Date(2025, 2, 20, 14, 30, 0, 0, time.UTC),
			Username:     "janesmith",
			AuthToken:    "token-jane-67890",
		},
	}
	m.groups = make(map[string]*Group)
	return nil
}

func (m *mockDB) GetUserByEmail(email string) *User {
	for _, user := range m.users {
		if user.Email == email {
			return user
		}
	}
	return nil
}

func (m *mockDB) GetUserByID(id string) *User {
	if user, exists := m.users[id]; exists {
		return user
	}
	return nil
}

func (m *mockDB) CreateUser(email string, username string, passwordHash string) *User {

	user := &User{
		Id:           uuid.New(),
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		AuthToken:    uuid.New().String(),
	}

	m.users[user.Id.String()] = user
	return user
}

func (m *mockDB) CreateGroup(userID string, name string) *Group {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil
	}

	if m.GetUserByID(userID) == nil {
		return nil
	}

	group := &Group{
		ID:        uuid.New(),
		Name:      name,
		CreatedBy: userUUID,
		CreatedAt: time.Now(),
	}

	m.groups[group.ID.String()] = group
	return group
}
