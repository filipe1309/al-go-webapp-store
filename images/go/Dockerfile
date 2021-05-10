FROM golang:1.16-alpine3.13
RUN apk update && apk upgrade && apk add curl git
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_HOME /go/src/go-webapp-store
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
WORKDIR $APP_HOME
EXPOSE 8000
COPY . .
CMD air