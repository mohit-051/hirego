package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobPosting struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	JobID            string             `json:"jobID" bson:"jobID"`
	Username         string             `json:"username" bson:"username"`
	Title            string             `json:"title" bson:"title"`
	Description      string             `json:"description" bson:"description"`
	Responsibilities []string           `json:"responsibilities" bson:"responsibilities"`
	Requirements     []string           `json:"requirements" bson:"requirements"`
	Skills           []string           `json:"skills" bson:"skills"`
	Benefits         []string           `json:"benefits" bson:"benefits"`
	Location         string             `json:"location" bson:"location"`
	EmploymentType   string             `json:"employmentType" bson:"employmentType"`
	Industry         string             `json:"industry" bson:"industry"`
	Company          struct {
		Name        string `json:"name" bson:"name"`
		Description string `json:"description" bson:"description"`
		Size        string `json:"size" bson:"size"`
		Type        string `json:"type" bson:"type"`
		Industry    string `json:"industry" bson:"industry"`
		Website     string `json:"website" bson:"website"`
	} `json:"company" bson:"company"`
	Contact struct {
		Name  string `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
		Phone string `json:"phone" bson:"phone"`
	} `json:"contact" bson:"contact"`
	Salary struct {
		Min      float64 `json:"min" bson:"min"`
		Max      float64 `json:"max" bson:"max"`
		Currency string  `json:"currency" bson:"currency"`
		Period   string  `json:"period" bson:"period"`
	} `json:"salary" bson:"salary"`
	Remote      bool   `json:"remote" bson:"remote"`
	PublishedAt string `json:"publishedAt" bson:"publishedAt"`
	ExpiresAt   string `json:"expiresAt" bson:"expiresAt"`
}

type JobApplication struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	JobID       string             `json:"jobID" bson:"jobID"`
	Username    string             `json:"Username" bson:"Username"`
	CompanyName string             `json:"companyName" bson:"companyName"`
	JobTitle    string             `json:"jobTitle" bson:"jobTitle"`
	Description string             `json:"description" bson:"description"`
	Location    string             `json:"location" bson:"location"`
}
