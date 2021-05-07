FROM golang:1.15-alpine

MAINTAINER GodYao liyaoo1995@gmail.com

EXPOSE 8080
EXPOSE 8081

WORKDIR /backend/

COPY . .

RUN cd cmd/admin && go build -mod=vendor -ldflags="-w -s" -o=../../bin/admin .

CMD ["-p=config", "-n=config.yaml", "-m=dev"]

ENTRYPOINT ["bin/admin"]



