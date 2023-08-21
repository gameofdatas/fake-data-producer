# Makefile for Fake Data Producer

# Targets
build:
	docker build -t fake-data-producer .
run-console:
	docker run -it -v $(CONFIG_DIR):/config fake-data-producer console --config-dir /config --file $(CONFIG_FILE) --nr-messages $(NUM_MESSAGES)

clean:
	docker rmi fake-data-producer

.PHONY: run-console
