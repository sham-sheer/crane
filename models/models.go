package models

import (
	"time"
	"gorm.io/gorm"

)
// User is a model in the "users" table.
type User struct {
	gorm.Model
	ID   				int     	`gorm:"primaryKey"`
	Name 				string 		`json:"name" `
	Email 				string 		`json:"email" `
	Contact			 	string 		`json:"contact" `
	Type 				string 		`json:"type" `
	Company 			string 		`json:"company" `
	Country 			string		`json:"country" `
	CreatedAt    		time.Time 	`json:"createdAt" `
	UpdatedAt    		time.Time 	`json:"updatedAt" `
}

type Job struct {
	gorm.Model
	ID    				int     			`gorm:"primaryKey"`
	Title  				string 				`json:"title"  		`
	AppLimit  			int 				`json:"appLimit" 	`
	PosLimit			int 				`json:"posLimit		`
	Applicants          int 				`json:"applicants" 	`
	Accepted		 	int 				`json:"accepted" 	`
	Deadline			time.Time 			`json:"deadline" 	`
	Skills   			string				`json:"skills" 		`
	JobType				string 				`json:"jobType" 	`
	Salary 				int 				`json:"salary" 		`
	CreatedAt    		time.Time 			`json:"createdAt" 	`
	UpdatedAt    		time.Time 			`json:"updatedAt" 	`
}

// Disease is a model in the "diseases" table.
type Recruiter struct {
	gorm.Model
	ID    				int     			`gorm:"primaryKey"`
	Name  				string 				`json:"name"  		`
	Contact  			int 				`json:"contact" 	`
	Company				int 				`json:"company"		`
	Rating 				int 				`json:"rating" 		`
	Country 			string				`json:"country" `
	CreatedAt    		time.Time 			`json:"createdAt" 	`
	UpdatedAt    		time.Time 			`json:"updatedAt" 	`
}

type Application struct {
	gorm.Model
	ID    				int     			`gorm:"primaryKey"`
	UserId    			int					`json:"userId"`
	User 				User 				`gorm:"foreignKey:UserId;references:ID"`
	RecruiterId    		int					`json:"recruiterId"`
	Recruiter 			Recruiter 			`gorm:"foreignKey:RecruiterId;references:ID"`
	JobId	    		int					`json:"jobId"`
	Job					Job 				`gorm:"foreignKey:JobId;references:ID"`
	Status 				string 				`json:"status" 		`
	CreatedAt    		time.Time 			`json:"createdAt" 	`
	UpdatedAt    		time.Time 			`json:"updatedAt" 	`
}