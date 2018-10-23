package models

//ProfileViewModel .
type UserModel struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Alias     string `json:"alias"`
	INumber   string `json:"inumber"`
}

type UsersModel struct {
	Models []*UserModel `json:"users"`
}
