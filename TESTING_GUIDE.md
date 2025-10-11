# Testing Guide for Freelancer Features

This guide provides comprehensive testing approaches for your new freelancer functionality in the Answer project.

## 🧪 Testing Overview

Your new freelancer features include:
- Freelancer profile management (create, read, update, delete)
- Job posting system
- Job application system
- Frontend components for freelancer functionality

## 📋 Testing Methods

### 1. Unit Testing

**Run unit tests:**
```bash
# Run freelancer service tests
go test ./internal/service/freelancer/ -v

# Run all tests
make test
```

**Test coverage:**
```bash
go test ./internal/service/freelancer/ -cover
```

### 2. Integration Testing

**Database integration tests:**
```bash
# Make sure your database is running
# Test with SQLite (default)
./answer

# Or test with PostgreSQL/MySQL
# Update config.yaml with your database settings
```

**Test database migrations:**
```bash
# Check if migrations run successfully
./answer --check-migrations
```

### 3. API Testing

**Using curl (see test_api_examples.sh):**
```bash
# Run the API test examples
./test_api_examples.sh
```

**Using Postman:**
1. Import the API collection (if available)
2. Set up environment variables:
   - `base_url`: http://localhost:9080
   - `api_base`: http://localhost:9080/answer/api/v1
   - `jwt_token`: Your authentication token

**Manual API testing:**
```bash
# Start the application
./answer

# Test in another terminal
curl -X GET http://localhost:9080/answer/api/v1/freelancer/profiles
```

### 4. Docker Testing

**Build and test with Docker:**
```bash
# Build the freelancer-specific image
docker build -f Dockerfile.freelancer -t answer-freelancer .

# Run the container
docker run -p 9080:80 answer-freelancer

# Test with docker-compose
docker-compose up --build
```

**Test different environments:**
```bash
# Test with different Dockerfiles
docker build -f Dockerfile.local -t answer-local .
docker build -f Dockerfile.custom -t answer-custom .
```

### 5. Frontend Testing

**Build and test UI:**
```bash
# Install UI dependencies
cd ui
pnpm install

# Build the frontend
pnpm build

# Test in development mode
pnpm start
```

**Test specific components:**
```bash
# Test the HireMeButton component
cd ui/src/components/HireMeButton
# Run component tests if available
```

### 6. End-to-End Testing

**Complete workflow testing:**
1. Start the application: `./answer`
2. Open browser: http://localhost:9080
3. Register a new user
4. Create a freelancer profile
5. Post a job
6. Apply to the job
7. Test the complete flow

## 🔧 Test Configuration

### Environment Variables
```bash
export ANSWER_DB_DRIVER=sqlite3
export ANSWER_DB_DSN=./data/sqlite3/answer.db
export ANSWER_PORT=9080
```

### Database Setup
```bash
# For SQLite (default)
mkdir -p data/sqlite3
touch data/sqlite3/answer.db

# For PostgreSQL
createdb answer_test
export ANSWER_DB_DRIVER=postgres
export ANSWER_DB_DSN="host=localhost port=5432 user=postgres password=password dbname=answer_test sslmode=disable"
```

## 🚀 Quick Start Testing

**Run the complete test suite:**
```bash
# Run the automated test script
./test_freelancer.sh
```

This script will:
1. Install dependencies
2. Run unit tests
3. Generate wire dependencies
4. Build the application
5. Provide Docker testing options

## 📊 Test Scenarios

### Freelancer Profile Tests
- ✅ Create profile with valid data
- ✅ Create profile with invalid data (should fail)
- ✅ Update existing profile
- ✅ Update non-existent profile (should fail)
- ✅ Get profile by user ID
- ✅ Delete profile

### Job Posting Tests
- ✅ Create job posting
- ✅ Update job posting
- ✅ Get job postings with pagination
- ✅ Search jobs by skills
- ✅ Filter jobs by location/type

### Job Application Tests
- ✅ Apply to job
- ✅ Withdraw application
- ✅ Get applications for a job
- ✅ Update application status

### Frontend Tests
- ✅ HireMeButton component renders
- ✅ Profile form validation
- ✅ Job posting form
- ✅ Application form

## 🐛 Debugging

**Check logs:**
```bash
# Run with verbose logging
./answer --log-level=debug

# Check Docker logs
docker logs <container_id>
```

**Database debugging:**
```bash
# For SQLite
sqlite3 data/sqlite3/answer.db
.tables
.schema freelancer_profile
```

**API debugging:**
```bash
# Use verbose curl
curl -v -X POST http://localhost:9080/answer/api/v1/freelancer/profile \
  -H "Content-Type: application/json" \
  -d '{"test": "data"}'
```

## 📝 Test Data

**Sample freelancer profile:**
```json
{
  "is_available": true,
  "hourly_rate": 50.0,
  "currency": "USD",
  "skills": ["Go", "React", "Docker"],
  "experience": "5+ years experience",
  "portfolio": ["https://github.com/user/project"],
  "availability": "Full-time",
  "preferred_projects": ["Web Development"],
  "contact_email": "freelancer@example.com",
  "bio": "Experienced developer",
  "languages": ["English"],
  "time_zone": "UTC-5",
  "response_time": "Within 24 hours"
}
```

**Sample job posting:**
```json
{
  "title": "Go Developer Needed",
  "description": "Looking for experienced Go developer",
  "budget": 10000.0,
  "currency": "USD",
  "budget_type": "fixed",
  "skills": ["Go", "PostgreSQL"],
  "experience_level": "senior",
  "duration": "3 months",
  "location": "remote"
}
```

## ✅ Success Criteria

Your testing is successful when:
- [ ] All unit tests pass
- [ ] API endpoints respond correctly
- [ ] Database migrations work
- [ ] Frontend components render properly
- [ ] Complete user workflows function
- [ ] Docker containers build and run
- [ ] No critical errors in logs

## 🆘 Troubleshooting

**Common issues:**
1. **Database connection errors**: Check database configuration
2. **API 404 errors**: Verify routes are registered
3. **Build failures**: Check Go module dependencies
4. **Docker issues**: Ensure Docker is running
5. **Frontend build errors**: Check Node.js/pnpm installation

**Get help:**
- Check application logs
- Verify configuration files
- Test with minimal data first
- Use debug mode for detailed output
