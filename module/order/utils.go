package order

func ToOrderResponse(order OrderDatabaseIO) HTTPOrderResponse {
	return HTTPOrderResponse{
		Id:          order.Id, 
		Notes:       order.Notes, 
		Count:       order.Count, 
		Status:      order.Status,
		Address:     order.Address,
		MenuId:      order.MenuId, 
		UserId:      order.UserId,
		CreatedAt:   order.CreatedAt, 
		UpdatedAt:   order.UpdatedAt, 
	}
}

func ToOrderResponses(orders []OrderDatabaseIO) []HTTPOrderResponse {
	var orderResponses []HTTPOrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, ToOrderResponse(order))
	}

	return orderResponses
}
