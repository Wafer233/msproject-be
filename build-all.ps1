# build-all.ps1
Write-Output "starting..."

# 先到项目根目录
Set-Location -Path $PSScriptRoot

# 编译 api-gateway
Write-Output "starting api-gateway..."
Set-Location -Path "./api-gateway"
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o "./cmd/bin/api-gateway" ./cmd/main.go

# 编译 user-service
Write-Output "starting user-service..."
Set-Location -Path "../user-service"
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o "./cmd/bin/user-service" ./cmd/main.go

# 编译 project-service
Write-Output "starting project-service..."
Set-Location -Path "../project-service"
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o "./cmd/bin/project-service" ./cmd/main.go

Write-Output "successfully build all services."
