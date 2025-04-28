# wafer233.ps1
# Build user-service, project-service, api-gateway (short version)

# api-gateway
docker build -f Dockerfile -t wafer233/project-service:latest .
