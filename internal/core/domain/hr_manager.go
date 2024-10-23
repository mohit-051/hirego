package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type HrManager struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type HrProfileInformation struct {
	ID               primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	Username         string                 `json:"username" bson:"username"`
	Email            string                 `json:"email" bson:"email"`
	FirstName        string                 `json:"firstname" bson:"firstname"`
	LastName         string                 `json:"lastname" bson:"lastname"`
	PhoneNumber      string                 `json:"phonenumber" bson:"phonenumber"`
	DOB              string                 `json:"dob" bson:"dob"`
	Address          string                 `json:"address" bson:"address"`
	ProfilePicture   string                 `json:"profilepicture" bson:"profilepicture"`
	Country          string                 `json:"country" bson:"country"`
	State            string                 `json:"state" bson:"state"`
	PushNotification string                 `json:"pushnotification" bson:"pushnotification"`
	Notifications    map[string]interface{} `json:"Notifications" bson:"Notifications"`
}
