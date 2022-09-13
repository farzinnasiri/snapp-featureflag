FROM ubuntu:20.04
WORKDIR /app

RUN mkdir "configs"
COPY ./config/config.yml ./config/
COPY bin/featureflag ./featureflag
ENTRYPOINT ["/app/featureflag"]

EXPOSE 6379 8080
