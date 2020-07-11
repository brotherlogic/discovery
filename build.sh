protoc --proto_path ../../../ -I=./proto --go_out=plugins=grpc:./proto proto/discovery.proto --js_out=import_style=commonjs:./proto  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./proto --dart_out=grpc:./proto

mv proto/github.com/brotherlogic/discovery/proto/* ./proto
