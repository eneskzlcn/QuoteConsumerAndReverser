build:
	go build -o bin/main main.go utility.go
clean:
	rm -R bin
run:
	bin/main