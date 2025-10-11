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
	"testing"

	"github.com/apache/answer/internal/entity"
	"github.com/apache/answer/internal/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockFreelancerRepo is a mock implementation of FreelancerRepo
type MockFreelancerRepo struct {
	mock.Mock
}

func (m *MockFreelancerRepo) CreateFreelancerProfile(ctx context.Context, profile *entity.FreelancerProfile) error {
	args := m.Called(ctx, profile)
	return args.Error(0)
}

func (m *MockFreelancerRepo) UpdateFreelancerProfile(ctx context.Context, profile *entity.FreelancerProfile) error {
	args := m.Called(ctx, profile)
	return args.Error(0)
}

func (m *MockFreelancerRepo) GetFreelancerProfileByUserID(ctx context.Context, userID string) (*entity.FreelancerProfile, bool, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*entity.FreelancerProfile), args.Bool(1), args.Error(2)
}

func (m *MockFreelancerRepo) GetFreelancerProfile(ctx context.Context, id string) (*entity.FreelancerProfile, bool, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.FreelancerProfile), args.Bool(1), args.Error(2)
}

func (m *MockFreelancerRepo) DeleteFreelancerProfile(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockFreelancerRepo) GetFreelancerProfiles(ctx context.Context, page, pageSize int) ([]*entity.FreelancerProfile, int64, error) {
	args := m.Called(ctx, page, pageSize)
	return args.Get(0).([]*entity.FreelancerProfile), args.Get(1).(int64), args.Error(2)
}

// MockUserRepo is a mock implementation of UserRepo
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetUserByID(ctx context.Context, userID string) (*entity.User, bool, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*entity.User), args.Bool(1), args.Error(2)
}

// MockEmailService is a mock implementation of EmailService
type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) SendAndSaveCode(ctx context.Context, toEmail, subject, code, notificationType string) error {
	args := m.Called(ctx, toEmail, subject, code, notificationType)
	return args.Error(0)
}

// MockSiteInfoService is a mock implementation of SiteInfoService
type MockSiteInfoService struct {
	mock.Mock
}

func (m *MockSiteInfoService) GetSiteGeneral(ctx context.Context) (*schema.SiteGeneralResp, error) {
	args := m.Called(ctx)
	return args.Get(0).(*schema.SiteGeneralResp), args.Error(1)
}

func TestCreateFreelancerProfile(t *testing.T) {
	ctx := context.Background()
	
	// Test case 1: Successful creation
	t.Run("successful_creation", func(t *testing.T) {
		mockRepo := new(MockFreelancerRepo)
		mockUserRepo := new(MockUserRepo)
		mockEmailService := new(MockEmailService)
		mockSiteInfoService := new(MockSiteInfoService)
		
		service := NewFreelancerService(mockRepo, mockUserRepo, mockEmailService, mockSiteInfoService)
		
		req := &schema.CreateFreelancerProfileReq{
			LoginUserID:      "user123",
			IsAvailable:      true,
			HourlyRate:       50.0,
			Currency:         "USD",
			Skills:           []string{"Go", "React", "Docker"},
			Experience:       "5 years of experience",
			Portfolio:        []string{"project1.com", "project2.com"},
			Availability:     "Full-time",
			PreferredProjects: []string{"Web Development", "API Development"},
			ContactEmail:     "freelancer@example.com",
			LinkedInProfile:  "linkedin.com/in/freelancer",
			GitHubProfile:    "github.com/freelancer",
			Website:          "freelancer.com",
			Bio:              "Experienced developer",
			Languages:        []string{"English", "Spanish"},
			TimeZone:         "UTC-5",
			ResponseTime:     "Within 24 hours",
		}
		
		// Mock that profile doesn't exist
		mockRepo.On("GetFreelancerProfileByUserID", ctx, "user123").Return((*entity.FreelancerProfile)(nil), false, nil)
		mockRepo.On("CreateFreelancerProfile", ctx, mock.AnythingOfType("*entity.FreelancerProfile")).Return(nil)
		
		err := service.CreateFreelancerProfile(ctx, req)
		
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
	
	// Test case 2: Profile already exists
	t.Run("profile_already_exists", func(t *testing.T) {
		mockRepo := new(MockFreelancerRepo)
		mockUserRepo := new(MockUserRepo)
		mockEmailService := new(MockEmailService)
		mockSiteInfoService := new(MockSiteInfoService)
		
		service := NewFreelancerService(mockRepo, mockUserRepo, mockEmailService, mockSiteInfoService)
		
		req := &schema.CreateFreelancerProfileReq{
			LoginUserID: "user123",
		}
		
		existingProfile := &entity.FreelancerProfile{ID: "1", UserID: "user123"}
		mockRepo.On("GetFreelancerProfileByUserID", ctx, "user123").Return(existingProfile, true, nil)
		
		err := service.CreateFreelancerProfile(ctx, req)
		
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateFreelancerProfile(t *testing.T) {
	ctx := context.Background()
	
	t.Run("successful_update", func(t *testing.T) {
		mockRepo := new(MockFreelancerRepo)
		mockUserRepo := new(MockUserRepo)
		mockEmailService := new(MockEmailService)
		mockSiteInfoService := new(MockSiteInfoService)
		
		service := NewFreelancerService(mockRepo, mockUserRepo, mockEmailService, mockSiteInfoService)
		
		req := &schema.UpdateFreelancerProfileReq{
			ID:            "profile123",
			LoginUserID:   "user123",
			IsAvailable:   false,
			HourlyRate:    75.0,
			Currency:      "EUR",
			Skills:        []string{"Python", "Django"},
			Experience:    "Updated experience",
			Portfolio:     []string{"newproject.com"},
			Availability:  "Part-time",
			PreferredProjects: []string{"Data Science"},
			ContactEmail:  "newemail@example.com",
			LinkedInProfile: "linkedin.com/in/newprofile",
			GitHubProfile: "github.com/newprofile",
			Website:       "newwebsite.com",
			Bio:           "Updated bio",
			Languages:     []string{"English", "French"},
			TimeZone:      "UTC+1",
			ResponseTime:  "Within 12 hours",
		}
		
		existingProfile := &entity.FreelancerProfile{
			ID:     "profile123",
			UserID: "user123",
		}
		
		mockRepo.On("GetFreelancerProfileByUserID", ctx, "user123").Return(existingProfile, true, nil)
		mockRepo.On("UpdateFreelancerProfile", ctx, mock.AnythingOfType("*entity.FreelancerProfile")).Return(nil)
		
		err := service.UpdateFreelancerProfile(ctx, req)
		
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("profile_not_found", func(t *testing.T) {
		mockRepo := new(MockFreelancerRepo)
		mockUserRepo := new(MockUserRepo)
		mockEmailService := new(MockEmailService)
		mockSiteInfoService := new(MockSiteInfoService)
		
		service := NewFreelancerService(mockRepo, mockUserRepo, mockEmailService, mockSiteInfoService)
		
		req := &schema.UpdateFreelancerProfileReq{
			ID:          "profile123",
			LoginUserID: "user123",
		}
		
		mockRepo.On("GetFreelancerProfileByUserID", ctx, "user123").Return((*entity.FreelancerProfile)(nil), false, nil)
		
		err := service.UpdateFreelancerProfile(ctx, req)
		
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
