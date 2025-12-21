# This is a test api server without framework
# The Data is only use ram to store

go run ./cmd/server
go build -o Server ./cmd/server

# windows test
curl http://localhost:8080/tasks `
  -Method POST `
  -Headers @{ "Content-Type" = "application/json" } `
  -Body '{ "title": "will disappear" }'

# Linux test
curl -X POST http://localhost:8080/tasks \ -H "Content-Type: application/json" \ -d '{"title":"will disappear"}'

# http://localhost:8080/tasks

# Add middleware

curl http://localhost:8080/tasks   # 401
curl http://localhost:8080/tasks -Method GET -Headers @{ "X-API-Key" = "dev-key" }

curl http://localhost:8080/tasks `
  -Method POST `
  -Headers @{
    "X-API-Key"    = "dev-key"
    "Content-Type" = "application/json"
  } `
  -Body '{ "title": "will disappear" }'

# CLI cmd
## add task 新增任務
go run ./cmd/cli -cmd=create -title="from cli"

## search list
go run ./cmd/cli -cmd=list

## serach single
go run ./cmd/cli -cmd=get -id=?

# cobra cmd
go run ./cmd/cli create "hello cobra"
go run ./cmd/cli list
go run ./cmd/cli get <id>
## help
go run ./cmd/cli --help
go run ./cmd/cli create --help
## build
go build -o task ./cmd/cli
# cmd
./task list
./task create "from binary" 