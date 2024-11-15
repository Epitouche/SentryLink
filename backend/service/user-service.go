package service

type UserService interface {
	User(username string, password string) bool
}

type userService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewUserService() UserService {
	return &userService{
		authorizedUsername: "root",
		authorizedPassword: "password",
	}
}

func (service *userService) User(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
