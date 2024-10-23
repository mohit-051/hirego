package ports

import "github.com/mohit-051/hirego/internal/core/domain"

type UserService interface {
	Signup(user domain.User) (string, error)
	Login(domain.User) (domain.AccessToken, error)
	SetUserProfileInformation(string, domain.UserInformation) (string, error)
	GetProfileInformation(string) (domain.UserInformation, error)
	GetUserWorkInformation(string) (domain.WorkInformation, error)
	SetUserWrorkInformation(string, domain.WorkInformation) (string, error)
	ApplyForJobPosting(domain.JobApplication) (string, error)
	GetAppliedJobPosting(string) ([]domain.JobApplication, error)
}

type UserRepository interface {
	BaseRepository[domain.User]
}
