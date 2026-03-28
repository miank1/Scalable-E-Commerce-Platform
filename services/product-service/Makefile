APP_NAME=product-service
PORT=8081

build:
	podman build -t $(APP_NAME) .

run:
	podman run -d \
	--name product-service \
	--network ecommerce-net \
	-p 8081:8081 \
	product-service

dev:
	go run main.go

logs:
	podman logs -f $(shell podman ps -q --filter ancestor=$(APP_NAME))

stop:
	podman stop -a

clean:
	podman rm -f $(shell podman ps -aq)

restart: stop build run