FROM ubuntu:latest
LABEL authors="LOQ"

ENTRYPOINT ["top", "-b"]