.PHONY: build clean run docker-build docker-run push tag build-linux-best build-linux-fast

BIN_NAME=admin
IMAGE_NAME=admin
IMAGE_TAG=$(shell date +"%Y%m%d%H%M%S")
IMAGE_ID=$(shell docker images | grep ${IMAGE_NAME} | awk '{print $3}' | awk 'NR==1{print}')
IMAGE_NAME_TAG=$(shell docker images | grep ${IMAGE_NAME} | awk '{print $$1":"$$2}' | awk 'NR==1{print}')


build:
	@echo "Building $(BIN_NAME) and Check Race"
	cd cmd/admin && go build -race -gcflags="-m -l" -o=../../bin .

clean:
	@echo 'Clean command:'
	@test ! -e bin/$(BIN_NAME) || rm bin/$(BIN_NAME)

run:
	@echo 'Run command:'
	cd bin && ./admin -p=../config

live:
	@echo 'Live Fresh command:'
	cd cmd/admin && go build -gcflags="-m -l" -o=../../bin . && \
	cd ../../bin  && ./admin -p=../config

pprof:
	@echo 'Pprof Web command:'
	cd ./bin && \
	curl -O http://localhost:8081/debug/pprof/profile && \
	go tool pprof -http=:8082 profile

build-linux-best:
	@echo 'Build Linux command:'
	cd cmd/admin && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o=../../bin/admin_linux . && \
	pwd && \
 	cd ../../bin && \
 	upx --best admin_linux -o admin_linux_upx_best && \
 	ls -hl

build-linux-fast:
	@echo 'Build Linux command:'
	cd cmd/admin && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o=../../bin/admin_linux . && \
	pwd && \
 	cd ../../bin && rm admin_linux_upx_fast && \
 	upx -1 admin_linux -o admin_linux_upx_fast && \
 	rm admin_linux && \
 	ls -hl

docker-build:
	@echo 'Docker Build command'
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

docker-run:
	@echo 'Docker Run command:'
	docker run -it -p 8080:8080  $(IMAGE_NAME_TAG)