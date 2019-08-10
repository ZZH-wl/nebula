#! /bin/bash
cd proto
for file in $(ls)
do
  if [ "${file##*.}" = "proto" ]; then
    if [ ! -d "srv/${file%.*}" ]; then
      mkdir -p srv/${file%.*}
    fi
    #if [ ! -d "../gateway/${file%.*}" ]; then
    #  mkdir -p ../gateway/${file%.*}
    #fi
    #if [ ! -d "../web/${file%.*}" ]; then
    #  mkdir -p ../web/${file%.*}
    #fi
    OUT_GO_DIR=srv/${file%.*}
    OUT_GATEWAY_DIR=gateway/${file%.*}
    OUT_WEB_DIR=web/${file%.*}
    #protoc $file --go_out=plugins=grpc:$OUT_GO_DIR
    #protoc $file --gofast_out=plugins=grpc:$OUT_GO_DIR
    #protoc $file --js_out=import_style=commonjs:$OUT_WEB_DIR --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$OUT_WEB_DIR
    protoc $file -I. \
      -I/usr/local/include \
      -I$GOPATH/src \
      #-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
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