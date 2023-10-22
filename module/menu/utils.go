package menu

func ToMenuResponse(menu MenuDatabaseIO) HTTPMenuResponse {
	return HTTPMenuResponse{
		Id:          menu.Id, 
		Title:       menu.Title, 
		Description: menu.Description, 
		Count:       menu.Count, 
		Price:       menu.Price, 
		CafeId:      menu.CafeId, 
		CreatedAt:   menu.CreatedAt, 
		UpdatedAt:   menu.UpdatedAt, 
	}
}

func ToMenuResponses(menus []MenuDatabaseIO) []HTTPMenuResponse {
	var menuResponses []HTTPMenuResponse
	for _, menu := range menus {
		menuResponses = append(menuResponses, ToMenuResponse(menu))
	}

	return menuResponses
}
