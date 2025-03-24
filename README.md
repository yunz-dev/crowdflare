# crowdflare

### Project Structure
```
myproject/
│── cmd/
│   ├── web/                # Web server entry point
│   │   ├── main.go
│   └── worker/             # Background workers (if needed)
│       ├── main.go
│
├── internal/
│   ├── handlers/           # HTMX handlers
│   │   ├── home.go
│   │   ├── user.go
│   ├── templates/          # HTML templates (HTMX components)
│   │   ├── layouts/
│   │   │   ├── base.html
│   │   ├── partials/
│   │   │   ├── navbar.html
│   │   ├── pages/
│   │   │   ├── index.html
│   ├── services/           # Business logic
│   │   ├── user_service.go
│   ├── db/                 # Database layer
│   │   ├── models.go
│   │   ├── queries.go
│   ├── static/             # Static files
│   │   ├── css/
│   │   │   ├── styles.css  # Compiled Tailwind CSS
│   │   ├── js/
│   │   │   ├── alpine.js   # Alpine.js
│   │   ├── images/
│   ├── middleware/         # Middleware for authentication, logging, etc.
│   │   ├── auth.go
│   │   ├── logging.go
│
├── config/                 # Configuration files
│   ├── config.yaml
│   ├── env.go
│
├── pkg/                    # Reusable packages
│   ├── utils/              # Helper functions
│   ├── validation/         # Input validation logic
│
├── migrations/             # Database migrations
│   ├── 001_initial.up.sql
│   ├── 001_initial.down.sql
│
├── public/                 # Public assets (e.g., favicon, robots.txt)
│
├── go.mod
├── go.sum
├── Makefile                # Automation scripts
└── README.md

```
