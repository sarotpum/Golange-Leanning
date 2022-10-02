package model

import "database/sql"

type UserModelDb struct {
	UserId    sql.NullString `json:"UserId"`
	FirstName sql.NullString `json:"FirstName"`
	LastName  sql.NullString `json:"LastName"`
	Age       sql.NullInt32 `json:"Age"`
	Status    sql.NullInt32 `json:"Status"`
	CreateAt  sql.NullString `json:"CreateAt"`
}

type UserModelReq struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int `json:"Age"`
	Status    int `json:"Status"`
}

type UserModelUpdate struct {
	UserId    string `json:"UserId"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int `json:"Age"`
	Status    int `json:"Status"`
}

type UserModel struct {
	UserId    string `json:"UserId"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int `json:"Age"`
	Status    int `json:"Status"`
	CreateAt  string `json:"CreateAt"`
}
