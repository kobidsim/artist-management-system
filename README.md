# Artist Management System

Management System for managing artists and their music.

Created using React and Go with sqlite3 database.

Setup:
Client:
- Run `npm install`

Server:
- Create a .env file with JWT_SECRET secret
- Run `go mod tidy` to install all packages
- Run `go run migration/migrate.go` to initialize the database
- Run `go run main.go` to start the server