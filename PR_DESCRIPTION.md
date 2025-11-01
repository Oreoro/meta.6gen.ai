# Pull Request Description

## Title
Fix custom UI serving and add hot reload support

## Summary
This PR fixes custom UI serving issues and adds hot reload support for development.

## Changes Made

### 🔧 Bug Fixes
- **Fixed custom UI serving**: Updated `internal/router/ui.go` to use `os.ReadFile()` instead of template rendering when `ANSWER_STATIC_PATH` is set
- **Fixed database permissions**: Added `init-db.sql` to grant proper MySQL permissions  
- **Fixed volume permissions**: Added entrypoint script to fix Docker volume permissions on startup
- **Fixed port binding**: Added `INSTALL_PORT` and `SITE_ADDR` env vars to use port 8080 (non-root friendly)

### ✨ New Features
- **Hot reload support**: Added `start-ui-dev.sh` for React dev server with HMR
- **Auto-rebuild script**: Added `watch-and-build.sh` for automatic rebuilds on file changes
- **Verification script**: Added `verify_custom_ui.sh` to verify custom UI is being served correctly

### 📝 Documentation
- Added `HOT_RELOAD_SETUP.md` with hot reload instructions
- Added `SYSTEM_STATUS.md` with system verification report

### 🐛 Merge Conflict Fixes
- Fixed merge conflicts in `Comment/index.tsx` and `HireMeButton/index.tsx`

## Files Changed
- `Dockerfile` - Added verification script
- `docker-compose.yaml` - Added ANSWER_STATIC_PATH, port configs, permission fixes
- `internal/router/ui.go` - Fixed UI serving with os.ReadFile()
- `init-db.sql` - Database permissions fix
- `start-ui-dev.sh` - Hot reload development server
- `verify_custom_ui.sh` - UI verification script
- `ui/watch-and-build.sh` - Auto-rebuild watch script
- `HOT_RELOAD_SETUP.md` - Documentation
- `SYSTEM_STATUS.md` - System status report
- Various UI component fixes

## Testing
- ✅ Code compiles without errors
- ✅ All scripts have valid syntax
- ✅ Docker-compose configuration verified
- ✅ Custom UI build verified (contains HireMeButton and earthly theme)

## Usage
After merge, rebuild the image:
```bash
docker build -f Dockerfile -t oreoro/answer:latest .
docker-compose down && docker-compose up -d
```

For hot reload during development:
```bash
./start-ui-dev.sh
```

## Related Issues
Fixes custom UI not being served when ANSWER_STATIC_PATH is set.

