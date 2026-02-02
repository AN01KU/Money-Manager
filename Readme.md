# Money Manager (Splitwise Clone)

A personal expense-sharing application built in Go.

## Tech Stack

- **Language:** Go 1.25
- **Router:** chi or gin (recommended)
- **Database:** PostgreSQL
- **Auth:** JWT

## Project Structure

```
money-manager/
├── cmd/
│   └── server/
│       └── main.go           # Entry point
├── internal/
│   ├── config/
│   │   └── config.go         # Env/config loading
│   ├── handler/
│   │   ├── auth.go
│   │   ├── group.go
│   │   ├── expense.go
│   │   └── settlement.go
│   ├── middleware/
│   │   └── auth.go           # JWT middleware
│   ├── model/
│   │   ├── user.go
│   │   ├── group.go
│   │   ├── expense.go
│   │   └── settlement.go
│   ├── repository/
│   │   ├── user.go
│   │   ├── group.go
│   │   ├── expense.go
│   │   └── settlement.go
│   └── service/
│       ├── auth.go
│       ├── group.go
│       ├── expense.go
│       └── balance.go        # Balance calculation logic
├── pkg/
│   └── utils/
│       └── response.go       # JSON response helpers
├── migrations/
│   └── 001_init.sql
├── go.mod
├── go.sum
└── Readme.md
```

## Getting Started

### 1. Set up PostgreSQL database
### 2. Create `.env` file
### 3. Run migrations
### 4. Start server

## Implementation Order

1. **Config & Database** - Set up config loading, DB connection
2. **Models** - Define structs for User, Group, Expense, Split, Settlement
3. **Auth** - Signup/Login with JWT
4. **Groups** - CRUD + member management
5. **Expenses** - Create expenses with splits
6. **Balances** - Calculate who owes whom
7. **Settlements** - Record payments between users

---

## API Endpoints

### Authentication

#### Signup
```bash
POST /auth/signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}

Response:
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "created_at": "2025-01-26T12:00:00Z"
  }
}
```

#### Login
```bash
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}

Response: Same as signup
```

### Groups

#### Create Group
```bash
POST /groups
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Weekend Trip"
}

Response:
{
  "id": "650e8400-e29b-41d4-a716-446655440000",
  "name": "Weekend Trip",
  "created_by": "550e8400-e29b-41d4-a716-446655440000",
  "created_at": "2025-01-26T12:00:00Z"
}
```

#### Add Member
```bash
POST /groups/:id/add-member
Authorization: Bearer <token>
Content-Type: application/json

{
  "user_id": "750e8400-e29b-41d4-a716-446655440000"
}

Response:
{
  "message": "member added"
}
```

### Expenses

#### Create Expense
```bash
POST /expenses
Authorization: Bearer <token>
Content-Type: application/json

{
  "group_id": "650e8400-e29b-41d4-a716-446655440000",
  "description": "Dinner",
  "total_amount": "100.00",
  "splits": [
    {
      "user_id": "550e8400-e29b-41d4-a716-446655440000",
      "amount": "50.00"
    },
    {
      "user_id": "750e8400-e29b-41d4-a716-446655440000",
      "amount": "50.00"
    }
  ]
}

Response:
{
  "id": "850e8400-e29b-41d4-a716-446655440000",
  "group_id": "650e8400-e29b-41d4-a716-446655440000",
  "description": "Dinner",
  "total_amount": "100.00",
  "paid_by": "550e8400-e29b-41d4-a716-446655440000",
  "created_at": "2025-01-26T12:00:00Z",
  "splits": [...]
}
```

#### Get Group Expenses
```bash
GET /groups/:id/expenses
Authorization: Bearer <token>

Response: Array of expenses
```

### Balances

#### Get Group Balances
```bash
GET /groups/:id/balances
Authorization: Bearer <token>

Response:
[
  {
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "amount": "-25.50"  # Negative = is owed, Positive = owes
  },
  {
    "user_id": "750e8400-e29b-41d4-a716-446655440000",
    "amount": "25.50"   # Owes 25.50
  }
]
```

### Settlements

#### Create Settlement
```bash
POST /settlements
Authorization: Bearer <token>
Content-Type: application/json

{
  "group_id": "650e8400-e29b-41d4-a716-446655440000",
  "from_user": "750e8400-e29b-41d4-a716-446655440000",
  "to_user": "550e8400-e29b-41d4-a716-446655440000",
  "amount": "25.50"
}

Response:
{
  "id": "950e8400-e29b-41d4-a716-446655440000",
  "group_id": "650e8400-e29b-41d4-a716-446655440000",
  "from_user": "750e8400-e29b-41d4-a716-446655440000",
  "to_user": "550e8400-e29b-41d4-a716-446655440000",
  "amount": "25.50",
  "created_at": "2025-01-26T12:00:00Z"
}
```
