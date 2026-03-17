# Makefile for Server ZZQ

.PHONY: build run test clean help deps fmt lint

# 变量
BINARY_NAME=server
BINARY_PATH=bin/$(BINARY_NAME)
MAIN_PATH=cmd/server/main.go
CONFIG_PATH=configs/config.yaml

# 默认目标
all: build

# 安装依赖
deps:
	go mod download
	go mod tidy

# 格式化代码
fmt:
	go fmt ./...

# 代码检查
lint:
	golangci-lint run

# 编译
build:
	go build -o $(BINARY_PATH) $(MAIN_PATH)

# 运行
run:
	go run $(MAIN_PATH)

# 测试
test:
	go test -v ./...

# 测试覆盖率
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# 清理
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean

# 运行迁移
migrate:
	@echo "Running database migrations..."
	# 添加迁移命令

# 帮助
help:
	@echo "可用命令:"
	@echo "  make deps     - 安装依赖"
	@echo "  make fmt      - 格式化代码"
	@echo "  make lint     - 代码检查"
	@echo "  make build    - 编译项目"
	@echo "  make run      - 运行开发服务器"
	@echo "  make test     - 运行测试"
	@echo "  make clean    - 清理构建文件"
	@echo "  make help     - 显示帮助"
