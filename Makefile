# Run the code locally 
run:
	go run main.go
#build the docker container. 
buildc:
	docker build -t testimage .