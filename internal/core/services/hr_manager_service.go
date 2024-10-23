package service

import (
	"fmt"

	"github.com/mohit-051/hirego/internal/core/domain"
	"github.com/mohit-051/hirego/internal/core/ports"
	"github.com/mohit-051/hirego/internal/helper"
	"go.mongodb.org/mongo-driver/bson"
)

type hrService struct {
	repo ports.HRRepository
}

func InitializeHRService(r ports.HRRepository) *hrService {
	return &hrService{
		repo: r,
	}
}

/*
The Signup function is used to signup a hr manager.
*/
func (s *hrService) Signup(hr domain.HrManager) (string, error) {

	// Call the sign up repo to insert the data of the user.
	message, err := s.repo.Create(hr, "hrmanager")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Failed to insert data", nil
	}

	// Return the success message.
	return "Successfully stored the data", nil
}

/*
The Login function is used to login a hr manager.
*/
func (s *hrService) Login(hr domain.HrManager) (domain.AccessToken, error) {

	// Call the login repo to insert the data of the user.
	token, err := helper.CreateToken(hr.Username, "hr_manager")

	if err != nil {
		panic(err)
	}

	// Return the success message.

	accessToken := domain.AccessToken{
		Token: token,
	}
	return accessToken, nil
}

/*
The SetHrProfileInformation function is used to set the profile information of the hr manager.
*/
func (s *hrService) SetHrProfileInformation(username string, hrInformation domain.HrProfileInformation) (string, error) {

	hrInformation.Username = username

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(hrInformation, "hrprofilenformation")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Failed to insert data", nil
	}

	// Return the success message.
	return "Hr information stored successfully", nil
}

func (s *hrService) GetProfileInformation(username string) (domain.HrProfileInformation, error) {
	// Call the login repo to insert the data of the user.
	hrInformation, err := s.repo.GetByField("username", username, "hrprofilenformation")
	if err != nil {
		panic(err)
	}

	// Marshal the MongoDB document into a byte slice.
	bsonBytes, err := bson.Marshal(hrInformation)
	if err != nil {
		return domain.HrProfileInformation{}, err
	}

	// Unmarshal the byte slice into a domain.HrProfileInformation struct.
	var hr domain.HrProfileInformation
	err = bson.Unmarshal(bsonBytes, &hr)
	if err != nil {
		return domain.HrProfileInformation{}, err
	}

	return hr, nil
}

func (s *hrService) JobPosting(jobPosting domain.JobPosting) (string, error) {

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(jobPosting, "jobposting")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Failed to insert data", nil
	}

	// Return the success message.
	return "Job posted successfully", nil
}

func (s *hrService) GetJobPosting(username string) (domain.JobPosting, error) {
	// Call the login repo to insert the data of the user.
	jobPosting, err := s.repo.GetByField("jobID", username, "jobposting")
	if err != nil {
		panic(err)
	}

	// Marshal the MongoDB document into a byte slice.
	bsonBytes, err := bson.Marshal(jobPosting)
	if err != nil {
		return domain.JobPosting{}, err
	}

	// Unmarshal the byte slice into a domain.JobPosting struct.
	var job domain.JobPosting
	err = bson.Unmarshal(bsonBytes, &job)
	if err != nil {
		return domain.JobPosting{}, err
	}

	return job, nil
}

func (s *hrService) GetAllJobPosting() ([]domain.JobPosting, error) {
	// Call the repository to get all job postings.
	jobPostings, err := s.repo.GetAll("jobposting")
	if err != nil {
		return nil, err
	}

	var postings []domain.JobPosting
	for _, jp := range jobPostings {
		// Marshal the MongoDB document into a byte slice.
		bsonBytes, err := bson.Marshal(jp)
		if err != nil {
			return nil, err
		}

		// Unmarshal the byte slice into a domain.JobPosting struct.
		var posting domain.JobPosting
		err = bson.Unmarshal(bsonBytes, &posting)
		if err != nil {
			return nil, err
		}

		postings = append(postings, posting)
	}

	return postings, nil
}

func (s *hrService) HrSpecificJobPosting(username string) ([]domain.JobPosting, error) {
	// Call the login repo to insert the data of the user.
	jobPostings, err := s.repo.GetAllByField("Username", username, "jobposting")
	if err != nil {
		panic(err)
	}

	var postings []domain.JobPosting
	for _, jp := range jobPostings {
		// Marshal the MongoDB document into a byte slice.
		bsonBytes, err := bson.Marshal(jp)
		if err != nil {
			return nil, err
		}

		// Unmarshal the byte slice into a domain.JobPosting struct.
		var posting domain.JobPosting
		err = bson.Unmarshal(bsonBytes, &posting)
		if err != nil {
			return nil, err
		}

		postings = append(postings, posting)
	}

	return postings, nil
}

func (s *hrService) GetAllApplicants(jobID string) ([]domain.UserInformation, error) {
	// Call the login repo to insert the data of the user.
	applicants, err := s.repo.GetAllByField("jobID", jobID, "jobapplications")
	if err != nil {
		panic(err)
	}

	var users []domain.UserInformation
	for index, user := range applicants {
		// Marshal the MongoDB document into a byte slice.
		bsonBytes, err := bson.Marshal(user)
		if err != nil {
			return nil, err
		}

		fmt.Println(index, user)

		// Unmarshal the byte slice into a domain.UserInformation struct.
		var u domain.UserInformation
		err = bson.Unmarshal(bsonBytes, &u)

		fmt.Println(u.Username)
		if err != nil {
			return nil, err
		}

		userInformation, err := s.repo.GetByField("username", u.Username, "userprofile")

		fmt.Println("The user information is", userInformation)

		if err != nil {
			panic(err)
		}

		// Marshal the MongoDB document into a byte slice.
		bsonBytes, err = bson.Marshal(userInformation)
		if err != nil {
			panic(err)
		}

		// Unmarshal the byte slice into a domain.UserInformation struct.
		var userData domain.UserInformation
		err = bson.Unmarshal(bsonBytes, &userData)
		if err != nil {
			panic(err)

		}

		fmt.Println(userData)

		users = append(users, userData)
	}

	return users, nil
}
