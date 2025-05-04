BUILD_DIR=build
CMD_DIR=src/cmd
.PHONY: build run clean

build: 
	mkdir -p $(BUILD_DIR)
	chmod 777 $(BUILD_DIR)
	go build -o $(BUILD_DIR)/main $(CMD_DIR)/main.go

run: build
	./$(BUILD_DIR)/main

clean: 
	rm -rf $(BUILD_DIR) 

.PHONY: format

format:
		gofmt -w .