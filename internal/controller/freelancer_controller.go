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

package controller

import (
	"github.com/apache/answer/internal/base/handler"
	"github.com/apache/answer/internal/base/middleware"
	"github.com/apache/answer/internal/base/reason"
	"github.com/apache/answer/internal/schema"
	"github.com/apache/answer/internal/service/freelancer"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// FreelancerController freelancer controller
type FreelancerController struct {
	freelancerService *freelancer.FreelancerService
}

// NewFreelancerController new freelancer controller
func NewFreelancerController(freelancerService *freelancer.FreelancerService) *FreelancerController {
	return &FreelancerController{
		freelancerService: freelancerService,
	}
}

// CreateFreelancerProfile godoc
// @Summary Create freelancer profile
// @Description Create freelancer profile
// @Tags Freelancer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.CreateFreelancerProfileReq true "CreateFreelancerProfile"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/freelancer/profile [post]
func (fc *FreelancerController) CreateFreelancerProfile(ctx *gin.Context) {
	req := &schema.CreateFreelancerProfileReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.LoginUserID = middleware.GetLoginUserIDFromContext(ctx)
	err := fc.freelancerService.CreateFreelancerProfile(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UpdateFreelancerProfile godoc
// @Summary Update freelancer profile
// @Description Update freelancer profile
// @Tags Freelancer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.UpdateFreelancerProfileReq true "UpdateFreelancerProfile"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/freelancer/profile [put]
func (fc *FreelancerController) UpdateFreelancerProfile(ctx *gin.Context) {
	req := &schema.UpdateFreelancerProfileReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.LoginUserID = middleware.GetLoginUserIDFromContext(ctx)
	err := fc.freelancerService.UpdateFreelancerProfile(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetFreelancerProfile godoc
// @Summary Get freelancer profile
// @Description Get freelancer profile
// @Tags Freelancer
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Success 200 {object} handler.RespBody{data=schema.FreelancerProfileResp}
// @Router /answer/api/v1/freelancer/profile [get]
func (fc *FreelancerController) GetFreelancerProfile(ctx *gin.Context) {
	req := &schema.GetFreelancerProfileReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := fc.freelancerService.GetFreelancerProfile(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// GetFreelancerProfiles godoc
// @Summary Get freelancer profiles
// @Description Get freelancer profiles
// @Tags Freelancer
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Param skills query string false "skills"
// @Param location query string false "location"
// @Param min_rate query number false "min_rate"
// @Param max_rate query number false "max_rate"
// @Param currency query string false "currency"
// @Success 200 {object} handler.RespBody{data=schema.GetFreelancerProfilesResp}
// @Router /answer/api/v1/freelancer/profiles [get]
func (fc *FreelancerController) GetFreelancerProfiles(ctx *gin.Context) {
	req := &schema.GetFreelancerProfilesReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := fc.freelancerService.GetFreelancerProfiles(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// CreateJobPosting godoc
// @Summary Create job posting
// @Description Create job posting
// @Tags Job
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.CreateJobPostingReq true "CreateJobPosting"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/job/posting [post]
func (fc *FreelancerController) CreateJobPosting(ctx *gin.Context) {
	req := &schema.CreateJobPostingReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.LoginUserID = middleware.GetLoginUserIDFromContext(ctx)
	err := fc.freelancerService.CreateJobPosting(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetJobPostings godoc
// @Summary Get job postings
// @Description Get job postings
// @Tags Job
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Param skills query string false "skills"
// @Param location query string false "location"
// @Param min_budget query number false "min_budget"
// @Param max_budget query number false "max_budget"
// @Param currency query string false "currency"
// @Param status query string false "status"
// @Success 200 {object} handler.RespBody{data=schema.GetJobPostingsResp}
// @Router /answer/api/v1/job/postings [get]
func (fc *FreelancerController) GetJobPostings(ctx *gin.Context) {
	req := &schema.GetJobPostingsReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := fc.freelancerService.GetJobPostings(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// GetJobPosting godoc
// @Summary Get job posting
// @Description Get job posting by ID
// @Tags Job
// @Accept json
// @Produce json
// @Param id path string true "job_id"
// @Success 200 {object} handler.RespBody{data=schema.JobPostingResp}
// @Router /answer/api/v1/job/posting/{id} [get]
func (fc *FreelancerController) GetJobPosting(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), nil)
		return
	}

	resp, err := fc.freelancerService.GetJobPosting(ctx, id)
	handler.HandleResponse(ctx, err, resp)
}

// CreateJobApplication godoc
// @Summary Create job application
// @Description Create job application
// @Tags Job
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.CreateJobApplicationReq true "CreateJobApplication"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/job/application [post]
func (fc *FreelancerController) CreateJobApplication(ctx *gin.Context) {
	req := &schema.CreateJobApplicationReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.LoginUserID = middleware.GetLoginUserIDFromContext(ctx)
	err := fc.freelancerService.CreateJobApplication(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// HireFreelancer godoc
// @Summary Hire freelancer
// @Description Send hiring message to freelancer
// @Tags Freelancer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.HireFreelancerReq true "HireFreelancer"
// @Success 200 {object} handler.RespBody{data=schema.HireFreelancerResp}
// @Router /answer/api/v1/freelancer/hire [post]
func (fc *FreelancerController) HireFreelancer(ctx *gin.Context) {
	req := &schema.HireFreelancerReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.LoginUserID = middleware.GetLoginUserIDFromContext(ctx)
	resp, err := fc.freelancerService.HireFreelancer(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
