CONF_PATH=

env.start:
	docker compose -f docker/local.yaml up -d
env.stop:
	docker compose -f docker/local.yaml down
run:
	CONF_PATH=$(shell pwd)/config/local.yaml go run main.go Ivan John Hans
