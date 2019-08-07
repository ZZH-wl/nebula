FROM alpine
ADD nebula /nebula
ADD nebula.json /nebula.json
ENTRYPOINT [ "/nebula" ]
