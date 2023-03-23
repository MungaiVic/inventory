package service

import (
	"inv-v2/internal/models"
)

func ConvertUserRegToUserModel(userReg UserRegistration) *models.User {
	userObj := new(models.User)
	userObj.FirstName = userReg.FirstName
	userObj.LastName = userReg.LastName
	userObj.Email = userReg.Email
	userObj.Username = userReg.Username
	userObj.Password = userReg.Password
	return userObj
}

func ConvertUserRegToUserResponse(userReg models.User) UserResponse {
	var userObj UserResponse
	userObj.FirstName = userReg.FirstName
	userObj.LastName = userReg.LastName
	userObj.Email = userReg.Email
	userObj.Username = userReg.Username

	return userObj
}

func ConvertUserModelsToUserResponses(userModels []*models.User) []UserResponse {
	var userResps []UserResponse

	for _, user := range userModels {
		userResponse := UserResponse{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Username:  user.Username,
		}
		userResps = append(userResps, userResponse)
	}
	return userResps
}

func ConvertUserModelToUserResponse(userModels *models.User) UserResponse {

	userResponse := UserResponse{
		FirstName: userModels.FirstName,
		LastName:  userModels.LastName,
		Email:     userModels.Email,
		Username:  userModels.Username,
	}

	return userResponse
}

func ConvertItemModelToItemResponse(itemModels []models.Item) []ItemResponse {
	var itemResps []ItemResponse

	for _, item := range itemModels {
		itemResponse := ItemResponse{
			Name: item.Name,
			Quantity: item.Quantity,
			Price: item.Price,
			Reorderlvl: item.Reorderlvl,
		}
		itemResps = append(itemResps, itemResponse)
	}
	return itemResps
}