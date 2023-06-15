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
	ID       *int32     `json:"id,omitempty"`       // UserID
	Username *string    `json:"username,omitempty"` // UserName
	Age      *int32     `json:"age,omitempty"`      // UserAge
	Sex      *string    `json:"sex,omitempty"`      // UserSex
	Birthday *time.Time `json:"birthday,omitempty"`
	Status   *string    `json:"status,omitempty"`
	State    *string    `json:"state,omitempty"`
}
