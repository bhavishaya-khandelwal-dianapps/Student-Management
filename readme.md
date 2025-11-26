student-management/
│
├── cmd/
│   └── server/
│       └── main.go
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
├── go.mod
└── go.sum
