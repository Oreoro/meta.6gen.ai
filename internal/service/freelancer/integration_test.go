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
	"os"
	"testing"

	"github.com/apache/answer/internal/entity"
	"github.com/apache/answer/internal/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// IntegrationTestSuite provides integration testing for freelancer features
type IntegrationTestSuite struct {
	service *FreelancerService
	ctx     context.Context
}

// SetupIntegrationTest sets up the integration test environment
func SetupIntegrationTest(t *testing.T) *IntegrationTestSuite {
	// Skip integration tests if not explicitly requested
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Integration tests skipped. Set RUN_INTEGRATION_TESTS=true to run.")
	}

	// This would normally set up a real database connection
	// For now, we'll use mocks but structure it for real integration testing
	suite := &IntegrationTestSuite{
		ctx: context.Background(),
	}

	// TODO: Set up real database connection for integration tests
	// This requires:
	// 1. Database connection setup
	// 2. Migration running
	// 3. Test data seeding
	// 4. Cleanup after tests

	return suite
}

func TestFreelancerProfileIntegration(t *testing.T) {
	suite := SetupIntegrationTest(t)

	t.Run("create_and_retrieve_profile", func(t *testing.T) {
		// Test creating a freelancer profile and then retrieving it
		req := &schema.CreateFreelancerProfileReq{
			LoginUserID:      "test_user_123",
			IsAvailable:      true,
			HourlyRate:       50.0,
			Currency:         "USD",
			Skills:           []string{"Go", "React", "Docker"},
			Experience:       "5+ years of experience",
			Portfolio:        []string{"https://github.com/test/project1"},
			Availability:     "Full-time",
			PreferredProjects: []string{"Web Development"},
			ContactEmail:     "test@example.com",
			LinkedInProfile:  "https://linkedin.com/in/test",
			GitHubProfile:    "https://github.com/test",
			Website:          "https://test.com",
			Bio:              "Test freelancer profile",
			Languages:        []string{"English"},
			TimeZone:         "UTC-5",
			ResponseTime:     "Within 24 hours",
		}

		// Create profile
		err := suite.service.CreateFreelancerProfile(suite.ctx, req)
		require.NoError(t, err, "Failed to create freelancer profile")

		// Retrieve profile
		profile, exist, err := suite.service.GetFreelancerProfile(suite.ctx, "test_user_123")
		require.NoError(t, err, "Failed to retrieve freelancer profile")
		require.True(t, exist, "Freelancer profile should exist")
		require.NotNil(t, profile, "Freelancer profile should not be nil")

		// Verify profile data
		assert.Equal(t, req.LoginUserID, profile.UserID)
		assert.Equal(t, req.IsAvailable, profile.IsAvailable)
		assert.Equal(t, req.HourlyRate, profile.HourlyRate)
		assert.Equal(t, req.Currency, profile.Currency)
		assert.Equal(t, req.Experience, profile.Experience)
		assert.Equal(t, req.Availability, profile.Availability)
		assert.Equal(t, req.ContactEmail, profile.ContactEmail)
		assert.Equal(t, req.LinkedInProfile, profile.LinkedInProfile)
		assert.Equal(t, req.GitHubProfile, profile.GitHubProfile)
		assert.Equal(t, req.Website, profile.Website)
		assert.Equal(t, req.Bio, profile.Bio)
		assert.Equal(t, req.TimeZone, profile.TimeZone)
		assert.Equal(t, req.ResponseTime, profile.ResponseTime)
	})

	t.Run("update_profile", func(t *testing.T) {
		// Test updating an existing freelancer profile
		updateReq := &schema.UpdateFreelancerProfileReq{
			ID:            "profile_id_here", // This would be the actual ID from the created profile
			LoginUserID:   "test_user_123",
			IsAvailable:   false,
			HourlyRate:    75.0,
			Currency:      "EUR",
			Skills:        []string{"Python", "Django"},
			Experience:    "Updated experience",
			Portfolio:     []string{"https://newproject.com"},
			Availability:  "Part-time",
			PreferredProjects: []string{"Data Science"},
			ContactEmail:  "newemail@example.com",
			LinkedInProfile: "https://linkedin.com/in/newprofile",
			GitHubProfile: "https://github.com/newprofile",
			Website:       "https://newwebsite.com",
			Bio:           "Updated bio",
			Languages:     []string{"English", "French"},
			TimeZone:      "UTC+1",
			ResponseTime:  "Within 12 hours",
		}

		err := suite.service.UpdateFreelancerProfile(suite.ctx, updateReq)
		require.NoError(t, err, "Failed to update freelancer profile")

		// Verify the update
		profile, exist, err := suite.service.GetFreelancerProfile(suite.ctx, "test_user_123")
		require.NoError(t, err, "Failed to retrieve updated freelancer profile")
		require.True(t, exist, "Updated freelancer profile should exist")

		assert.Equal(t, updateReq.IsAvailable, profile.IsAvailable)
		assert.Equal(t, updateReq.HourlyRate, profile.HourlyRate)
		assert.Equal(t, updateReq.Currency, profile.Currency)
		assert.Equal(t, updateReq.Experience, profile.Experience)
		assert.Equal(t, updateReq.Availability, profile.Availability)
		assert.Equal(t, updateReq.ContactEmail, profile.ContactEmail)
	})

	t.Run("delete_profile", func(t *testing.T) {
		// Test deleting a freelancer profile
		err := suite.service.DeleteFreelancerProfile(suite.ctx, "profile_id_here")
		require.NoError(t, err, "Failed to delete freelancer profile")

		// Verify deletion
		_, exist, err := suite.service.GetFreelancerProfile(suite.ctx, "test_user_123")
		require.NoError(t, err, "Error should not occur when checking deleted profile")
		assert.False(t, exist, "Freelancer profile should not exist after deletion")
	})
}

// Helper function to run integration tests
func RunIntegrationTests() {
	// This function can be called from a separate test runner
	// Set RUN_INTEGRATION_TESTS=true before running
	os.Setenv("RUN_INTEGRATION_TESTS", "true")
}
