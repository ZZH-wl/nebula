FROM alpine
ADD nebula /nebula
ADD nebula.json /runtime/nebula.json
VOLUME ["/runtime"]
WORKDIR /runtime
ENTRYPOINT [ "/nebula" ]
EXPOSE 8080