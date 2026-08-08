.PHONY: run down

# Live-reloading dev server on http://localhost:4000
run:
	set -a; . ./.env; set +a; air

# Kill a stray server still holding port 4000
down:
	@pkill -f 'tmp/web' || true
	@fuser -k 4000/tcp 2>/dev/null || true
