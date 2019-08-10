@echo off
cd proto
for /f "delims=" %%f in ('dir  /b/a-d/s  *.proto') do (
echo %%~nf
protoc %%~nf.proto -I. -I %GOPATH%/src  --go_out=plugins=grpc:.\srv\%%~nf --micro_out=.\srv\%%~nf
)
cd ..