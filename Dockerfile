FROM golang:1.15.1-alpine3.12

RUN apk update && \
  apk add --no-cache gcc libc-dev curl git pkgconfig openssl && \
  wget https://github.com/neo4j-drivers/seabolt/releases/download/v1.7.4/seabolt-1.7.4-Linux-alpine-3.9.3.tar.gz && \
  tar zxvf seabolt-1.7.4-Linux-alpine-3.9.3.tar.gz --strip-components=1 -C / && \
  rm seabolt-1.7.4-Linux-alpine-3.9.3.tar.gz
ENV LD_LIBRARY_PATH ${LD_LIBRARY_PATH}:/usr/local/lib64

RUN mkdir /app
WORKDIR /app
