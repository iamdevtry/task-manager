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
}

func (u *User) GetUserId() int64 {
	return u.Id
}

type UserCreate struct {
	FirstName  string `json:"first_name" db:"FIRSTNAME"`
	MiddleName string `json:"middle_name" db:"MIDDLENAME"`
	LastName   string `json:"last_name" db:"LASTNAME"`
	Username   string `json:"username" db:"USERNAME"`
	Mobile     string `json:"mobile" db:"MOBILE"`
	Email      string `json:"email" db:"EMAIL"`
	Password   string `json:"password" `
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
