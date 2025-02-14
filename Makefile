# 项目名
PROJECT_NAME := mikumiku_clouddrive

# 源代码目录
SRC_DIR := .

# 输出目录
OUT_DIR := ./bin

# Go 编译器
GO := go

# 目标平台
PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64

# 默认目标
.PHONY: all
all: clean build

# 清理
.PHONY: clean
clean:
	rm -rf $(OUT_DIR)

# 创建输出目录
.PHONY: create-out-dir
create-out-dir:
	mkdir -p $(OUT_DIR)

# 构建
.PHONY: build
build: create-out-dir $(PLATFORMS)

# 针对每个平台编译
$(PLATFORMS):
	GOOS=$(word 1, $(subst /, ,$@)) GOARCH=$(word 2, $(subst /, ,$@)) \
	$(GO) build -o $(OUT_DIR)/$(PROJECT_NAME)-$(word 1, $(subst /, ,$@))-$(word 2, $(subst /, ,$@))$(if $(findstring windows,$@),.exe) $(SRC_DIR)

# 运行项目
.PHONY: run
run:
	$(GO) run $(SRC_DIR)

# 格式化代码
.PHONY: format
format:
	gofumpt -w .

# 生成swagger文档
.PHONY: doc
doc:
	swag init -g ./main.go -o docs

# 初始化数据库
.PHONY: initdb
initdb:
	go run main.go --initdb

# 测试
.PHONY: test
test:
	$(GO) test ./...
	
# 安装依赖
.PHONY: deps
deps:
	$(GO) mod tidy
	$(GO) install mvdan.cc/gofumpt@latest
	$(GO) install github.com/swaggo/swag/cmd/swag@latest

# 使用方法
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make            - 编译所有平台的可执行文件"
	@echo "  make clean      - 清理输出目录"
	@echo "  make build      - 编译所有平台的可执行文件"
	@echo "  make run        - 编译并运行项目"
	@echo "	 make format     - 格式化代码"
	@echo "  make doc		 - 生成swagger文档"
	@echo "  make test       - 运行测试"
	@echo "  make deps       - 安装依赖"
	@echo "  make initdb     - 初始化数据库"
	@echo "  make help       - 显示此帮助信息"
