#!/bin/bash
# Start UI development server with hot reload
# This allows you to see UI changes immediately without rebuilding

echo "=== Starting UI Development Server ==="
echo ""
echo "The dev server will run on http://localhost:3000"
echo "API requests will be proxied to http://localhost:8081"
echo ""
echo "⚠️  Make sure your backend is running: docker-compose up -d"
echo ""
echo "Press Ctrl+C to stop"
echo ""

cd "$(dirname "$0")/ui" || exit 1

# Set API URL for development
export REACT_APP_API_URL=http://localhost:8081
export REACT_APP_BASE_URL=
export TSC_COMPILE_ON_ERROR=true
export ESLINT_NO_DEV_ERRORS=true

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
    echo "Installing dependencies..."
    pnpm install
fi

# Start the development server
pnpm start

