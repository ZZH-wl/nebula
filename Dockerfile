FROM alpine
ADD nebula /nebula
ADD nebula.json /nebula.json
VOLUME ["/runtime"]
WORKDIR /runtime
ENTRYPOINT [ "/nebula" ]
EXPOSE 8080