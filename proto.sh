function proto {
  DOMAIN=.
  PROTO_FILE=$1.proto
  VERSION=$2
  PROTO_PATH=${DOMAIN}/api/qvbilam/file/"$VERSION"
  OUT_PATH=./${DOMAIN}/api/qvbilam/file/"$VERSION"
  protoc -I="$PROTO_PATH" --go_out "$OUT_PATH" --go_opt paths=source_relative --go-grpc_out "$OUT_PATH" --go-grpc_opt=paths=source_relative "$PROTO_FILE"
}

function userProto {
  DOMAIN=.
  PROTO_FILE=$1.proto
  VERSION=$2
  PROTO_PATH=${DOMAIN}/api/qvbilam/user/"$VERSION"
  OUT_PATH=./${DOMAIN}/api/qvbilam/user/"$VERSION"
  protoc -I="$PROTO_PATH" --go_out "$OUT_PATH" --go_opt paths=source_relative --go-grpc_out "$OUT_PATH" --go-grpc_opt=paths=source_relative "$PROTO_FILE"
}

proto video v1
proto image v1
userProto user v1

