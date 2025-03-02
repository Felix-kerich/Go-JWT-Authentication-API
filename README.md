# Go JWT Authentication API

A robust REST API built with Go (Gin) that implements JWT authentication, email verification, and user management with soft delete functionality.

## Features

- JWT-based authentication
- Email verification
- User management (CRUD operations)
- Password hashing with bcrypt
- Soft delete with recovery option
- MySQL database with GORM
- Role-based authorization
- Password reset functionality

## Tech Stack

- Go (Golang)
- Gin Web Framework
- GORM (MySQL)
- JWT-Go
- Bcrypt
- SMTP for email

## Configuration

- SMTP server credentials for email verification
- MySQL database credentials
- JWT secret key
- Bcrypt cost factor

## Environment Variables

- `SMTP_HOST`: SMTP server host
- `SMTP_PORT`: SMTP server port
- `SMTP_USERNAME`: SMTP server username
- `SMTP_PASSWORD`: SMTP server password
- `DB_HOST`: MySQL database host
- `DB_PORT`: MySQL database port
- `DB_USER`: MySQL database username
- `DB_PASSWORD`: MySQL database password
- `DB_NAME`: MySQL database name
- `JWT_SECRET`: JWT secret key
- `BCRYPT_COST`: Bcrypt cost factor


## API Endpoints

### Authentication
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/auth/register` | Register a new user |
| POST | `/auth/login` | Login and receive JWT token |
| GET | `/auth/verify` | Verify email with token |

### User Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | Get all users (Admin only) |
| GET | `/users/:id` | Get user by ID |
| PUT | `/users/:id` | Update user |
| DELETE | `/users/:id` | Soft delete user |
| POST | `/users/:id/recover` | Recover deleted user |

## Getting Started

### Prerequisites

- Go 1.16+
- MySQL
- SMTP server credentials for email verification

### Installation

1. Clone the repository

```bash
git clone https://github.com/yourusername/go-jwt-auth.git
cd go-jwt-auth
```

2. Install dependencies

```bash
go mod tidy
```

3. Configure environment variables

```bash
cp .env.example .env
```

4. Run the application

```bash
go run main.go
```

## API Documentation

The API documentation is available in the `docs` directory.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

