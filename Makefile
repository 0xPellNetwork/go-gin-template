.PHONY: run build test clean install

# 运行项目
run: docs
	go run cmd/server/main.go

# 构建项目
build:
	go build -o bin/gin-template cmd/server/main.go

# 运行测试
test:
	go test ./test/... -v

# 运行测试并显示覆盖率
test-cover:
	go test ./test/... -v -cover

# 安装依赖
install:
	go mod download
	go mod tidy

# 清理构建文件
clean:
	rm -rf bin/
	go clean

# 格式化代码
fmt:
	go fmt ./...

# 代码检查
lint:
	golangci-lint run

# 开发模式（自动重载）
dev:
	air

# 生成 Swagger API 文档
docs:
	swag init -g cmd/server/main.go -o docs

# 格式化 Swagger 注释
docs-fmt:
	swag fmt -g cmd/server/main.go

# Docker 构建
docker-build:
	docker build -t gin-template .

# Docker 运行
docker-run:
	docker run -p 8080:8080 gin-template 