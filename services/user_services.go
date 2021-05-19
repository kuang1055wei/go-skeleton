package services

var UserService = newUserTokenService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}
