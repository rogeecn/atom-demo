package dto

type UserForm struct {
	Name string `json:"name"`
}

type UserListQueryFilter struct {
}

type UserItem struct {
	Name string `json:"name"`
}
