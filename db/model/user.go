package model

import "time"

type User struct {
	Id           int64      `json:"id"  db:"ID"`
	RoleID       int64      `json:"role_id" db:"ROLEID"`
	FirstName    string     `json:"first_name" db:"FIRSTNAME"`
	MiddleName   string     `json:"middle_name" db:"MIDDLENAME"`
	LastName     string     `json:"last_name" db:"LASTNAME"`
	Username     string     `json:"username" db:"USERNAME"`
	Mobile       string     `json:"mobile" db:"MOBILE"`
	Email        string     `json:"email" db:"EMAIL"`
	PasswordHash string     `json:"password_hash" db:"PASSWORDHASH"`
	RegisteredAt *time.Time `json:"registered_at" db:"REGISTEREDAT"`
	LastLogin    *time.Time `json:"last_login" db:"LASTLOGIN"`
	Intro        string     `json:"intro" db:"INTRO"`
	Profile      string     `json:"profile" db:"PROFILE"`
}
