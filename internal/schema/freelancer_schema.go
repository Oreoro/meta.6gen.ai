/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package schema

// FreelancerProfileResp freelancer profile response
type FreelancerProfileResp struct {
	ID                 string   `json:"id"`
	UserID             string   `json:"user_id"`
	IsAvailable        bool     `json:"is_available"`
	HourlyRate         float64  `json:"hourly_rate"`
	Currency           string   `json:"currency"`
	Skills             []string `json:"skills"`
	Experience         string   `json:"experience"`
	Portfolio          []string `json:"portfolio"`
	Availability       string   `json:"availability"`
	PreferredProjects  []string `json:"preferred_projects"`
	ContactEmail       string   `json:"contact_email"`
	LinkedInProfile    string   `json:"linkedin_profile"`
	GitHubProfile      string   `json:"github_profile"`
	Website            string   `json:"website"`
	Bio                string   `json:"bio"`
	BioHTML            string   `json:"bio_html"`
	Languages          []string `json:"languages"`
	TimeZone           string   `json:"time_zone"`
	ResponseTime       string   `json:"response_time"`
	CompletedProjects  int      `json:"completed_projects"`
	ClientSatisfaction float64  `json:"client_satisfaction"`
	IsVerified         bool     `json:"is_verified"`
	VerificationDate   int64    `json:"verification_date"`
	CreatedAt          int64    `json:"created_at"`
	UpdatedAt          int64    `json:"updated_at"`
}

// CreateFreelancerProfileReq create freelancer profile request
type CreateFreelancerProfileReq struct {
	IsAvailable       bool     `json:"is_available"`
	HourlyRate        float64  `json:"hourly_rate"`
	Currency          string   `json:"currency"`
	Skills            []string `json:"skills"`
	Experience        string   `json:"experience"`
	Portfolio         []string `json:"portfolio"`
	Availability      string   `json:"availability"`
	PreferredProjects []string `json:"preferred_projects"`
	ContactEmail      string   `json:"contact_email"`
	LinkedInProfile   string   `json:"linkedin_profile"`
	GitHubProfile     string   `json:"github_profile"`
	Website           string   `json:"website"`
	Bio               string   `json:"bio"`
	Languages         []string `json:"languages"`
	TimeZone          string   `json:"time_zone"`
	ResponseTime      string   `json:"response_time"`
	LoginUserID       string   `json:"-"`
}

// UpdateFreelancerProfileReq update freelancer profile request
type UpdateFreelancerProfileReq struct {
	ID                string   `validate:"required" json:"id"`
	IsAvailable       bool     `json:"is_available"`
	HourlyRate        float64  `json:"hourly_rate"`
	Currency          string   `json:"currency"`
	Skills            []string `json:"skills"`
	Experience        string   `json:"experience"`
	Portfolio         []string `json:"portfolio"`
	Availability      string   `json:"availability"`
	PreferredProjects []string `json:"preferred_projects"`
	ContactEmail      string   `json:"contact_email"`
	LinkedInProfile   string   `json:"linkedin_profile"`
	GitHubProfile     string   `json:"github_profile"`
	Website           string   `json:"website"`
	Bio               string   `json:"bio"`
	Languages         []string `json:"languages"`
	TimeZone          string   `json:"time_zone"`
	ResponseTime      string   `json:"response_time"`
	LoginUserID       string   `json:"-"`
}

// GetFreelancerProfileReq get freelancer profile request
type GetFreelancerProfileReq struct {
	UserID string `validate:"required" json:"user_id" form:"user_id"`
}

// GetFreelancerProfilesReq get freelancer profiles request
type GetFreelancerProfilesReq struct {
	Page     int     `json:"page" form:"page"`
	PageSize int     `json:"page_size" form:"page_size"`
	Skills   string  `json:"skills" form:"skills"`
	Location string  `json:"location" form:"location"`
	MinRate  float64 `json:"min_rate" form:"min_rate"`
	MaxRate  float64 `json:"max_rate" form:"max_rate"`
	Currency string  `json:"currency" form:"currency"`
}

// GetFreelancerProfilesResp get freelancer profiles response
type GetFreelancerProfilesResp struct {
	Count int                      `json:"count"`
	List  []*FreelancerProfileResp `json:"list"`
}

// JobPostingResp job posting response
type JobPostingResp struct {
	ID               string   `json:"id"`
	UserID           string   `json:"user_id"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	DescriptionHTML  string   `json:"description_html"`
	Budget           float64  `json:"budget"`
	Currency         string   `json:"currency"`
	BudgetType       string   `json:"budget_type"`
	Skills           []string `json:"skills"`
	ExperienceLevel  string   `json:"experience_level"`
	Duration         string   `json:"duration"`
	Location         string   `json:"location"`
	Status           string   `json:"status"`
	ContactEmail     string   `json:"contact_email"`
	ApplicationCount int      `json:"application_count"`
	ViewsCount       int      `json:"views_count"`
	IsActive         bool     `json:"is_active"`
	ExpiresAt        int64    `json:"expires_at"`
	CreatedAt        int64    `json:"created_at"`
	UpdatedAt        int64    `json:"updated_at"`
}

// CreateJobPostingReq create job posting request
type CreateJobPostingReq struct {
	Title           string   `validate:"required" json:"title"`
	Description     string   `validate:"required" json:"description"`
	Budget          float64  `json:"budget"`
	Currency        string   `json:"currency"`
	BudgetType      string   `json:"budget_type"`
	Skills          []string `json:"skills"`
	ExperienceLevel string   `json:"experience_level"`
	Duration        string   `json:"duration"`
	Location        string   `json:"location"`
	ContactEmail    string   `json:"contact_email"`
	ExpiresAt       int64    `json:"expires_at"`
	LoginUserID     string   `json:"-"`
}

// UpdateJobPostingReq update job posting request
type UpdateJobPostingReq struct {
	ID              string   `validate:"required" json:"id"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Budget          float64  `json:"budget"`
	Currency        string   `json:"currency"`
	BudgetType      string   `json:"budget_type"`
	Skills          []string `json:"skills"`
	ExperienceLevel string   `json:"experience_level"`
	Duration        string   `json:"duration"`
	Location        string   `json:"location"`
	ContactEmail    string   `json:"contact_email"`
	Status          string   `json:"status"`
	ExpiresAt       int64    `json:"expires_at"`
	LoginUserID     string   `json:"-"`
}

// GetJobPostingsReq get job postings request
type GetJobPostingsReq struct {
	Page      int     `json:"page" form:"page"`
	PageSize  int     `json:"page_size" form:"page_size"`
	Skills    string  `json:"skills" form:"skills"`
	Location  string  `json:"location" form:"location"`
	MinBudget float64 `json:"min_budget" form:"min_budget"`
	MaxBudget float64 `json:"max_budget" form:"max_budget"`
	Currency  string  `json:"currency" form:"currency"`
	Status    string  `json:"status" form:"status"`
}

// GetJobPostingsResp get job postings response
type GetJobPostingsResp struct {
	Count int               `json:"count"`
	List  []*JobPostingResp `json:"list"`
}

// JobApplicationResp job application response
type JobApplicationResp struct {
	ID           string  `json:"id"`
	JobID        string  `json:"job_id"`
	ApplicantID  string  `json:"applicant_id"`
	CoverLetter  string  `json:"cover_letter"`
	ProposedRate float64 `json:"proposed_rate"`
	Currency     string  `json:"currency"`
	Status       string  `json:"status"`
	Message      string  `json:"message"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
}

// CreateJobApplicationReq create job application request
type CreateJobApplicationReq struct {
	JobID        string  `validate:"required" json:"job_id"`
	CoverLetter  string  `json:"cover_letter"`
	ProposedRate float64 `json:"proposed_rate"`
	Currency     string  `json:"currency"`
	Message      string  `json:"message"`
	LoginUserID  string  `json:"-"`
}

// GetJobApplicationsReq get job applications request
type GetJobApplicationsReq struct {
	JobID       string `json:"job_id" form:"job_id"`
	ApplicantID string `json:"applicant_id" form:"applicant_id"`
	Status      string `json:"status" form:"status"`
	Page        int    `json:"page" form:"page"`
	PageSize    int    `json:"page_size" form:"page_size"`
}

// GetJobApplicationsResp get job applications response
type GetJobApplicationsResp struct {
	Count int                   `json:"count"`
	List  []*JobApplicationResp `json:"list"`
}

// HireFreelancerReq hire freelancer request
type HireFreelancerReq struct {
	FreelancerUserID string `validate:"required" json:"freelancer_user_id"`
	Subject          string `validate:"required" json:"subject"`
	Message          string `validate:"required" json:"message"`
	LoginUserID      string `json:"-"`
}

// HireFreelancerResp hire freelancer response
type HireFreelancerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
