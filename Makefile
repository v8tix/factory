include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/temp-srv: runs the cmd app (from code)
.PHONY: run/temp-srv
run/temp-srv:
	@go run ./cmd

## run/temp-srv-log: runs the cmd app (from code) and logs would be written to external files (check the .envrc file to customize).
.PHONY: run/temp-srv-log
run/temp-srv-log:
	@go run ./cmd >>${LINUX_ERROR_LOG_DIR} 2>>${LINUX_INFO_LOG_DIR}

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/srv: builds the cmd app
.PHONY: build/srv
build/srv:
	@echo 'Building ...'
	@GOOS=linux GOARCH=amd64 go build -o=./bin/linux_amd64/${BIN_NAME} ./cmd
	@GOOS=windows GOARCH=amd64 go build -o=./bin/win_amd64/${BIN_NAME} ./cmd
	@echo 'Artifact ready !'

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #

## run/srv-log-linux: runs the linux cmd app and logs would be written to external files (check the .envrc file to customize).
.PHONY: run/srv-log-linux
run/srv-log-linux:
	@echo 'Running ...'
	@./bin/linux_amd64/${BIN_NAME} >>${LINUX_ERROR_LOG_DIR} 2>>${LINUX_INFO_LOG_DIR}

## run/srv-linux: runs the linux cmd app and logs to your console.
.PHONY: run/srv-linux
run/srv-linux:
	@echo 'Running ...'
	@./bin/linux_amd64/${BIN_NAME}

## run/srv-log-win: runs the Windows cmd app and logs would be written to external files (check the .envrc file to customize).
.PHONY: run/srv-log-win
run/srv-log-win:
	@echo 'Running ...'
	@./bin/win_amd64/${BIN_NAME} >>${WIN_ERROR_LOG_DIR} 2>>${WIN_INFO_LOG_DIR}

## run/srv-win: runs the Windows cmd app and logs to your console.
.PHONY: run/srv-win
run/srv-win:
	@echo 'Running ...'
	@./bin/win_amd64/${BIN_NAME}

