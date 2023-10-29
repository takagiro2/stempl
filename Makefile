templ:
	cd ./src/components
	templ generate
	cd ../../
	mv ./src/components/*.go ./bin/components

clean:
	rm ./bin/components/*.go

serv:
	go run ./server/main.go