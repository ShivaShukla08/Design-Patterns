package main

type UserService struct {
	db *DatabaseConnection
}

func NewUserService() *UserService {
	return &UserService{
		db: GetInstance(), 
	}
}

func (u *UserService) GetUsers() {
	u.db.ExecuteQuery("SELECT * FROM users")
}