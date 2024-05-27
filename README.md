# GO Games Fiber CRUD API with PostgreSQL.
## Setup
```
go mod tidy
go install github.com/cosmtrek/air@latest
docker compose up -d
air
```
## Routes
go see [router.go](./cmd/api/router/router.go)
everything is received in JSON, to create a game should use : 
{
 "title": "",
 "description": ""
}
```
POST       "/game"         // Create a game
GET        "/game/:id"     // Get a game by ID
GET        "/games"        // Lists All Games
PUT        "/game/:id"     // Update a game by ID
DELETE     "game/:id"      // Delete a game by ID
```