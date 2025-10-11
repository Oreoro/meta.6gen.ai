# How to Test Your New Freelancer Changes

## 🚀 Quick Start

**Run all tests at once:**
```bash
./run_all_tests.sh
```

This comprehensive script will test everything automatically!

## 📋 Individual Testing Methods

### 1. **Unit Testing**
```bash
# Test freelancer service specifically
go test ./internal/service/freelancer/ -v

# Test everything
make test
```

### 2. **Build and Run**
```bash
# Generate dependencies and build
make generate
make build

# Run the application
./answer
```

### 3. **API Testing**
```bash
# View API test examples
./test_api_examples.sh

# Test manually (after starting the app)
curl http://localhost:9080/answer/api/v1/freelancer/profiles
```

### 4. **Docker Testing**
```bash
# Test with Docker
./test_docker.sh

# Or manually
docker build -f Dockerfile.freelancer -t answer-freelancer .
docker run -p 9080:80 answer-freelancer
```

### 5. **Frontend Testing**
```bash
# Test frontend components
./test_frontend.sh

# Or manually
cd ui
pnpm install
pnpm build
pnpm start
```

## 🧪 What Gets Tested

### Backend Features
- ✅ Freelancer profile creation/update/retrieval
- ✅ Job posting system
- ✅ Job application system
- ✅ Database migrations
- ✅ API endpoints
- ✅ Service layer logic

### Frontend Features
- ✅ HireMeButton component
- ✅ Freelancer profile forms
- ✅ Job posting interface
- ✅ Application forms
- ✅ UI build process

### Infrastructure
- ✅ Docker container builds
- ✅ Database connectivity
- ✅ API routing
- ✅ Configuration files

## 🔧 Testing Scenarios

### Scenario 1: Complete User Workflow
1. Start the application: `./answer`
2. Open browser: http://localhost:9080
3. Register a new user
4. Create a freelancer profile
5. Post a job
6. Apply to the job
7. Test the complete flow

### Scenario 2: API Testing
1. Start the application: `./answer`
2. Get authentication token
3. Test API endpoints with curl:
   ```bash
   curl -X POST http://localhost:9080/answer/api/v1/freelancer/profile \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -d '{"is_available": true, "hourly_rate": 50.0, ...}'
   ```

### Scenario 3: Docker Testing
1. Build the container: `docker build -f Dockerfile.freelancer -t answer-freelancer .`
2. Run the container: `docker run -p 9080:80 answer-freelancer`
3. Test the application: `curl http://localhost:9080`

## 📊 Test Data Examples

### Sample Freelancer Profile
```json
{
  "is_available": true,
  "hourly_rate": 50.0,
  "currency": "USD",
  "skills": ["Go", "React", "Docker"],
  "experience": "5+ years of experience",
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

### Sample Job Posting
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

## 🐛 Troubleshooting

### Common Issues

**1. Database Connection Errors**
```bash
# Check database file
ls -la data/sqlite3/answer.db

# Check configuration
cat configs/config.yaml
```

**2. API 404 Errors**
- Verify routes are registered in `internal/router/answer_api_router.go`
- Check if the application is running on the correct port

**3. Build Failures**
```bash
# Clean and rebuild
make clean
make generate
make build
```

**4. Docker Issues**
```bash
# Check Docker status
docker ps
docker logs <container_id>

# Clean up
docker system prune -f
```

**5. Frontend Issues**
```bash
# Clean node_modules and reinstall
cd ui
rm -rf node_modules
pnpm install
pnpm build
```

## ✅ Success Criteria

Your testing is successful when:
- [ ] All unit tests pass
- [ ] Application builds without errors
- [ ] API endpoints respond correctly
- [ ] Database migrations work
- [ ] Frontend components render
- [ ] Docker containers build and run
- [ ] Complete user workflows function

## 🎯 Next Steps After Testing

1. **Deploy to staging environment**
2. **Run performance tests**
3. **Test with real user data**
4. **Monitor application logs**
5. **Deploy to production**

## 📞 Getting Help

If you encounter issues:
1. Check the application logs
2. Verify configuration files
3. Test with minimal data first
4. Use debug mode for detailed output
5. Check the comprehensive testing guide: `TESTING_GUIDE.md`

---

**Happy Testing! 🚀**

Your freelancer features are ready to be tested and deployed!
