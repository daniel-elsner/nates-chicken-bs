@echo off
docker build -t nates-chicken-bs .
docker run -v %USERPROFILE%\.aws:/root/.aws -e AWS_PROFILE=%profile% -p 8080:8080 -e PORT=8080 nates-chicken-bs