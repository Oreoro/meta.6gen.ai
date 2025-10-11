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

package entity

import "time"

// FreelancerProfile freelancer profile
type FreelancerProfile struct {
	ID                string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt          time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt          time.Time `xorm:"updated TIMESTAMP updated_at"`
	UserID             string    `xorm:"not null BIGINT(20) user_id"`
	IsAvailable        bool      `xorm:"not null default true BOOL is_available"`
	HourlyRate         float64   `xorm:"not null default 0 DECIMAL(10,2) hourly_rate"`
	Currency           string    `xorm:"not null default 'USD' VARCHAR(10) currency"`
	Skills             string    `xorm:"TEXT skills"` // JSON array of skills
	Experience         string    `xorm:"TEXT experience"` // JSON object with experience details
	Portfolio          string    `xorm:"TEXT portfolio"` // JSON array of portfolio items
	Availability       string    `xorm:"VARCHAR(100) availability"` // e.g., "Full-time", "Part-time", "Project-based"
	PreferredProjects  string    `xorm:"TEXT preferred_projects"` // JSON array of preferred project types
	ContactEmail       string    `xorm:"VARCHAR(100) contact_email"` // Optional different contact email
	LinkedInProfile    string    `xorm:"VARCHAR(255) linkedin_profile"`
	GitHubProfile      string    `xorm:"VARCHAR(255) github_profile"`
	Website            string    `xorm:"VARCHAR(255) website"`
	Bio                string    `xorm:"TEXT bio"` // Extended bio for freelancing
	BioHTML            string    `xorm:"TEXT bio_html"`
	Languages          string    `xorm:"TEXT languages"` // JSON array of languages
	TimeZone           string    `xorm:"VARCHAR(50) time_zone"`
	ResponseTime       string    `xorm:"VARCHAR(50) response_time"` // e.g., "Within 24 hours"
	CompletedProjects  int       `xorm:"not null default 0 INT(11) completed_projects"`
	ClientSatisfaction float64   `xorm:"not null default 0 DECIMAL(3,2) client_satisfaction"` // 0.00 to 5.00
	IsVerified         bool      `xorm:"not null default false BOOL is_verified"`
	VerificationDate   time.Time `xorm:"TIMESTAMP verification_date"`
}

// TableName freelancer profile table name
func (FreelancerProfile) TableName() string {
	return "freelancer_profile"
}

// JobPosting job posting
type JobPosting struct {
	ID              string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt       time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt       time.Time `xorm:"updated TIMESTAMP updated_at"`
	UserID          string    `xorm:"not null BIGINT(20) user_id"`
	Title           string    `xorm:"not null VARCHAR(255) title"`
	Description     string    `xorm:"TEXT description"`
	DescriptionHTML string    `xorm:"TEXT description_html"`
	Budget          float64   `xorm:"not null default 0 DECIMAL(10,2) budget"`
	Currency        string    `xorm:"not null default 'USD' VARCHAR(10) currency"`
	BudgetType      string    `xorm:"not null default 'fixed' VARCHAR(20) budget_type"` // "fixed", "hourly", "negotiable"
	Skills          string    `xorm:"TEXT skills"` // JSON array of required skills
	ExperienceLevel string    `xorm:"VARCHAR(50) experience_level"` // "entry", "intermediate", "senior", "expert"
	Duration        string    `xorm:"VARCHAR(100) duration"` // e.g., "1-3 months", "3-6 months"
	Location        string    `xorm:"VARCHAR(100) location"` // "remote", "onsite", "hybrid"
	Status          string    `xorm:"not null default 'open' VARCHAR(20) status"` // "open", "closed", "filled"
	ContactEmail    string    `xorm:"VARCHAR(100) contact_email"`
	ApplicationCount int      `xorm:"not null default 0 INT(11) application_count"`
	ViewsCount      int       `xorm:"not null default 0 INT(11) views_count"`
	IsActive        bool      `xorm:"not null default true BOOL is_active"`
	ExpiresAt       time.Time `xorm:"TIMESTAMP expires_at"`
}

// TableName job posting table name
func (JobPosting) TableName() string {
	return "job_posting"
}

// JobApplication job application
type JobApplication struct {
	ID           string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt    time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt    time.Time `xorm:"updated TIMESTAMP updated_at"`
	JobID        string    `xorm:"not null BIGINT(20) job_id"`
	ApplicantID  string    `xorm:"not null BIGINT(20) applicant_id"`
	CoverLetter  string    `xorm:"TEXT cover_letter"`
	ProposedRate float64   `xorm:"not null default 0 DECIMAL(10,2) proposed_rate"`
	Currency     string    `xorm:"not null default 'USD' VARCHAR(10) currency"`
	Status       string    `xorm:"not null default 'pending' VARCHAR(20) status"` // "pending", "accepted", "rejected", "withdrawn"
	Message      string    `xorm:"TEXT message"` // Additional message from applicant
}

// TableName job application table name
func (JobApplication) TableName() string {
	return "job_application"
}
