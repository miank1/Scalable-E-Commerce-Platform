.PHONY: build run-product run-user run-gateway run-all dev-product dev-user dev-gateway stop clean restart

# ==============================
# PATHS
# ==============================
USER_SERVICE=services/user-service
PRODUCT_SERVICE=services/product-service
GATEWAY=gateway

USER_BIN=user-service.exe
PRODUCT_BIN=product-service.exe
GATEWAY_BIN=gateway.exe

# ==============================
# BUILD
# ==============================
build:
	@echo Building services...
	cd $(PRODUCT_SERVICE) && go build -o $(PRODUCT_BIN)
	cd $(USER_SERVICE) && go build -o $(USER_BIN)
	cd $(GATEWAY) && go build -o $(GATEWAY_BIN)
	@echo Build complete

# ==============================
# RUN SERVICES (BINARY MODE)
# ==============================

run-product:
	@echo Starting product-service (8081)...
	cd $(PRODUCT_SERVICE) && $(PRODUCT_BIN)

run-user:
	@echo Starting user-service (8080)...
	cd $(USER_SERVICE) && $(USER_BIN)

run-gateway:
	@echo Starting gateway (8000)...
	cd $(GATEWAY) && $(GATEWAY_BIN)

run-all:
	@echo Starting all services...
	cd $(PRODUCT_SERVICE) && start /B $(PRODUCT_BIN)
	timeout /t 2 >nul
	cd $(USER_SERVICE) && start /B $(USER_BIN)
	timeout /t 2 >nul
	cd $(GATEWAY) && start /B $(GATEWAY_BIN)
	@echo All services started

# ==============================
# DEV MODE (go run)
# ==============================
dev-product:
	cd $(PRODUCT_SERVICE) && go run main.go

dev-user:
	cd $(USER_SERVICE) && go run main.go

dev-gateway:
	cd $(GATEWAY) && go run main.go

# ==============================
# STOP
# ==============================
stop:
	taskkill /IM $(PRODUCT_BIN) /F >nul 2>&1 || echo product not running
	taskkill /IM $(USER_BIN) /F >nul 2>&1 || echo user not running
	taskkill /IM $(GATEWAY_BIN) /F >nul 2>&1 || echo gateway not running

# ==============================
# CLEAN
# ==============================
clean:
	-del /Q $(PRODUCT_SERVICE)/$(PRODUCT_BIN)
	-del /Q $(USER_SERVICE)/$(USER_BIN)
	-del /Q $(GATEWAY)/$(GATEWAY_BIN)

# ==============================
# RESTART
# ==============================
restart: stop build run-all