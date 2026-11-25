@echo off
echo Running unit tests...
go test -v ./tests/...
if %errorlevel% neq 0 exit /b %errorlevel%

echo Building config_translator...
go build -o config_translator.exe ./src
if %errorlevel% neq 0 exit /b %errorlevel%

echo Build and tests completed successfully!
