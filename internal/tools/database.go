package tools

type DatabaseInterface interface {
	SetupDatabase() error

	CreateUser(email string, username string, passwordHash string) *User
	GetUserByEmail(email string) *User
	GetUserByID(id string) *User

	CreateGroup(userID string, name string) *Group
}

func NewDatabase() (DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	err := database.SetupDatabase()
	if err != nil {
		return nil, err
	}

	return database, nil
}
