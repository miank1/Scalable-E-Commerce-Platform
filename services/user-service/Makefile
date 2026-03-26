APP_NAME=user-service

build:
	podman build -t $(APP_NAME) .

run:
	podman run --rm -it --env-file .env -p 8080:8080 $(APP_NAME)

dev:
	go run main.go

logs:
	podman logs -f $(shell podman ps -q --filter ancestor=$(APP_NAME))

stop:
	podman stop -a

clean:
	podman rm -f $(shell podman ps -aq)

restart: stop build run