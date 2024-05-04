.PHONY: build-test
build-test:
	go build 
	go test

.PHONY: run
run:
	go build 
	./testing-actions

.PHONY: docker-build
docker-build:
	docker build -t github.com/sachin-raja-sap/testing-actions -f Dockerfile .

.PHONY: docker-run
docker-run:
	docker run --rm -p 8000:8000 github.com/sachin-raja-sap/testing-actions