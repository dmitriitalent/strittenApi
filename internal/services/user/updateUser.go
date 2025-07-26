package userService

import "github.com/dmitriitalent/strittenApi/internal/entity"

func (service *UserService) UpdateUser(user entity.User) (entity.User, error) {
	newPassword, err := service.Crypto.HashPassword(user.Password)		
	user.Password = newPassword;
	if err != nil {
		return user, err
	}

	return service.Repositories.User.UpdateUser(user);
}