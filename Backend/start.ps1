# 設置 CGO_ENABLED 為 0（使用純 Go SQLite 驅動）
$env:CGO_ENABLED = "0"

# 啟動 Go 伺服器
Write-Host "Starting Go backend server..." -ForegroundColor Green
go run main.go
