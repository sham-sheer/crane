package server

import (
	"time"
	"github.com/lib/pq"
	"gorm.io/gorm"

)
// User is a model in the "users" table.
type User struct {
	gorm.Model
	ID   				int     	`json:"id,omitempty"`
	Name 				string 		`json:"name" gorm:"not null"`
	Email 				string 		`json:"email" gorm:"not null"`
	Contact			 	string 		`json:"contact" gorm:"not null"`
	Type 				string 		`json:"type" gorm:"not null"`
	Company 			string 		`json:"company" gorm:"not null"`
	Country 			string		`json:"country" gorm:"not null"`
	CreatedAt    		time.Time 	`json:"createdAt" gorm:"not null"`
	UpdatedAt    		time.Time 	`json:"updatedAt" gorm:"not null"`
}

type Job struct {
	gorm.Model
	ID    				int     			`json:"id,omitempty"`
	Title  				string 				`json:"title"  		gorm:"not null"`
	AppLimit  			int 				`json:"appLimit" 	gorm:"not null"`
	PosLimit			int 				`json:"posLimit		gorm:"not null"`
	Applicants          int 				`json:"applicants" 	gorm:"not null"`
	Accepted		 	int 				`json:"accepted" 	gorm:"not null"`
	Deadline			time.Time 			`json:"deadline" 	gorm:"not null"`
	Skills   			pq.StringArray 		`json: "skills" 	gorm:"type:text[]"`
	JobType				string 				`json:"jobType" 	gorm:"not null"`
	Salary 				int 				`json:"salary" 		gorm:"not null"`
	CreatedAt    		time.Time 			`json:"createdAt" 	gorm:"not null"`
	UpdatedAt    		time.Time 			`json:"updatedAt" 	gorm:"not null"`
}

// Disease is a model in the "diseases" table.
type Recruiter struct {
	gorm.Model
	ID    				int     			`json:"id,omitempty"`
	Name  				string 				`json:"name"  		gorm:"not null"`
	Contact  			int 				`json:"contact" 	gorm:"not null"`
	Company				int 				`json:"company"		gorm:"not null"`
	Rating 				int 				`json:"rating" 		gorm:"not null"`
	Country 			string				`json:"country" gorm:"not null"`
	CreatedAt    		time.Time 			`json:"createdAt" 	gorm:"not null"`
	UpdatedAt    		time.Time 			`json:"updatedAt" 	gorm:"not null"`
}

type Application struct {
	gorm.Model
	ID    				int     			`json:"id,omitempty"`
	User 				User 				`json:"userId" 		gorm:"foreignKey:UserId"`
	Recruiter 			Recruiter 			`json:"recruiterId" gorm:"foreignKey:RecruiterId"`
	Job					Job 				`json:"jobId 		gorm:"foreignKey:JobId"`
	Status 				string 				`json:"rating" 		gorm:"not null"`
	CreatedAt    		time.Time 			`json:"createdAt" 	gorm:"not null"`
	UpdatedAt    		time.Time 			`json:"updatedAt" 	gorm:"not null"`
}