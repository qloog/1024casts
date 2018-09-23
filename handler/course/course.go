package course

import (
	"1024casts/backend/model"
)

type CreateRequest struct {
	Name string `json:"username"`
	Type string `json:"password"`
	Description    string `json:"email"`
	Slug string `json:"slug"`
	CoverImage string `json:"cover_image"`
	UserId int `json:"user_id"`
	IsPublish int `json:"is_publish"`
}

type CreateResponse struct {
	Id string `json:"id"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64               `json:"totalCount"`
	List       []*model.CourseModel `json:"list"`
}

type SwaggerListResponse struct {
	TotalCount uint64              `json:"totalCount"`
	List       []model.CourseModel `json:"list"`
}
