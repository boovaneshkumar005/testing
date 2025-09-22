
---

# Go Testing Example Project

This project demonstrates **unit tests, integration tests, and functional tests** in Go using **Testify** and a real database (SQLite). It shows a clean folder structure and a layered testing approach.

---

## **Folder Structure**

```
go-testing-example/
├── go.mod
├── main.go
├── config/
│   └── db.go
├── service/
│   ├── calculator.go
│   └── calculator_test.go
├── repository/
│   ├── user_repo.go
│   └── user_repo_test.go
├── handler/
│   ├── user_handler.go
│   └── user_handler_test.go
├── integration/
│   └── user_integration_test.go
└── functional/
    └── user_functional_test.go
```

---

## **1. About Testify**

[Testify](https://github.com/stretchr/testify) is a Go testing toolkit that provides:

* **Assertions** – simple, readable checks for your tests
* **Mocks** – mock dependencies for isolated unit tests
* **Suites** – organize tests with common setup and teardown

### **1.1 Assertions**

* `assert` continues even if the test fails
* `require` stops the test immediately on failure

**Examples:**

```go
assert.Equal(t, expected, actual, "values should match")
assert.Nil(t, err)
require.NotNil(t, user)
```

**Common assertions:** `Equal`, `NotEqual`, `Nil`, `NotNil`, `True`, `False`, `Contains`

---

### **1.2 Mocks**

Mock external dependencies to **isolate unit tests**:

```go
type UserServiceMock struct { mock.Mock }
func (m *UserServiceMock) GetUser(id int) string { return m.Called(id).String(0) }
```

* Define expectations and return values
* Ensure methods are called with correct parameters

---

### **1.3 Test Suites**

Group related tests with **common setup/teardown**:

```go
type CalculatorTestSuite struct { suite.Suite }

func (s *CalculatorTestSuite) SetupTest() { /* runs before each test */ }
func (s *CalculatorTestSuite) TestAdd() { s.Equal(5, Add(2,3)) }

func TestCalculatorTestSuite(t *testing.T) {
    suite.Run(t, new(CalculatorTestSuite))
}
```

---

### **1.4 Why Use Testify**

* Improves **readability** and **maintainability**
* Reduces **boilerplate** code
* Supports **mocking and suites**
* Works for **unit, integration, and functional tests**

---

## **2. Why modernc.org/sqlite**

By default, Go projects often use `github.com/mattn/go-sqlite3`. However:

* It **requires CGO**, which may fail in **CI/CD** or Docker alpine images.
* Error example:

```
Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work.
```

**Solution:** Use [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite), a **pure-Go SQLite driver**:

```bash
go get modernc.org/sqlite
```

**Benefits:**

* Works with **CGO disabled**
* Runs on **any platform** without extra dependencies
* Ideal for **testing and CI pipelines**

---

## **3. How Tests Are Organized**

| Test Type             | Example Files                                                     | DB Used?       | Description                                 |
| --------------------- | ----------------------------------------------------------------- | -------------- | ------------------------------------------- |
| **Unit Tests**        | `service/calculator_test.go`, `repository/user_repo_unit_test.go` | Mocked/None    | Test **small, isolated logic** using mocks  |
| **Integration Tests** | `integration/user_integration_test.go`                            | Real SQLite DB | Test repository logic against **real DB**   |
| **Functional Tests**  | `functional/user_functional_test.go`                              | Real SQLite DB | Test **full HTTP flow** with service and DB |

---

## **4. Testing Flow Diagram**

```mermaid
flowchart LR
    A[Unit Tests] --> B[Integration Tests] --> C[Functional Tests]
    
    subgraph Unit Tests
        A1[Service Layer Logic]
        A2[Repository Logic (Mocked DB)]
        A1 --> A2
    end

    subgraph Integration Tests
        B1[Repository Layer + Real DB (SQLite)]
        B1 --> B2[Test DB Queries]
    end

    subgraph Functional Tests
        C1[HTTP Handlers + Service + Repository + Real DB]
        C1 --> C2[Test Full Application Flow]
    end
```

**Explanation:**

1. **Unit Tests:** Isolated functions or repository methods with **mocked dependencies**.
2. **Integration Tests:** Repository + real database, validating **SQL queries and DB interaction**.
3. **Functional Tests:** Full application flow via HTTP endpoints, including **service logic + DB**.

---

## **5. How to Run Tests**

```bash
# Run all tests
go test ./...

# Run unit tests only
go test ./service ./repository

# Run integration tests only
go test ./integration

# Run functional tests only
go test ./functional
```

---

## **6. How to Run the Application**

```bash
go run main.go
```

* The HTTP server runs at **:8080**
* Endpoints:

```
POST /user      -> Add a user
GET  /user?id=  -> Get user by ID
```

---

## **7. Advantages of This Approach**

* **Readable tests** using Testify assertions
* **Mocking** allows isolated unit testing
* **Integration tests** verify DB interaction
* **Functional tests** verify end-to-end application behavior
* Using **modernc.org/sqlite** avoids CGO issues and simplifies CI/CD

---

This README now contains:

1. Full **project explanation**
2. Detailed **Testify section**
3. **modernc.org/sqlite explanation**
4. Folder structure and file descriptions
5. **Testing flow diagram** for unit → integration → functional

---
