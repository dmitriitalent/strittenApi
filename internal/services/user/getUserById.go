package userService

import "github.com/dmitriitalent/strittenApi/internal/entity"

func (service *UserService) GetUserById(userId int) (entity.User, error){
	user, err := service.Repositories.User.FindUserById(userId)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}