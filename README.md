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

### Backend
- Code structure for backend follow the clean architecture pattern and domain driven design that allow the project scalable and maintainable.
- Use customizeable struck validator that help to define fields validation for the http request payload.
- Implemented unit testing with `testify` make it easy to cover basic unit testing for each endpoint combined with self defined `httptsestutils` to help test end to end http request.

### Frontend
- Code structure for the frontend use reusable UI component to reduce redudant UI definition like data view, form and modal.
- Use MaterialUI as UI Framework, easy to use with declarative approach. Its provide ready to use component like AppBar, buttons, layouts, and etc. Very helpfull when facing fast development.
- Implemnt builtin `fetch` from NextJs and  `react-query`  from tanstack to integrate the API. Make state management for query or manipulation more easy, simple, clean.
- Use `zod` combined with `react-form` to handle the form validation. Zod help define the data schema and validatiom, react-form help manage the UI/UX state to show the current state form data and errors.

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
- [ ] For the built in net/http in golang, its quite hard to manage the endpoint, it would be better if we use http framework like Gin, Chi, Echo or Fiber.
- [ ] More advance error handling.
- [ ] Tracer to trace the span and process that happen on the backend. So when any issue happen on production it is easier to debug. Add more capabllities to do observability for monitoring the apps.
- [ ] Add Kubernets manifest for the apps.
- [ ] Increase the unit testing coverage.
- [ ] Implement linter for each backend and frontend.
- [ ] Separate the business logic and UI on frontend to make the code more clean and easier to add unit testing on the frontend logic.
- [ ] Implement frontend test with bdd test combined with cypress, help to test end to end test between backend and frontend.
- [ ] API Docs implementation with Swaggo for backend golang.
- [ ] Advance logging for backend. 
- [ ] Change the YAML storage to postgresql to make the APP stateless and easy to scale