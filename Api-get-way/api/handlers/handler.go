package handlers

import "Github.com/LocalEats/Api-get-way/genproto"

type Handler struct {
	UserService  genproto.UserServiceClient
	OrderService genproto.OrderServiceClient
}

func NewHandler(us genproto.UserServiceClient, o genproto.OrderServiceClient) *Handler {
	return &Handler{UserService: us, OrderService: o}
}
