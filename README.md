# MikuMikuCloudDrive

[![Go Version](https://img.shields.io/badge/go-1.20+-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/gin-1.9.0-green.svg)](https://github.com/gin-gonic/gin)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

MikuMikuCloudDrive 是一个基于 Gin 框架开发的云存储系统后端，提供文件上传、下载、目录管理等核心功能。

## 功能特性

- 用户认证与授权 (JWT)
- 文件分块上传与断点续传
- 文件下载
- 目录管理
  - 创建目录
  - 删除目录
  - 重命名目录
  - 获取目录内容
- 用户管理
  - 用户注册/登录/注销
  - 用户信息管理
- Swagger API 文档
- Redis 缓存支持
- 日志记录

## 技术栈

- 编程语言: Go 1.20+
- Web 框架: Gin
- 数据库: MySQL
- 缓存: Redis
- 认证: JWT
- API 文档: Swagger

## 安装与运行

### 前置要求

- Go 1.20+
- MySQL 5.7+
- Redis 6.0+

### 安装步骤

1. 克隆仓库

   ```bash
   git clone https://github.com/yourusername/MikuMikuCloudDrive.git
   cd MikuMikuCloudDrive
   ```

2. 配置环境

   - 复制配置文件模板
     ```bash
     cp config/config.toml.example config/config.toml
     ```
   - 修改 `config/config.toml` 中的数据库和 Redis 配置

3. 安装依赖

   ```bash
   go mod tidy
   ```

4. 初始化数据库

   ```bash
   go run main.go --initdb
   ```

5. 运行项目

   ```bash
   go run main.go
   ```

6. 访问 API 文档
   ```
   http://localhost:8888/swagger/index.html
   ```

## 数据库结构

![](https://blog.meowrain.cn/api/i/2025/01/15/2MfTvn1736939923222826046.avif)

## API 文档

项目使用 Swagger 自动生成 API 文档，启动项目后访问：

```
http://localhost:8888/swagger/index.html
```

## 项目结构

```
MikuMikuCloudDrive/
├── common/            # 通用组件
├── config/            # 配置文件
├── controllers/       # 控制器
├── core/              # 核心初始化
├── docs/              # Swagger 文档
├── middleware/        # 中间件
├── models/            # 数据模型
├── routes/            # 路由配置
├── services/          # 业务逻辑
├── types/             # 类型定义
├── utils/             # 工具类
├── go.mod
├── go.sum
├── main.go            # 入口文件
└── README.md
```

## TODO

- [ ] 回收站功能
- [ ] 文件分享功能
- [ ] 文件版本控制
- [ ] 文件搜索功能
- [ ] 用户配额管理

## 许可证

本项目采用 MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件
