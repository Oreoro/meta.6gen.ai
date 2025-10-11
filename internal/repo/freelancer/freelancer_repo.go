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

package freelancer

import (
	"context"

	"github.com/apache/answer/internal/base/data"
	"github.com/apache/answer/internal/base/reason"
	"github.com/apache/answer/internal/entity"
	"github.com/apache/answer/internal/schema"
	"github.com/segmentfault/pacman/errors"
)

// FreelancerRepo freelancer repository
type FreelancerRepo interface {
	CreateFreelancerProfile(ctx context.Context, profile *entity.FreelancerProfile) error
	UpdateFreelancerProfile(ctx context.Context, profile *entity.FreelancerProfile) error
	GetFreelancerProfileByUserID(ctx context.Context, userID string) (*entity.FreelancerProfile, bool, error)
	GetFreelancerProfiles(ctx context.Context, req *schema.GetFreelancerProfilesReq) ([]*entity.FreelancerProfile, int64, error)
	DeleteFreelancerProfile(ctx context.Context, userID string) error

	CreateJobPosting(ctx context.Context, posting *entity.JobPosting) error
	UpdateJobPosting(ctx context.Context, posting *entity.JobPosting) error
	GetJobPostingByID(ctx context.Context, id string) (*entity.JobPosting, bool, error)
	GetJobPostings(ctx context.Context, req *schema.GetJobPostingsReq) ([]*entity.JobPosting, int64, error)
	GetJobPostingsByUserID(ctx context.Context, userID string) ([]*entity.JobPosting, error)
	DeleteJobPosting(ctx context.Context, id string) error
	IncrementJobViews(ctx context.Context, id string) error

	CreateJobApplication(ctx context.Context, application *entity.JobApplication) error
	UpdateJobApplication(ctx context.Context, application *entity.JobApplication) error
	GetJobApplicationByID(ctx context.Context, id string) (*entity.JobApplication, bool, error)
	GetJobApplications(ctx context.Context, req *schema.GetJobApplicationsReq) ([]*entity.JobApplication, int64, error)
	GetJobApplicationsByJobID(ctx context.Context, jobID string) ([]*entity.JobApplication, error)
	GetJobApplicationsByApplicantID(ctx context.Context, applicantID string) ([]*entity.JobApplication, error)
	DeleteJobApplication(ctx context.Context, id string) error
}

type freelancerRepo struct {
	data *data.Data
}

// NewFreelancerRepo new freelancer repository
func NewFreelancerRepo(data *data.Data) FreelancerRepo {
	return &freelancerRepo{
		data: data,
	}
}

// CreateFreelancerProfile create freelancer profile
func (fr *freelancerRepo) CreateFreelancerProfile(ctx context.Context, profile *entity.FreelancerProfile) error {
	_, err := fr.data.DB.Context(ctx).Insert(profile)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// UpdateFreelancerProfile update freelancer profile
func (fr *freelancerRepo) UpdateFreelancerProfile(ctx context.Context, profile *entity.FreelancerProfile) error {
	_, err := fr.data.DB.Context(ctx).Where("user_id = ?", profile.UserID).Update(profile)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// GetFreelancerProfileByUserID get freelancer profile by user ID
func (fr *freelancerRepo) GetFreelancerProfileByUserID(ctx context.Context, userID string) (*entity.FreelancerProfile, bool, error) {
	profile := &entity.FreelancerProfile{}
	exist, err := fr.data.DB.Context(ctx).Where("user_id = ?", userID).Get(profile)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return profile, exist, nil
}

// GetFreelancerProfiles get freelancer profiles
func (fr *freelancerRepo) GetFreelancerProfiles(ctx context.Context, req *schema.GetFreelancerProfilesReq) ([]*entity.FreelancerProfile, int64, error) {
	session := fr.data.DB.Context(ctx).Where("is_available = ?", true)
	
	// Apply filters
	if req.Skills != "" {
		session = session.Where("skills LIKE ?", "%"+req.Skills+"%")
	}
	if req.Location != "" {
		session = session.Where("location LIKE ?", "%"+req.Location+"%")
	}
	if req.MinRate > 0 {
		session = session.Where("hourly_rate >= ?", req.MinRate)
	}
	if req.MaxRate > 0 {
		session = session.Where("hourly_rate <= ?", req.MaxRate)
	}
	if req.Currency != "" {
		session = session.Where("currency = ?", req.Currency)
	}

	// Get total count
	total, err := session.Count(&entity.FreelancerProfile{})
	if err != nil {
		return nil, 0, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Apply pagination
	if req.Page > 0 {
		session = session.Limit(req.PageSize, (req.Page-1)*req.PageSize)
	}

	var profiles []*entity.FreelancerProfile
	err = session.OrderBy("created_at DESC").Find(&profiles)
	if err != nil {
		return nil, 0, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	return profiles, total, nil
}

// DeleteFreelancerProfile delete freelancer profile
func (fr *freelancerRepo) DeleteFreelancerProfile(ctx context.Context, userID string) error {
	_, err := fr.data.DB.Context(ctx).Where("user_id = ?", userID).Delete(&entity.FreelancerProfile{})
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// CreateJobPosting create job posting
func (fr *freelancerRepo) CreateJobPosting(ctx context.Context, posting *entity.JobPosting) error {
	_, err := fr.data.DB.Context(ctx).Insert(posting)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// UpdateJobPosting update job posting
func (fr *freelancerRepo) UpdateJobPosting(ctx context.Context, posting *entity.JobPosting) error {
	_, err := fr.data.DB.Context(ctx).Where("id = ?", posting.ID).Update(posting)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// GetJobPostingByID get job posting by ID
func (fr *freelancerRepo) GetJobPostingByID(ctx context.Context, id string) (*entity.JobPosting, bool, error) {
	posting := &entity.JobPosting{}
	exist, err := fr.data.DB.Context(ctx).Where("id = ?", id).Get(posting)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return posting, exist, nil
}

// GetJobPostings get job postings
func (fr *freelancerRepo) GetJobPostings(ctx context.Context, req *schema.GetJobPostingsReq) ([]*entity.JobPosting, int64, error) {
	session := fr.data.DB.Context(ctx).Where("is_active = ?", true)
	
	// Apply filters
	if req.Skills != "" {
		session = session.Where("skills LIKE ?", "%"+req.Skills+"%")
	}
	if req.Location != "" {
		session = session.Where("location LIKE ?", "%"+req.Location+"%")
	}
	if req.MinBudget > 0 {
		session = session.Where("budget >= ?", req.MinBudget)
	}
	if req.MaxBudget > 0 {
		session = session.Where("budget <= ?", req.MaxBudget)
	}
	if req.Currency != "" {
		session = session.Where("currency = ?", req.Currency)
	}
	if req.Status != "" {
		session = session.Where("status = ?", req.Status)
	}

	// Get total count
	total, err := session.Count(&entity.JobPosting{})
	if err != nil {
		return nil, 0, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Apply pagination
	if req.Page > 0 {
		session = session.Limit(req.PageSize, (req.Page-1)*req.PageSize)
	}

	var postings []*entity.JobPosting
	err = session.OrderBy("created_at DESC").Find(&postings)
	if err != nil {
		return nil, 0, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	return postings, total, nil
}

// GetJobPostingsByUserID get job postings by user ID
func (fr *freelancerRepo) GetJobPostingsByUserID(ctx context.Context, userID string) ([]*entity.JobPosting, error) {
	var postings []*entity.JobPosting
	err := fr.data.DB.Context(ctx).Where("user_id = ?", userID).OrderBy("created_at DESC").Find(&postings)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return postings, nil
}

// DeleteJobPosting delete job posting
func (fr *freelancerRepo) DeleteJobPosting(ctx context.Context, id string) error {
	_, err := fr.data.DB.Context(ctx).Where("id = ?", id).Delete(&entity.JobPosting{})
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// IncrementJobViews increment job views count
func (fr *freelancerRepo) IncrementJobViews(ctx context.Context, id string) error {
	_, err := fr.data.DB.Context(ctx).Where("id = ?", id).Incr("views_count").Update(&entity.JobPosting{})
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// CreateJobApplication create job application
func (fr *freelancerRepo) CreateJobApplication(ctx context.Context, application *entity.JobApplication) error {
	_, err := fr.data.DB.Context(ctx).Insert(application)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// UpdateJobApplication update job application
func (fr *freelancerRepo) UpdateJobApplication(ctx context.Context, application *entity.JobApplication) error {
	_, err := fr.data.DB.Context(ctx).Where("id = ?", application.ID).Update(application)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// GetJobApplicationByID get job application by ID
func (fr *freelancerRepo) GetJobApplicationByID(ctx context.Context, id string) (*entity.JobApplication, bool, error) {
	application := &entity.JobApplication{}
	exist, err := fr.data.DB.Context(ctx).Where("id = ?", id).Get(application)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return application, exist, nil
}

// GetJobApplications get job applications
func (fr *freelancerRepo) GetJobApplications(ctx context.Context, req *schema.GetJobApplicationsReq) ([]*entity.JobApplication, int64, error) {
	session := fr.data.DB.Context(ctx)
	
	// Apply filters
	if req.JobID != "" {
		session = session.Where("job_id = ?", req.JobID)
	}
	if req.ApplicantID != "" {
		session = session.Where("applicant_id = ?", req.ApplicantID)
	}
	if req.Status != "" {
		session = session.Where("status = ?", req.Status)
	}

	// Get total count
	total, err := session.Count(&entity.JobApplication{})
	if err != nil {
		return nil, 0, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Apply pagination
	if req.Page > 0 {
		session = session.Limit(req.PageSize, (req.Page-1)*req.PageSize)
	}

	var applications []*entity.JobApplication
	err = session.OrderBy("created_at DESC").Find(&applications)
	if err != nil {
		return nil, 0, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	return applications, total, nil
}

// GetJobApplicationsByJobID get job applications by job ID
func (fr *freelancerRepo) GetJobApplicationsByJobID(ctx context.Context, jobID string) ([]*entity.JobApplication, error) {
	var applications []*entity.JobApplication
	err := fr.data.DB.Context(ctx).Where("job_id = ?", jobID).OrderBy("created_at DESC").Find(&applications)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return applications, nil
}

// GetJobApplicationsByApplicantID get job applications by applicant ID
func (fr *freelancerRepo) GetJobApplicationsByApplicantID(ctx context.Context, applicantID string) ([]*entity.JobApplication, error) {
	var applications []*entity.JobApplication
	err := fr.data.DB.Context(ctx).Where("applicant_id = ?", applicantID).OrderBy("created_at DESC").Find(&applications)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return applications, nil
}

// DeleteJobApplication delete job application
func (fr *freelancerRepo) DeleteJobApplication(ctx context.Context, id string) error {
	_, err := fr.data.DB.Context(ctx).Where("id = ?", id).Delete(&entity.JobApplication{})
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}
