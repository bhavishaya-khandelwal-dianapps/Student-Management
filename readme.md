student-management/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── configs/                  <-- YAML / JSON config files
│   ├── app.yaml
│   ├── db.yaml
│   └── logger.yaml
│
├── deployments/              <-- Kubernetes YAMLs
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
│
├── api/                      <-- Swagger / OpenAPI YAML
│   └── openapi.yaml
│
├── internal/
│   ├── handlers/        # controllers
│   │   └── student_handler.go
│   │
│   ├── services/        # service layer
│   │   └── student_service.go
│   │
│   ├── repositories/    # DB queries
│   │   └── student_repository.go
│   │
│   ├── models/          # models / structs
│   │   └── student.go
│   │
│   ├── validations/     # request validations
│   │   └── student_validation.go
│   │
│   ├── routes/          # router files
│   │   └── routes.go
│   │
│   ├── config/          # env, DB config
│   │   └── config.go
│   │
│   └── database/        # DB connection
│       └── db.go
│
├── pkg/                      <-- optional reusable code
│
├── go.mod
└── go.sum
