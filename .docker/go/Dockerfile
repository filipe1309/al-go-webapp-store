FROM golang:1.16-alpine3.13
RUN apk update && apk upgrade && apk add curl && apk add --no-cache git
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV CGO_ENABLED=0
ENV APP_HOME /go/src/go-webapp-store
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
WORKDIR $APP_HOME
COPY . .
EXPOSE 8000
# COPY go.* .
# RUN go mod vendor
# RUN go mod download 
CMD air