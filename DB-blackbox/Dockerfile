FROM golang:1.12
RUN go get github.com/wenkaler/blackbox
WORKDIR  /go/src/github.com/blackbox/cmd/blackbox
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM postgres:9.6

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    python3 postgresql-plpython3-9.6
COPY --from=0 /go/src/github.com/wenkaler/blackbox/app .
RUN ./app >> app.log &
COPY . /docker-entrypoint-initdb.d/ 

# Add VOLUMEs to allow backup of config, logs and databases
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]
