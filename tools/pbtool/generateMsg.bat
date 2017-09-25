@echo off

echo "delete all *.pb Files"
del /F *.pb.*

echo "generate .pb File By .proto"

protoc --go_out=. client2server.proto
protoc --go_out=plugins=grpc:. game2center.proto

echo "copy client .pb File to bin/pb"
copy /Y ".\client2server.pb.go" "../../lib/pb/client2server.pb.go"
copy /Y ".\game2center.pb.go" "../../lib/pb3/game2center.pb.go"


pause

