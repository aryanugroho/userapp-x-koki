package domain

type User struct {
	ID          string `json:"_id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"password" bson:"password"`
}
