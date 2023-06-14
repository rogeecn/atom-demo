package dto

import (
	"time"
)

type UserForm struct {
	UserItem `json:",inline"`
}

type UserListQueryFilter struct {
	UserItem `json:",inline"`
}

type UserItem struct {
	ID       *int32     `json:"id,omitempty""`
	Username *string    `json:"username,omitempty"`
	Age      *int32     `json:"age,omitempty"`
	Sex      *string    `json:"sex,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
	Status   *string    `json:"status,omitempty"`
	State    *string    `json:"state,omitempty"`
}
