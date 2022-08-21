function proto {
  DOMAIN=.
  PROTO_FILE=$1.proto
  PROTO_PATH=${DOMAIN}/api/pb
  OUT_PATH=./${DOMAIN}/api/pb
  protoc -I=$PROTO_PATH --go_out $OUT_PATH --go_opt paths=source_relative --go-grpc_out $OUT_PATH --go-grpc_opt=paths=source_relative $PROTO_FILE
}

function userProto {
  DOMAIN=.
  PROTO_FILE=$1.proto
  PROTO_PATH=${DOMAIN}/api/user/pb
  OUT_PATH=./${DOMAIN}/api/user/pb
  protoc -I=$PROTO_PATH --go_out $OUT_PATH --go_opt paths=source_relative --go-grpc_out $OUT_PATH --go-grpc_opt=paths=source_relative $PROTO_FILE
}

proto file
proto config
userProto user
