FROM alpine
ADD nebula /nebula
ADD nebula.json /nebula.json
WORKDIR /nebula
ENTRYPOINT [ "/nebula" ]
