package cafe

func ToCafeResponse(cafe CafeDatabaseIO) HTTPCafeResponse {
	return HTTPCafeResponse{
		Id:        cafe.Id, 
		Email:     cafe.Email, 
		Title:     cafe.Title, 
		CreatedAt: cafe.CreatedAt, 
		UpdatedAt: cafe.UpdatedAt, 
	}
}

func ToCafeResponses(cafes []CafeDatabaseIO) []HTTPCafeResponse {
	var cafeResponses []HTTPCafeResponse
	for _, cafe := range cafes {
		cafeResponses = append(cafeResponses, ToCafeResponse(cafe))
	}

	return cafeResponses
}
