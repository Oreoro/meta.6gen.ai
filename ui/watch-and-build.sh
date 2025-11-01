#!/bin/bash
# Alternative: Watch for changes and rebuild automatically
# Use this if you prefer to keep using the production build

cd "$(dirname "$0")"

echo "=== UI Watch & Build Mode ==="
echo "Watching for changes in ui/src..."
echo "Rebuilding automatically when files change..."
echo ""
echo "Press Ctrl+C to stop"
echo ""

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
    echo "Installing dependencies..."
    pnpm install
fi

# Install chokidar-cli if not available
if ! command -v chokidar &> /dev/null; then
    echo "Installing chokidar-cli for file watching..."
    pnpm add -D chokidar-cli
fi

# Watch for changes and rebuild
echo "Watching for changes..."
chokidar "src/**/*.{ts,tsx,scss,css}" "public/**/*" -c "echo 'Change detected, rebuilding...' && pnpm build && echo '✅ Rebuild complete! Refresh your browser.'"

