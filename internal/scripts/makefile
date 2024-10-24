SRC=../../cmd/main.go 
CLIENT=../../client/client.go 


RED=\033[31m
GREEN=\033[32m
YELLOW=\033[33m
BLUE=\033[34m
RESET=\033[0m

.PHONY: run 

run:
	@echo "$(YELLOW)Application is starting...$(RESET)"
	go run $(SRC) & # Run the server in the background
	go run $(CLIENT) # Run the client

	@echo "$(GREEN)Application is running...$(RESET)"
