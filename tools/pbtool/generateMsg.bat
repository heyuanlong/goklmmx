@echo off

echo "delete all *.pb Files"
del /F *.pb.*

echo "generate .pb File By .proto"

protoc --go_out=. client2server.proto

echo "copy client .pb File to bin/pb"
copy /Y ".\client2server.pb.go" "../../lib/pb/client2server.pb.go"



pause
