run:
	@echo "running..."
	./generator-app --number-of-groups=1000000 --batch-size=5000 --interval=1 --output-directory="~/Desktop"

build:
	@echo "Building executable file"
	go build

clean:
	@echo "Cleaning up..."
	rm *.json
	rm generator-app
	