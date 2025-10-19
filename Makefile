
.PHONY: prep prep-backend prep-frontend build build-backend build-frontend clean docker-build docker-run run-backend run-frontend run test-backend

# Preparation steps
prep-backend:
	@echo "Prepare backend"
	cd ./backend && make prep

prep-frontend:
	@echo "Prepare frontend"
	cd ./frontend && rm -rf node_modules && npm ci

prep: prep-backend prep-frontend

# Build steps
build-backend:
	@echo "Build backend binary"
	cd ./backend && make build

build-frontend:
	@echo "Build frontend"
	cd ./frontend && npm run build

build: build-backend build-frontend

clean:
	@echo "Clean backend and frontend artifacts"
	@cd ./backend && make clean
	@cd ./frontend && rm -rf .next node_modules/.cache

docker-build:
	@echo "Building Docker image"
	docker build -t route-portal:latest .

docker-run:
	@echo "Run Docker image"
	docker run --name route-portal -p 3000:3000 -p 8080:8080 route-portal

# Run steps for local development
run-backend:
	@echo "Run backend locally"
	cd ./backend && make run

run-frontend:
	@echo "Run frontend locally"
	cd ./frontend && npm run dev

run: run-backend run-frontend

# Test steps
test-backend:
	@echo "Run backend unit tests"
	cd ./backend && make test
