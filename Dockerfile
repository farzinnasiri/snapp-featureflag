FROM ubuntu:20.04
WORKDIR /app

RUN mkdir "configs"
COPY ./configs/config.yml ./configs/
COPY bin/featureflag ./featureflag
ENTRYPOINT ["/app/featureflag"]

EXPOSE 6379 9000
