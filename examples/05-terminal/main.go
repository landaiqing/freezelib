package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/freezelib"
)

func main() {
	fmt.Println("💻 Terminal Output Examples")
	fmt.Println("============================")

	// Create output directory
	os.MkdirAll("output", 0755)

	// Run terminal examples
	basicAnsiExample()
	buildOutputExample()
	testResultsExample()
	dockerOutputExample()
	gitOutputExample()
	systemLogsExample()

	fmt.Println("\n✅ Terminal examples completed!")
	fmt.Println("📁 Check the 'output' directory for generated files.")
}

// Basic ANSI color example
func basicAnsiExample() {
	fmt.Println("\n🌈 Basic ANSI Colors")
	fmt.Println("--------------------")

	// Terminal preset optimized for ANSI output
	freeze := freezelib.NewWithPreset("terminal")

	ansiOutput := `\033[32m✓ SUCCESS\033[0m: Application started successfully
\033[33m⚠ WARNING\033[0m: Configuration file not found, using defaults
\033[31m✗ ERROR\033[0m: Failed to connect to database
\033[36mINFO\033[0m: Server listening on port 8080
\033[35mDEBUG\033[0m: Loading user preferences
\033[37mTRACE\033[0m: Function call: getUserById(123)

\033[1mBold text\033[0m and \033[4munderlined text\033[0m
\033[7mReversed text\033[0m and \033[9mstrikethrough text\033[0m

Background colors:
\033[41mRed background\033[0m
\033[42mGreen background\033[0m
\033[43mYellow background\033[0m
\033[44mBlue background\033[0m`

	svgData, err := freeze.GenerateFromANSI(ansiOutput)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	err = os.WriteFile("output/basic_ansi.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Println("✅ Generated: output/basic_ansi.svg")
}

// Build output example
func buildOutputExample() {
	fmt.Println("\n🔨 Build Output")
	fmt.Println("---------------")

	freeze := freezelib.New().
		WithTheme("github-dark").
		WithFont("Cascadia Code", 13).
		WithWindow(true).
		WithPadding(20).
		WithBackground("#0d1117")

	buildOutput := `$ go build -v ./...
github.com/myproject/internal/config
github.com/myproject/internal/database
github.com/myproject/internal/handlers
github.com/myproject/cmd/server

\033[32m✓ Build completed successfully\033[0m

$ go test -v ./...
=== RUN   TestUserService_CreateUser
--- PASS: TestUserService_CreateUser (0.01s)
=== RUN   TestUserService_GetUser
--- PASS: TestUserService_GetUser (0.00s)
=== RUN   TestUserService_UpdateUser
--- PASS: TestUserService_UpdateUser (0.01s)
=== RUN   TestUserService_DeleteUser
--- PASS: TestUserService_DeleteUser (0.00s)
=== RUN   TestDatabaseConnection
--- PASS: TestDatabaseConnection (0.05s)

\033[32mPASS\033[0m
\033[32mok  \033[0m	github.com/myproject	0.123s

\033[36mCoverage: 85.7% of statements\033[0m

$ docker build -t myapp:latest .
Sending build context to Docker daemon  2.048kB
Step 1/8 : FROM golang:1.21-alpine AS builder
 ---> 7642119cd161
Step 2/8 : WORKDIR /app
 ---> Using cache
 ---> 8f3b8c9d4e5f
Step 3/8 : COPY go.mod go.sum ./
 ---> Using cache
 ---> 1a2b3c4d5e6f
Step 4/8 : RUN go mod download
 ---> Using cache
 ---> 2b3c4d5e6f7g
Step 5/8 : COPY . .
 ---> 3c4d5e6f7g8h
Step 6/8 : RUN go build -o main .
 ---> Running in 4d5e6f7g8h9i
 ---> 5e6f7g8h9i0j
Step 7/8 : FROM alpine:latest
 ---> 6f7g8h9i0j1k
Step 8/8 : COPY --from=builder /app/main /main
 ---> 7g8h9i0j1k2l
\033[32mSuccessfully built 7g8h9i0j1k2l\033[0m
\033[32mSuccessfully tagged myapp:latest\033[0m`

	svgData, err := freeze.GenerateFromANSI(buildOutput)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	err = os.WriteFile("output/build_output.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Println("✅ Generated: output/build_output.svg")
}

// Test results example
func testResultsExample() {
	fmt.Println("\n🧪 Test Results")
	fmt.Println("---------------")

	freeze := freezelib.New().
		WithTheme("dracula").
		WithFont("JetBrains Mono", 14).
		WithWindow(true).
		WithLineNumbers(false).
		WithShadow(15, 0, 10).
		WithPadding(25)

	testOutput := `$ npm test

> myapp@1.0.0 test
> jest --coverage

 PASS  src/components/Button.test.js
  Button Component
    \033[32m✓\033[0m renders correctly (15ms)
    \033[32m✓\033[0m handles click events (8ms)
    \033[32m✓\033[0m applies custom className (3ms)

 PASS  src/services/api.test.js
  API Service
    \033[32m✓\033[0m fetches user data (45ms)
    \033[32m✓\033[0m handles network errors (12ms)
    \033[32m✓\033[0m retries failed requests (23ms)

 FAIL  src/utils/validation.test.js
  Validation Utils
    \033[32m✓\033[0m validates email addresses (5ms)
    \033[31m✗\033[0m validates phone numbers (8ms)
    \033[32m✓\033[0m validates passwords (3ms)

  ● Validation Utils › validates phone numbers

    expect(received).toBe(expected)

    Expected: true
    Received: false

      12 |   test('validates phone numbers', () => {
      13 |     const phoneNumber = '+1-555-123-4567';
    > 14 |     expect(isValidPhoneNumber(phoneNumber)).toBe(true);
         |                                             ^
      15 |   });

    at Object.<anonymous> (src/utils/validation.test.js:14:45)

\033[33mTest Suites: 1 failed, 2 passed, 3 total\033[0m
\033[33mTests:       1 failed, 7 passed, 8 total\033[0m
\033[33mSnapshots:   0 total\033[0m
\033[33mTime:        2.847s\033[0m

\033[36m----------------------|---------|----------|---------|---------|-------------------\033[0m
\033[36mFile                  | % Stmts | % Branch | % Funcs | % Lines | Uncovered Line #s\033[0m
\033[36m----------------------|---------|----------|---------|---------|-------------------\033[0m
\033[36mAll files             |   87.5   |   75.0   |   90.0  |   87.5  |\033[0m
\033[36m src/components       |   95.0   |   85.0   |  100.0  |   95.0  |\033[0m
\033[36m  Button.js           |   95.0   |   85.0   |  100.0  |   95.0  | 23\033[0m
\033[36m src/services         |   80.0   |   65.0   |   80.0  |   80.0  |\033[0m
\033[36m  api.js              |   80.0   |   65.0   |   80.0  |   80.0  | 45,67\033[0m
\033[36m src/utils             |   87.5   |   75.0   |   90.0  |   87.5  |\033[0m
\033[36m  validation.js       |   87.5   |   75.0   |   90.0  |   87.5  | 34\033[0m
\033[36m----------------------|---------|----------|---------|---------|-------------------\033[0m`

	svgData, err := freeze.GenerateFromANSI(testOutput)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	err = os.WriteFile("output/test_results.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Println("✅ Generated: output/test_results.svg")
}

// Docker output example
func dockerOutputExample() {
	fmt.Println("\n🐳 Docker Output")
	fmt.Println("----------------")

	freeze := freezelib.New().
		WithTheme("nord").
		WithFont("SF Mono", 13).
		WithWindow(true).
		WithPadding(20).
		WithBackground("#2e3440")

	dockerOutput := `$ docker-compose up -d
Creating network "myapp_default" with the default driver
Creating volume "myapp_postgres_data" with default driver
Creating volume "myapp_redis_data" with default driver

\033[33mPulling postgres (postgres:13)...\033[0m
13: Pulling from library/postgres
\033[36m7b1a6ab2e44d\033[0m: Pull complete
\033[36m5c9d4e5f6a7b\033[0m: Pull complete
\033[36m8c1d2e3f4a5b\033[0m: Pull complete
\033[36m9d2e3f4a5b6c\033[0m: Pull complete
\033[32mDigest: sha256:abc123def456...\033[0m
\033[32mStatus: Downloaded newer image for postgres:13\033[0m

\033[33mPulling redis (redis:6-alpine)...\033[0m
6-alpine: Pulling from library/redis
\033[36m4c0d5e6f7a8b\033[0m: Pull complete
\033[36m5d1e6f7a8b9c\033[0m: Pull complete
\033[32mDigest: sha256:def456ghi789...\033[0m
\033[32mStatus: Downloaded newer image for redis:6-alpine\033[0m

Creating myapp_postgres_1 ... \033[32mdone\033[0m
Creating myapp_redis_1    ... \033[32mdone\033[0m
Creating myapp_web_1      ... \033[32mdone\033[0m

$ docker ps
CONTAINER ID   IMAGE           COMMAND                  CREATED         STATUS         PORTS                    NAMES
\033[36m1a2b3c4d5e6f\033[0m   myapp:latest    "go run main.go"         2 minutes ago   Up 2 minutes   \033[35m0.0.0.0:8080->8080/tcp\033[0m   myapp_web_1
\033[36m2b3c4d5e6f7g\033[0m   postgres:13     "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes   \033[35m5432/tcp\033[0m                 myapp_postgres_1
\033[36m3c4d5e6f7g8h\033[0m   redis:6-alpine  "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes   \033[35m6379/tcp\033[0m                 myapp_redis_1

$ docker logs myapp_web_1
\033[36m2024/01/15 10:30:00\033[0m \033[32mINFO\033[0m Starting server...
\033[36m2024/01/15 10:30:00\033[0m \033[32mINFO\033[0m Connected to database
\033[36m2024/01/15 10:30:00\033[0m \033[32mINFO\033[0m Connected to Redis
\033[36m2024/01/15 10:30:00\033[0m \033[32mINFO\033[0m Server listening on :8080
\033[36m2024/01/15 10:30:15\033[0m \033[36mDEBUG\033[0m GET /api/health - 200 OK (2ms)
\033[36m2024/01/15 10:30:20\033[0m \033[36mDEBUG\033[0m POST /api/users - 201 Created (45ms)`

	svgData, err := freeze.GenerateFromANSI(dockerOutput)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	err = os.WriteFile("output/docker_output.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Println("✅ Generated: output/docker_output.svg")
}

// Git output example
func gitOutputExample() {
	fmt.Println("\n📚 Git Output")
	fmt.Println("-------------")

	freeze := freezelib.New().
		WithTheme("github").
		WithFont("Menlo", 13).
		WithWindow(true).
		WithPadding(20).
		WithBackground("#ffffff")

	gitOutput := `$ git status
On branch feature/user-authentication
Your branch is ahead of 'origin/main' by 3 commits.
  (use "git push" to publish your local commits)

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	\033[32mnew file:   src/auth/login.js\033[0m
	\033[32mnew file:   src/auth/register.js\033[0m
	\033[32mmodified:   src/app.js\033[0m
	\033[32mmodified:   package.json\033[0m

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	\033[31mmodified:   README.md\033[0m
	\033[31mmodified:   src/components/Header.js\033[0m

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	\033[31msrc/auth/middleware.js\033[0m
	\033[31mtests/auth.test.js\033[0m

$ git log --oneline -5
\033[33m7a8b9c0\033[0m \033[32m(HEAD -> feature/user-authentication)\033[0m Add user registration functionality
\033[33m6a7b8c9\033[0m Add login form validation
\033[33m5a6b7c8\033[0m Implement JWT authentication
\033[33m4a5b6c7\033[0m \033[36m(origin/main, main)\033[0m Update project dependencies
\033[33m3a4b5c6\033[0m Fix responsive design issues

$ git diff --stat
 README.md                  |  15 \033[32m+++++++++\033[0m\033[31m------\033[0m
 package.json               |   3 \033[32m+++\033[0m
 src/app.js                 |  42 \033[32m++++++++++++++++++++++++++++++\033[0m\033[31m----------\033[0m
 src/auth/login.js          |  67 \033[32m+++++++++++++++++++++++++++++++++++++++++++++++++++++++\033[0m
 src/auth/register.js       |  89 \033[32m++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\033[0m
 src/components/Header.js   |   8 \033[32m+++++\033[0m\033[31m---\033[0m
 6 files changed, 201 insertions(+), 23 deletions(-)`

	svgData, err := freeze.GenerateFromANSI(gitOutput)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	err = os.WriteFile("output/git_output.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Println("✅ Generated: output/git_output.svg")
}

// System logs example
func systemLogsExample() {
	fmt.Println("\n📋 System Logs")
	fmt.Println("--------------")

	freeze := freezelib.New().
		WithTheme("monokai").
		WithFont("Ubuntu Mono", 13).
		WithWindow(true).
		WithPadding(20).
		WithBackground("#272822")

	systemLogs := `$ tail -f /var/log/application.log

\033[90m2024-01-15 10:30:00.123\033[0m [\033[32mINFO \033[0m] \033[36mApplication\033[0m - Server starting up
\033[90m2024-01-15 10:30:00.456\033[0m [\033[32mINFO \033[0m] \033[36mDatabase \033[0m - Connection pool initialized (size: 10)
\033[90m2024-01-15 10:30:00.789\033[0m [\033[32mINFO \033[0m] \033[36mCache    \033[0m - Redis connection established
\033[90m2024-01-15 10:30:01.012\033[0m [\033[32mINFO \033[0m] \033[36mSecurity \033[0m - JWT secret loaded from environment
\033[90m2024-01-15 10:30:01.234\033[0m [\033[32mINFO \033[0m] \033[36mHTTP     \033[0m - Server listening on port 8080

\033[90m2024-01-15 10:30:15.567\033[0m [\033[34mDEBUG\033[0m] \033[36mAuth     \033[0m - User login attempt: user@example.com
\033[90m2024-01-15 10:30:15.678\033[0m [\033[32mINFO \033[0m] \033[36mAuth     \033[0m - User authenticated successfully: user@example.com
\033[90m2024-01-15 10:30:15.789\033[0m [\033[34mDEBUG\033[0m] \033[36mHTTP     \033[0m - POST /api/auth/login - 200 OK (223ms)

\033[90m2024-01-15 10:30:30.123\033[0m [\033[33mWARN \033[0m] \033[36mDatabase \033[0m - Slow query detected (1.2s): SELECT * FROM users WHERE...
\033[90m2024-01-15 10:30:30.234\033[0m [\033[34mDEBUG\033[0m] \033[36mHTTP     \033[0m - GET /api/users - 200 OK (1234ms)

\033[90m2024-01-15 10:30:45.456\033[0m [\033[31mERROR\033[0m] \033[36mPayment  \033[0m - Payment processing failed: insufficient funds
\033[90m2024-01-15 10:30:45.567\033[0m [\033[31mERROR\033[0m] \033[36mPayment  \033[0m - Stack trace:
    at PaymentService.processPayment (payment.js:45:12)
    at OrderController.createOrder (order.js:23:8)
    at Router.handle (express.js:123:5)
\033[90m2024-01-15 10:30:45.678\033[0m [\033[34mDEBUG\033[0m] \033[36mHTTP     \033[0m - POST /api/orders - 400 Bad Request (112ms)

\033[90m2024-01-15 10:31:00.789\033[0m [\033[32mINFO \033[0m] \033[36mScheduler\033[0m - Running daily cleanup task
\033[90m2024-01-15 10:31:05.012\033[0m [\033[32mINFO \033[0m] \033[36mScheduler\033[0m - Cleanup completed: removed 1,234 expired sessions
\033[90m2024-01-15 10:31:05.123\033[0m [\033[32mINFO \033[0m] \033[36mScheduler\033[0m - Next cleanup scheduled for 2024-01-16 10:31:00`

	svgData, err := freeze.GenerateFromANSI(systemLogs)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	err = os.WriteFile("output/system_logs.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Println("✅ Generated: output/system_logs.svg")
}
