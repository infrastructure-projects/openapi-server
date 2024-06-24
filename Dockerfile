FROM centos:centos7
COPY app /app
ENV PARAMS=""
ENV GIN_MODE="release"
WORKDIR /data
ENTRYPOINT ["sh", "-c", "/app $PARAMS"]