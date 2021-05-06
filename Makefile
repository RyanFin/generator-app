makefile:
all: build run test

run:
	@echo "running..."
	./generator-app --number-of-groups=1000000 --batch-size=5000 --interval=1 --output-directory="Desktop/"

build:
	@echo "Building executable file"
	go build

clean:
	@echo "Cleaning up..."
	rm *.json
	rm generator-app

test:
	@echo "Testing Application - main"
	go test -v -coverprofile=coverage.out
	go tool cover -html=coverage.out
	@echo "Testing Application - /pkg"
	cd pkg && go test -v -coverprofile=coverage.out
	cd pkg && go tool cover -html=coverage.out

	