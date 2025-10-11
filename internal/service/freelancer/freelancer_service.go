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
	"encoding/json"
	"time"

	"github.com/apache/answer/internal/base/reason"
	"github.com/apache/answer/internal/entity"
	"github.com/apache/answer/internal/repo/freelancer"
	"github.com/apache/answer/internal/schema"
	"github.com/apache/answer/internal/service/export"
	"github.com/apache/answer/internal/service/siteinfo_common"
	"github.com/apache/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/errors"
)

// FreelancerService freelancer service
type FreelancerService struct {
	freelancerRepo freelancer.FreelancerRepo
	userRepo       usercommon.UserRepo
	emailService   *export.EmailService
	siteInfoService siteinfo_common.SiteInfoCommonService
}

// NewFreelancerService new freelancer service
func NewFreelancerService(
	freelancerRepo freelancer.FreelancerRepo,
	userRepo usercommon.UserRepo,
	emailService *export.EmailService,
	siteInfoService siteinfo_common.SiteInfoCommonService,
) *FreelancerService {
	return &FreelancerService{
		freelancerRepo:  freelancerRepo,
		userRepo:        userRepo,
		emailService:    emailService,
		siteInfoService: siteInfoService,
	}
}

// CreateFreelancerProfile create freelancer profile
func (fs *FreelancerService) CreateFreelancerProfile(ctx context.Context, req *schema.CreateFreelancerProfileReq) error {
	// Check if profile already exists
	_, exist, err := fs.freelancerRepo.GetFreelancerProfileByUserID(ctx, req.LoginUserID)
	if err != nil {
		return err
	}
	if exist {
		return errors.BadRequest(reason.FreelancerProfileAlreadyExists)
	}

	// Convert skills to JSON
	skillsJSON, err := json.Marshal(req.Skills)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Convert portfolio to JSON
	portfolioJSON, err := json.Marshal(req.Portfolio)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Convert preferred projects to JSON
	preferredProjectsJSON, err := json.Marshal(req.PreferredProjects)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Convert languages to JSON
	languagesJSON, err := json.Marshal(req.Languages)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	profile := &entity.FreelancerProfile{
		UserID:            req.LoginUserID,
		IsAvailable:       req.IsAvailable,
		HourlyRate:        req.HourlyRate,
		Currency:          req.Currency,
		Skills:            string(skillsJSON),
		Experience:        req.Experience,
		Portfolio:         string(portfolioJSON),
		Availability:      req.Availability,
		PreferredProjects: string(preferredProjectsJSON),
		ContactEmail:      req.ContactEmail,
		LinkedInProfile:   req.LinkedInProfile,
		GitHubProfile:     req.GitHubProfile,
		Website:           req.Website,
		Bio:               req.Bio,
		Languages:         string(languagesJSON),
		TimeZone:          req.TimeZone,
		ResponseTime:      req.ResponseTime,
	}

	return fs.freelancerRepo.CreateFreelancerProfile(ctx, profile)
}

// UpdateFreelancerProfile update freelancer profile
func (fs *FreelancerService) UpdateFreelancerProfile(ctx context.Context, req *schema.UpdateFreelancerProfileReq) error {
	// Get existing profile
	profile, exist, err := fs.freelancerRepo.GetFreelancerProfileByUserID(ctx, req.LoginUserID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.NotFound(reason.FreelancerProfileNotFound)
	}

	// Convert skills to JSON
	skillsJSON, err := json.Marshal(req.Skills)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Convert portfolio to JSON
	portfolioJSON, err := json.Marshal(req.Portfolio)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Convert preferred projects to JSON
	preferredProjectsJSON, err := json.Marshal(req.PreferredProjects)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Convert languages to JSON
	languagesJSON, err := json.Marshal(req.Languages)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// Update fields
	profile.IsAvailable = req.IsAvailable
	profile.HourlyRate = req.HourlyRate
	profile.Currency = req.Currency
	profile.Skills = string(skillsJSON)
	profile.Experience = req.Experience
	profile.Portfolio = string(portfolioJSON)
	profile.Availability = req.Availability
	profile.PreferredProjects = string(preferredProjectsJSON)
	profile.ContactEmail = req.ContactEmail
	profile.LinkedInProfile = req.LinkedInProfile
	profile.GitHubProfile = req.GitHubProfile
	profile.Website = req.Website
	profile.Bio = req.Bio
	profile.Languages = string(languagesJSON)
	profile.TimeZone = req.TimeZone
	profile.ResponseTime = req.ResponseTime

	return fs.freelancerRepo.UpdateFreelancerProfile(ctx, profile)
}

// GetFreelancerProfile get freelancer profile
func (fs *FreelancerService) GetFreelancerProfile(ctx context.Context, req *schema.GetFreelancerProfileReq) (*schema.FreelancerProfileResp, error) {
	profile, exist, err := fs.freelancerRepo.GetFreelancerProfileByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.NotFound(reason.FreelancerProfileNotFound)
	}

	return fs.convertFreelancerProfileToResp(profile), nil
}

// GetFreelancerProfiles get freelancer profiles
func (fs *FreelancerService) GetFreelancerProfiles(ctx context.Context, req *schema.GetFreelancerProfilesReq) (*schema.GetFreelancerProfilesResp, error) {
	profiles, total, err := fs.freelancerRepo.GetFreelancerProfiles(ctx, req)
	if err != nil {
		return nil, err
	}

	var resp []*schema.FreelancerProfileResp
	for _, profile := range profiles {
		resp = append(resp, fs.convertFreelancerProfileToResp(profile))
	}

	return &schema.GetFreelancerProfilesResp{
		Count: int(total),
		List:  resp,
	}, nil
}

// CreateJobPosting create job posting
func (fs *FreelancerService) CreateJobPosting(ctx context.Context, req *schema.CreateJobPostingReq) error {
	// Convert skills to JSON
	skillsJSON, err := json.Marshal(req.Skills)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	posting := &entity.JobPosting{
		UserID:          req.LoginUserID,
		Title:           req.Title,
		Description:     req.Description,
		Budget:          req.Budget,
		Currency:        req.Currency,
		BudgetType:      req.BudgetType,
		Skills:          string(skillsJSON),
		ExperienceLevel: req.ExperienceLevel,
		Duration:        req.Duration,
		Location:        req.Location,
		ContactEmail:    req.ContactEmail,
		ExpiresAt:       time.Unix(req.ExpiresAt, 0),
	}

	return fs.freelancerRepo.CreateJobPosting(ctx, posting)
}

// GetJobPostings get job postings
func (fs *FreelancerService) GetJobPostings(ctx context.Context, req *schema.GetJobPostingsReq) (*schema.GetJobPostingsResp, error) {
	postings, total, err := fs.freelancerRepo.GetJobPostings(ctx, req)
	if err != nil {
		return nil, err
	}

	var resp []*schema.JobPostingResp
	for _, posting := range postings {
		resp = append(resp, fs.convertJobPostingToResp(posting))
	}

	return &schema.GetJobPostingsResp{
		Count: int(total),
		List:  resp,
	}, nil
}

// GetJobPosting get job posting by ID
func (fs *FreelancerService) GetJobPosting(ctx context.Context, id string) (*schema.JobPostingResp, error) {
	posting, exist, err := fs.freelancerRepo.GetJobPostingByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.NotFound(reason.JobPostingNotFound)
	}

	// Increment views count
	go func() {
		ctx := context.Background()
		fs.freelancerRepo.IncrementJobViews(ctx, id)
	}()

	return fs.convertJobPostingToResp(posting), nil
}

// CreateJobApplication create job application
func (fs *FreelancerService) CreateJobApplication(ctx context.Context, req *schema.CreateJobApplicationReq) error {
	// Check if job exists
	_, exist, err := fs.freelancerRepo.GetJobPostingByID(ctx, req.JobID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.NotFound(reason.JobPostingNotFound)
	}

	// Check if user already applied
	applications, err := fs.freelancerRepo.GetJobApplicationsByApplicantID(ctx, req.LoginUserID)
	if err != nil {
		return err
	}
	for _, app := range applications {
		if app.JobID == req.JobID {
			return errors.BadRequest(reason.JobApplicationAlreadyExists)
		}
	}

	application := &entity.JobApplication{
		JobID:        req.JobID,
		ApplicantID:  req.LoginUserID,
		CoverLetter:  req.CoverLetter,
		ProposedRate: req.ProposedRate,
		Currency:     req.Currency,
		Message:      req.Message,
	}

	return fs.freelancerRepo.CreateJobApplication(ctx, application)
}

// HireFreelancer hire freelancer
func (fs *FreelancerService) HireFreelancer(ctx context.Context, req *schema.HireFreelancerReq) (*schema.HireFreelancerResp, error) {
	// Get freelancer user info
	freelancerUser, exist, err := fs.userRepo.GetByUserID(ctx, req.FreelancerUserID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.NotFound(reason.UserNotFound)
	}

	// Get freelancer profile
	freelancerProfile, exist, err := fs.freelancerRepo.GetFreelancerProfileByUserID(ctx, req.FreelancerUserID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.NotFound(reason.FreelancerProfileNotFound)
	}

	// Get current user info
	currentUser, exist, err := fs.userRepo.GetByUserID(ctx, req.LoginUserID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.NotFound(reason.UserNotFound)
	}

	// Get site info
	siteInfo, err := fs.siteInfoService.GetSiteGeneral(ctx)
	if err != nil {
		return nil, err
	}

	// Prepare email
	contactEmail := freelancerUser.EMail
	if freelancerProfile.ContactEmail != "" {
		contactEmail = freelancerProfile.ContactEmail
	}

	// Send hiring email
	subject := req.Subject
	body := fs.generateHiringEmailBody(req.Message, currentUser, freelancerUser, siteInfo.Name)

	fs.emailService.Send(ctx, contactEmail, subject, body)

	return &schema.HireFreelancerResp{
		Success: true,
		Message: "Hiring message sent successfully",
	}, nil
}

// convertFreelancerProfileToResp convert freelancer profile to response
func (fs *FreelancerService) convertFreelancerProfileToResp(profile *entity.FreelancerProfile) *schema.FreelancerProfileResp {
	var skills []string
	json.Unmarshal([]byte(profile.Skills), &skills)

	var portfolio []string
	json.Unmarshal([]byte(profile.Portfolio), &portfolio)

	var preferredProjects []string
	json.Unmarshal([]byte(profile.PreferredProjects), &preferredProjects)

	var languages []string
	json.Unmarshal([]byte(profile.Languages), &languages)

	return &schema.FreelancerProfileResp{
		ID:                profile.ID,
		UserID:            profile.UserID,
		IsAvailable:       profile.IsAvailable,
		HourlyRate:        profile.HourlyRate,
		Currency:          profile.Currency,
		Skills:            skills,
		Experience:        profile.Experience,
		Portfolio:         portfolio,
		Availability:      profile.Availability,
		PreferredProjects: preferredProjects,
		ContactEmail:      profile.ContactEmail,
		LinkedInProfile:   profile.LinkedInProfile,
		GitHubProfile:     profile.GitHubProfile,
		Website:           profile.Website,
		Bio:               profile.Bio,
		BioHTML:           profile.BioHTML,
		Languages:         languages,
		TimeZone:          profile.TimeZone,
		ResponseTime:      profile.ResponseTime,
		CompletedProjects: profile.CompletedProjects,
		ClientSatisfaction: profile.ClientSatisfaction,
		IsVerified:        profile.IsVerified,
		VerificationDate:  profile.VerificationDate.Unix(),
		CreatedAt:         profile.CreatedAt.Unix(),
		UpdatedAt:         profile.UpdatedAt.Unix(),
	}
}

// convertJobPostingToResp convert job posting to response
func (fs *FreelancerService) convertJobPostingToResp(posting *entity.JobPosting) *schema.JobPostingResp {
	var skills []string
	json.Unmarshal([]byte(posting.Skills), &skills)

	return &schema.JobPostingResp{
		ID:               posting.ID,
		UserID:           posting.UserID,
		Title:            posting.Title,
		Description:      posting.Description,
		DescriptionHTML:  posting.DescriptionHTML,
		Budget:           posting.Budget,
		Currency:        posting.Currency,
		BudgetType:       posting.BudgetType,
		Skills:           skills,
		ExperienceLevel:  posting.ExperienceLevel,
		Duration:         posting.Duration,
		Location:         posting.Location,
		Status:           posting.Status,
		ContactEmail:     posting.ContactEmail,
		ApplicationCount: posting.ApplicationCount,
		ViewsCount:       posting.ViewsCount,
		IsActive:         posting.IsActive,
		ExpiresAt:        posting.ExpiresAt.Unix(),
		CreatedAt:        posting.CreatedAt.Unix(),
		UpdatedAt:        posting.UpdatedAt.Unix(),
	}
}

// generateHiringEmailBody generate hiring email body
func (fs *FreelancerService) generateHiringEmailBody(message string, currentUser, freelancerUser *entity.User, siteName string) string {
	return `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Job Opportunity from ` + siteName + `</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
    <div style="max-width: 600px; margin: 0 auto; padding: 20px;">
        <h2 style="color: #2c3e50;">Job Opportunity from ` + siteName + `</h2>
        
        <p>Hello ` + freelancerUser.DisplayName + `,</p>
        
        <p>You have received a job opportunity from <strong>` + currentUser.DisplayName + `</strong> (` + currentUser.Username + `) on ` + siteName + `.</p>
        
        <div style="background-color: #f8f9fa; padding: 15px; border-left: 4px solid #007bff; margin: 20px 0;">
            <h3 style="margin-top: 0; color: #007bff;">Message:</h3>
            <p style="margin-bottom: 0;">` + message + `</p>
        </div>
        
        <p>You can contact them directly at: <a href="mailto:` + currentUser.EMail + `">` + currentUser.EMail + `</a></p>
        
        <p>Best regards,<br>The ` + siteName + ` Team</p>
        
        <hr style="border: none; border-top: 1px solid #eee; margin: 30px 0;">
        <p style="font-size: 12px; color: #666;">
            This message was sent through ` + siteName + `. Please do not reply to this email directly.
        </p>
    </div>
</body>
</html>
`
}
