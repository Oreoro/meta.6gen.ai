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

package migrations

import (
	"context"
	"fmt"

	"github.com/apache/answer/internal/entity"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

// MigrateFreelancerTables create freelancer related tables
func MigrateFreelancerTables(ctx context.Context, db *xorm.Engine) error {
	log.Info("Creating freelancer profile table...")
	err := db.Context(ctx).Sync2(new(entity.FreelancerProfile))
	if err != nil {
		return fmt.Errorf("failed to create freelancer_profile table: %w", err)
	}

	log.Info("Creating job posting table...")
	err = db.Context(ctx).Sync2(new(entity.JobPosting))
	if err != nil {
		return fmt.Errorf("failed to create job_posting table: %w", err)
	}

	log.Info("Creating job application table...")
	err = db.Context(ctx).Sync2(new(entity.JobApplication))
	if err != nil {
		return fmt.Errorf("failed to create job_application table: %w", err)
	}

	log.Info("Freelancer tables created successfully")
	return nil
}
