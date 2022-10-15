docker stop dashboardproject
docker rm dashboardproject
docker rmi dashboardproject
docker build -t dashboardproject:latest .
docker run -d -p 80:8000 -e SERVER_PORT="8000" -e DB_USERNAME="root" -e DB_PASSWORD="" -e DB_HOST="192.168.0.112" -e DB_PORT="3306" -e DB_NAME="dashboardproject" --name dashboardproject dashboardproject:latest