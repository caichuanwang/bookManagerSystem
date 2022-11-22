# book-mananger-system

​    book-manager-system是一款用于图书管理的中后台系统。系统初始化极度简单，只需要配置文件中，修改数据库连接，可以使用docker一键部署应用

[前端项目](https://github.com/caichuanwang/bms-fe)

## 🐯技术栈

- [echo](https://echo.labstack.com/)

- [Ant design](https://ant.design/components/overview-cn)

- [React](https://reactjs.org/)

  

## 🔥页面展示

![image-20221122161036195](http://gicgo-images.oss-cn-shanghai.aliyuncs.com/img/image-20221122161036195.png)

![image-20221122161102987](http://gicgo-images.oss-cn-shanghai.aliyuncs.com/img/image-20221122161102987.png)

![image-20221122161123307](http://gicgo-images.oss-cn-shanghai.aliyuncs.com/img/image-20221122161123307.png)



## ✨ 特性

- 遵循 RESTful API 设计规范
- 基于 ECHO WEB API 框架，提供了丰富的中间件支持（用户认证、访问日志等）
- 基于Casbin的 RBAC 访问控制模型
- JWT 认证
- 支持 Swagger 文档(基于swaggo)
- 基于 GORM 的数据库存储
- 配置文件简单的模型映射，快速能够得到想要的配置
- TODO: 单元测试



## 🎁 内置

1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2. 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
3. 参数管理：对系统动态配置常用参数。
4. 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
5. 登录日志：系统登录日志记录查询包含登录异常。
6. 接口文档：根据业务代码自动生成相关的api接口文档。
7. 定时任务：自动化任务，目前支持接口调用和函数调用。
8. 邮件通知：可以配置邮件，不会错过任何信息。

## 准备工作

你需要在本地安装 [go] [mysql] [redis] [node](http://nodejs.org/) 和 [git](https://git-scm.com/)

## 📦 本地开发

### 环境要求

go 1.18

node版本: v16.17.0

npm版本: 8.15.0

redis版本：6.0.6



### 开发目录创建

```
# 创建开发目录
mkdir bookManagerSystem
cd bookManagerSystem
```

### 获取代码

```
# 获取后端代码
git clone https://github.com/caichuanwang/bookManagerSystem.git

# 获取前端代码
git clone https://github.com/caichuanwang/bms-fe.git
```

### 启动说明

#### 服务端启动说明

```
# 进入 bookManagerSystem 后端项目
cd ./bookManagerSystem

# 更新整理依赖
go mod tidy

# 编译项目
go build

# 修改配置 
# 文件路径  bookManagerSystem/app.conf
vi bookManagerSystem/app.conf
```

:::tip ⚠️注意 在windows环境如果没有安装中CGO，会出现这个问题；

```
E:\bookManagerSystem>go build
# github.com/mattn/go-sqlite3
cgo: exec /missing-cc: exec: "/missing-cc": file does not exist
```

or

```
D:\Code\bookManagerSystem>go build
# github.com/mattn/go-sqlite3
cgo: exec gcc: exec: "gcc": executable file not found in %PATH%
```

[解决cgo问题进入](https://doc.go-admin.dev/zh-CN/guide/faq#cgo-的问题)

:::

#### 初始化数据库，以及服务启动

```
# 首次配置需要初始化数据库资源信息
book_management_system.sql
```

#### 使用docker 编译启动

```
# 编译镜像
docker build -t bookManagementSystem .

# 启动容器，第一个bookManagementSystem是容器名字，第二个bookManagementSystem是镜像名称
# -v 映射配置文件 本地路径：容器路径
docker run --name bookManagementSystem -p 8888:8888 -v /app.conf:/app.conf -d bookManagementSystem
```

#### 文档生成

```
swag init
```

#### 交叉编译

```
# windows
env GOOS=windows GOARCH=amd64 go build main.go

# or
# linux
env GOOS=linux GOARCH=amd64 go build main.go
```

### UI交互端启动说明

```
# 安装依赖
yarn install

# 启动服务
yarn start
```



## 🔑 License

[MIT](https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md)

Copyright (c) 2022 caichuanwang
