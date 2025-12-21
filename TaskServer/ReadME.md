# This is a test api server without framework
# The Data is only use ram to store

go run ./cmd/server

# windows test
curl http://localhost:8080/tasks `
  -Method POST `
  -Headers @{ "Content-Type" = "application/json" } `
  -Body '{ "title": "will disappear" }'

# Linux test
curl http://localhost:8080/tasks `
  -Method POST `
  -Headers @{ "Content-Type" = "application/json" } `
  -Body '{ "title": "will disappear" }'

# http://localhost:8080/tasks