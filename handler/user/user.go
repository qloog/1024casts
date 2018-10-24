package user

import (
	"1024casts/backend/model"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Id uint64 `json:"id"`
}

type UpdateReq struct {
	Status int `json:"status"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64             `json:"totalCount"`
	UserList   []*model.UserModel `json:"userList"`
}

type SwaggerListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []model.UserModel `json:"userList"`
}
