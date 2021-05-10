FROM golang:1.16-alpine3.13
# RUN go get -u github.com/beego/bee
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/go-webapp-store
# RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER # FROM golang:1.16.3-buster
RUN addgroup -S $APP_USER && adduser -S $APP_USER -G $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
USER $APP_USER
WORKDIR $APP_HOME
EXPOSE 8000
# CMD ["bee", "run"]
COPY . .
#RUN CGO_ENABLED=0 go build -o /go/bin/store
ENTRYPOINT ["go", "run", "main.go"]