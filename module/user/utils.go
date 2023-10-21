package user

func ToUserResponse(user UserDatabaseIO) HTTPUserResponse {
	return HTTPUserResponse{
		Id:        user.Id, 
		Email:     user.Email, 
		CreatedAt: user.CreatedAt, 
		UpdatedAt: user.UpdatedAt, 
	}
}

func ToUserResponses(users []UserDatabaseIO) []HTTPUserResponse {
	var userResponses []HTTPUserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
