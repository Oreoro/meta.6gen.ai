# Hot Reload Setup for UI Development

You have two options for hot reloading UI changes:

## Option 1: Development Server (Recommended) 🚀

This runs React's development server with built-in hot module replacement (HMR).

### Setup:
1. **Start your backend:**
   ```bash
   docker-compose up -d
   ```

2. **Start the UI dev server:**
   ```bash
   ./start-ui-dev.sh
   ```

3. **Access the app:**
   - Open http://localhost:3000 (dev server with hot reload)
   - API calls are automatically proxied to http://localhost:8081

### How it works:
- React dev server runs on port 3000
- Changes to `ui/src/` files trigger automatic rebuilds
- Browser automatically refreshes (or uses HMR for instant updates)
- API requests are proxied to your backend container

### Advantages:
- ✅ True hot reload (HMR)
- ✅ Fast feedback loop
- ✅ Source maps for debugging
- ✅ Error overlay in browser

---

## Option 2: Watch & Rebuild Mode

This automatically rebuilds when files change, then you refresh the browser.

### Setup:
1. **Start your backend:**
   ```bash
   docker-compose up -d
   ```

2. **Start the watch script:**
   ```bash
   cd ui && ./watch-and-build.sh
   ```

3. **Access the app:**
   - Open http://localhost:8081
   - Refresh browser after rebuild completes

### How it works:
- Watches `ui/src/` for changes
- Automatically runs `pnpm build` when files change
- Rebuilds go to `ui/build/` which is mounted in Docker
- Container serves the new build immediately

### Advantages:
- ✅ Uses production build (same as deployed)
- ✅ No separate dev server needed
- ✅ Works with existing Docker setup

---

## Comparison

| Feature | Dev Server | Watch & Rebuild |
|---------|-----------|-----------------|
| Speed | ⚡ Fast (HMR) | 🐢 Slower (full rebuild) |
| Auto-refresh | ✅ Instant | ⚠️ Manual refresh needed |
| Production-like | ❌ Dev mode | ✅ Production build |
| Setup complexity | Simple | Simple |
| Port | 3000 | 8081 |

## Recommended Workflow

**For active development:** Use Option 1 (Dev Server)
- Fastest feedback
- Best developer experience

**For testing production builds:** Use Option 2 (Watch & Rebuild)
- See exactly what users will see
- Test production optimizations

---

## Troubleshooting

### Dev Server Issues:
- **Port 3000 already in use:** Change port in `ui/package.json` or kill the process
- **API calls failing:** Ensure backend is running on port 8081
- **Hot reload not working:** Clear browser cache or restart dev server

### Watch & Rebuild Issues:
- **Changes not detected:** Install chokidar-cli: `pnpm add -D chokidar-cli`
- **Container not seeing changes:** Check volume mount in docker-compose.yaml
- **Build errors:** Check console output for TypeScript/ESLint errors

