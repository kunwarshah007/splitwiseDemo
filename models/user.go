package models

type User struct {
	UserId       string `json:"user_id"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobileNumber"`
	Email        string `json:"email"`
}

func (u *User) Validate() {

}
