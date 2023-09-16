# Variables
IMAGE_NAME = central

# Comandos
build:
	sudo docker build . -t $(IMAGE_NAME) 

docker-central:
	sudo docker run $(IMAGE_NAME)