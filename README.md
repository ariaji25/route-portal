# Route Management Portal

A fullstack project built with golang and NextJS for api routes management

## Tech Stack

**Backend (Go):**
- Go 1.24 with standard library (`net/http`)
- YAML-based persistent storage

**Frontend (TypeScript/NextJs):**
- NextJs 14 with App Router

**DevOps:**
- Multi-stage Docker build
- Github Actions

## Project Structure

```
route-portal/
├── backend/                   # Go backend service
│   ├── cmd/api/               # Application entry point
│   ├── domain/                # Domain models and errors
│   ├── internal/              # Private application code
│   │   └── route/
│   │       ├── delivery/http/ # HTTP handlers
│   │       ├── repository/    # Data persistence layer
│   │       └── usecase/       # Business logic
│   ├── pkg/                   # Shared utilities
│   │   ├── validations/       # Custom validators
│   │   ├── httputils/         # HTTP utilities
│   │   └── yamlutils/         # YAML utilities
│   ├── test/                  # Unit tests
│   └── .data/                 # YAML storage
├── frontend/                  # Next.js frontend
│   └── src/
│       ├── app/               # App Router pages
│       ├── components/        # Reusable UI components
│       ├── services/          # API client
│       └── types/             # TypeScript definitions
├── Dockerfile                 # Multi-stage container build
├── Makefile                   # Build automation
└── README.md                  # This file
```

- Code structure for backend follow the clean architecture pattern and domain driven design that allow the project scalable to maintain and add another feature
- Code structure for the frontend use reusable UI component and light weight API integration to make it easier to maintain

## 🚀 Quick Start

### Prerequisites

- Go 1.24+
- Node.js 18+
- Docker
- Make

### Local Development
1. **Clone this repository**
   ```bash
   git clone <repo-url>
   cd route-portal
   make prep
   ```

2. **Start backend**
   ```bash
   make run-backend
   # Runs on http://localhost:8080

3. **Start frontend**
   ```bash
   make run-frontend
   # Runs on http://localhost:3000


### Docker

1. **Build docker image**
   ```bash
   make docker-build

2. **Run with docker image (complete docker build first)**
   ```bash
   make docker-run

### Improvement
- [ ] For the net/http its quite hard to manage the endpoint, it would be better if we use http framework like Gin, Echo, Fiber and etc
- [ ] More advance error handling
- [ ] Tracer to trace the span and process that happen on the backend. So when any issue happen on production it is easier to debug. And also add more capable to do observability for monitoring the apps
- [ ] Add Kubernets manifest for apps
- [ ] Increase the unit testing coverage
- [ ] Implement linter for each backend and frontend
- [ ] Separate the business logic and UI on frontend to make the code more clean and easier to add unit testing on the frontend logic
- [ ] Implement frontend test with bdd test combined with cypress
