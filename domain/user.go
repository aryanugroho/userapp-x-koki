package domain

type User struct {
	ObjID       string `json:"-" bson:"_id,omitempty"`
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"-" bson:"password"`
}
