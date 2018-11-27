GOOS   := linux
GOARCH := amd64

deps:
	dep ensure

run:
	go run main.go

run-mysql:
	docker run --rm --name vg-1day-mysql -p 127.0.0.1:13306:3306 \
	  -v $(CURDIR)/migrations:/docker-entrypoint-initdb.d \
	  -e MYSQL_DATABASE=sampledb -e MYSQL_ROOT_PASSWORD=password -d mysql:5.6

clean:
	docker stop vg-1day-mysql

mysql-debug:
	mysql -uroot -ppassword -h127.0.0.1 -P13306

event:
	 curl http://localhost:8081/event\?name\=hoge\&value\=fuga

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o app
