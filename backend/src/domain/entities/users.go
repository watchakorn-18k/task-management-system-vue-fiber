package entities

type User struct {
	UserID   string `json:"user_id" bson:"user_id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
	JWT      string `json:"jwt" bson:"jwt,omitempty"`
}

type UserModel struct {
	Username        string `json:"username" bson:"username"`
	Password        string `json:"password" bson:"password"`
	PasswordConfirm string `json:"password_confirm" bson:"password_confirm"`
}
