package ports

import "github.com/mohit-051/hirego/internal/core/domain"

type HRService interface {
	Signup(domain.HrManager) (string, error)
	Login(domain.HrManager) (domain.AccessToken, error)
	SetHrProfileInformation(string, domain.HrProfileInformation) (string, error)
	GetProfileInformation(string) (domain.HrProfileInformation, error)
	JobPosting(domain.JobPosting) (string, error)
	GetJobPosting(string) (domain.JobPosting, error)
	GetAllJobPosting() ([]domain.JobPosting, error)
	HrSpecificJobPosting(string) ([]domain.JobPosting, error)
	GetAllApplicants(string) ([]domain.UserInformation, error)
}

type HRRepository interface {
	BaseRepository[domain.HrManager]
}
