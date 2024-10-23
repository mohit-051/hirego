package service

import (
	"fmt"

	"github.com/mohit-051/hirego/internal/core/domain"
	"github.com/mohit-051/hirego/internal/core/ports"
	"github.com/mohit-051/hirego/internal/helper"
	"go.mongodb.org/mongo-driver/bson"
)

type userService struct {
	repo ports.UserRepository
}

func InitializeUserService(r ports.UserRepository) *userService {
	return &userService{
		repo: r,
	}
}

// Service to signup a user.
func (s *userService) Signup(user domain.User) (string, error) {

	// Call the sign up repo to insert the data of the user.
	message, err := s.repo.Create(user, "users")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Failed to insert data", nil

	}

	// Return the success message.
	return "Successfully stored the information", nil
}

func (s *userService) Login(user domain.User) (domain.AccessToken, error) {

	// Call the login repo to insert the data of the user.
	token, err := helper.CreateToken(user.Username, "user")

	if err != nil {
		panic(err)
	}

	// Define the access token instance.
	accessToken := domain.AccessToken{
		Token: token,
	}
	return accessToken, nil
}

/*
The setUserWorkInformation function is used to set the work information of the user.
This will be public information that will be visible to all the hr managers who signs up in the platform.
*/
func (s *userService) SetUserWrorkInformation(username string, workInformation domain.WorkInformation) (string, error) {

	workInformation.Username = username

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(workInformation, "workinformation")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Some error occured", nil
	}

	// Return the success message.
	return "User work information stored successfully", nil
}

/*
The getUserWorkInformation function is used to get the work information of the user.
This will mainly be used by the HR managers to view the work information of the user.
*/
func (s *userService) GetUserWorkInformation(username string) (domain.WorkInformation, error) {

	// Call the login repo to insert the data of the user.
	workInformation, err := s.repo.GetData(username, "workinformation")

	if err != nil {
		panic(err)
	}

	// Define a variable to store the converted data.
	var convertedData domain.WorkInformation

	// Convert the data into bson json like document.
	bsonBytes, err := bson.Marshal(workInformation)

	if err != nil {
		panic(err)
	}

	// Convert bson to golang data structure.
	err = bson.Unmarshal(bsonBytes, &convertedData)
	if err != nil {
		panic(err)
	}

	fmt.Println("convertedData", convertedData)

	// Return the success message.
	return convertedData, nil
}

/*
The GetProfileInformation function is used to get the profile information of the user.
This will be completely private and only the ueer will be able to view this information. This is one to one.
*/
func (s *userService) GetProfileInformation(username string) (domain.UserInformation, error) {
	// Call the login repo to insert the data of the user.
	userInformation, err := s.repo.GetByField("username", username, "userprofile")

	if err != nil {
		return domain.UserInformation{}, err
	}
	// Define a variable to store the converted data.
	var user domain.UserInformation

	// Convert the data into bson json like document.
	bsonBytes, err := bson.Marshal(userInformation)

	if err != nil {
		panic(err)
	}

	// Convert bson to golang data structure.
	err = bson.Unmarshal(bsonBytes, &user)

	if err != nil {
		panic(err)
	}

	// Return the populated domain.UserInformation instance.
	return user, nil
}

/*
This function is used to set the profile information of the user. Users will set them for profile in our website.
*/
func (s *userService) SetUserProfileInformation(username string, userInformation domain.UserInformation) (string, error) {

	userInformation.Username = username

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(userInformation, "userprofile")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Some error occurred", nil
	}

	// Return the success message.
	return "Your profile information stored successfully", nil
}

func (s *userService) ApplyForJobPosting(jobPosting domain.JobApplication) (string, error) {

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(jobPosting, "jobapplications")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Some error occurred", nil
	}

	// Return the success message.
	return "You have successfully applied for the job", nil
}

func (s *userService) GetAppliedJobPosting(username string) ([]domain.JobApplication, error) {

	// Call the login repo to insert the data of the user.
	jobApplications, err := s.repo.GetAllByField("Username", username, "jobapplications")

	if err != nil {
		panic(err)
	}

	// Define a variable to store the converted data.
	var convertedData []domain.JobApplication

	for _, jobApplication := range jobApplications {
		// Convert the data into bson json like document.
		bsonBytes, err := bson.Marshal(jobApplication)

		if err != nil {
			panic(err)
		}

		// Convert bson to golang data structure.
		var job domain.JobApplication
		err = bson.Unmarshal(bsonBytes, &job)
		if err != nil {
			panic(err)
		}

		// Append the converted data to the convertedData slice.
		convertedData = append(convertedData, job)

	}

	// Return the success message.
	return convertedData, nil
}
