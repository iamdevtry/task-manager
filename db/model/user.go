package model

type User struct {
	Id           int64   `json:"id"  db:"ID"`
	FirstName    *string `json:"first_name,omitempty" db:"FIRSTNAME"`
	MiddleName   *string `json:"middle_name,omitempty" db:"MIDDLENAME"`
	LastName     *string `json:"last_name,omitempty" db:"LASTNAME"`
	Username     string  `json:"username,omitempty" db:"USERNAME"`
	Mobile       string  `json:"mobile,omitempty" db:"MOBILE"`
	Email        string  `json:"email,omitempty" db:"EMAIL"`
	PasswordHash string  `json:"-" db:"PASSWORDHASH"`
	Intro        *string `json:"intro,omitempty" db:"INTRO"`
	Profile      *string `json:"profile,omitempty" db:"PROFILE"`
}

type CreateUser struct {
	FirstName  string `json:"first_name" db:"FIRSTNAME"`
	MiddleName string `json:"middle_name" db:"MIDDLENAME"`
	LastName   string `json:"last_name" db:"LASTNAME"`
	Username   string `json:"username" db:"USERNAME"`
	Mobile     string `json:"mobile" db:"MOBILE"`
	Email      string `json:"email" db:"EMAIL"`
	Password   string `json:"password" `
	Intro      string `json:"intro" db:"INTRO"`
	Profile    string `json:"profile" db:"PROFILE"`
}
