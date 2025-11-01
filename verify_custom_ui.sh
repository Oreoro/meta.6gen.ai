#!/bin/bash
# Verification script for custom UI after rebuild

echo "=== Custom UI Verification Script ==="
echo ""

# Check 1: Image exists
echo "1. Checking Docker image..."
if docker images | grep -q "oreoro/answer.*latest"; then
    echo "   ✅ Image exists"
    docker images | grep "oreoro/answer.*latest"
else
    echo "   ❌ Image not found"
    exit 1
fi

# Check 2: Container is running
echo ""
echo "2. Checking container status..."
if docker-compose ps | grep -q "answer.*Up"; then
    echo "   ✅ Container is running"
else
    echo "   ⚠️  Container not running. Start with: docker-compose up -d"
fi

# Check 3: ANSWER_STATIC_PATH is set
echo ""
echo "3. Checking environment variable..."
if docker-compose exec answer sh -c 'echo $ANSWER_STATIC_PATH' 2>/dev/null | grep -q "/data/ui"; then
    echo "   ✅ ANSWER_STATIC_PATH is set to /data/ui"
else
    echo "   ❌ ANSWER_STATIC_PATH not set correctly"
fi

# Check 4: UI files are mounted
echo ""
echo "4. Checking mounted UI files..."
if docker-compose exec answer test -f /data/ui/index.html 2>/dev/null; then
    echo "   ✅ index.html exists in container"
    JS_FILE=$(docker-compose exec answer sh -c 'ls /data/ui/static/js/main*.js 2>/dev/null | head -1')
    if [ -n "$JS_FILE" ]; then
        echo "   ✅ Main JS file: $(basename $JS_FILE)"
        if echo "$JS_FILE" | grep -q "main.37a084e1.js"; then
            echo "   ✅ Correct custom build (main.37a084e1.js)"
        else
            echo "   ⚠️  Different JS file - may be old build"
        fi
    fi
else
    echo "   ❌ UI files not mounted correctly"
fi

# Check 5: Server responds with HTML
echo ""
echo "5. Checking HTTP response..."
RESPONSE=$(curl -s http://localhost:8081 2>/dev/null)
if [ -n "$RESPONSE" ] && echo "$RESPONSE" | grep -q "main.37a084e1.js"; then
    echo "   ✅ Server returns HTML with custom main.js"
    echo "   ✅ Custom UI is being served!"
elif [ -n "$RESPONSE" ]; then
    JS_FILE=$(echo "$RESPONSE" | grep -o 'main\.[^"]*\.js' | head -1)
    echo "   ⚠️  Server returns HTML but with: $JS_FILE"
    if [ "$JS_FILE" != "main.37a084e1.js" ]; then
        echo "   ⚠️  This is NOT the custom build"
    fi
else
    echo "   ❌ Server returns empty response"
    echo "   This means the code fix may not be working"
fi

# Check 6: Custom features in build
echo ""
echo "6. Checking for custom features in build..."
if docker-compose exec answer sh -c 'grep -q "HireMeButton\|FreelancerProvider" /data/ui/static/js/main*.js 2>/dev/null'; then
    echo "   ✅ HireMeButton code found in JS"
else
    echo "   ⚠️  HireMeButton code not found"
fi

if docker-compose exec answer sh -c 'grep -q "brand-green\|#2f6b3a" /data/ui/static/css/main*.css 2>/dev/null'; then
    echo "   ✅ Earthly green theme found in CSS"
else
    echo "   ⚠️  Theme colors not found"
fi

echo ""
echo "=== Verification Complete ==="
echo ""
echo "To test in browser:"
echo "1. Open http://localhost:8081"
echo "2. Hard refresh (Cmd+Shift+R / Ctrl+Shift+R)"
echo "3. Look for:"
echo "   - Green theme colors (#2f6b3a)"
echo "   - '💼 Hire Me' buttons on comments"

