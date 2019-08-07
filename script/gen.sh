#! /bin/bash
echo "$(cd "$(dirname "$0")";pwd)"
cd $(dirname "$0") # SHELL_FOLDER
for file in $(ls)
do
  if [ "${file##*.}" = "proto" ]; then
    if [ ! -d "../grpc/service/${file%.*}" ]; then
      mkdir -p ../grpc/service/${file%.*}
    fi
#    if [ ! -d "../grpc/gateway/${file%.*}" ]; then
#      mkdir -p ../grpc/gateway/${file%.*}
#    fi
    #if [ ! -d "../web/grpc/${file%.*}" ]; then
    #  mkdir -p ../web/grpc/${file%.*}
    #fi
    OUT_GO_DIR=../grpc/service/${file%.*}
    OUT_GATEWAY_DIR=../grpc/gateway/${file%.*}
    OUT_WEB_DIR=../grpc/web/${file%.*}
    #protoc $file --go_out=plugins=grpc:$OUT_GO_DIR
    #protoc $file --gofast_out=plugins=grpc:$OUT_GO_DIR
    #protoc $file --js_out=import_style=commonjs:$OUT_WEB_DIR --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$OUT_WEB_DIR
    protoc $file -I/usr/local/include -I. \
      -I$GOPATH/src \
      -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --go_out=plugins=grpc:$OUT_GO_DIR \
      --micro_out=$OUT_GO_DIR
#    protoc $file -I/usr/local/include -I. \
#      -I$GOPATH/src \
#      -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#      --go_out=plugins=grpc:$OUT_GATEWAY_DIR \
#      --grpc-gateway_out=logtostderr=true:$OUT_GATEWAY_DIR
    echo "generate $file"
  fi
done