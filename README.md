# msproject-be

This was originally a file from a go training camp. Its [front end](https://github.com/Wafer233/msproject-fe) was written in vue. After I built it, I found that its front end was unexpectedly good. Here I diss my previous project that didn't even have a go test written in the front end, and the interface documents were quite complete, so I decided to use this project.

# Improvement
The original author's structure was incomprehensible. He wrote in his doc that he designed it according to DDD, but it was basically a mess. Unable to bear it, I redesigned the project structure and redesigned the structure using the dependency injection I learned in [webook](https://github.com/Wafer233/webook).
I personally feel that it basically strictly complies with the design concept of DDD. Of course, since I am not a professional, I consulted teachers Claude and GPT in many cases ^^

# Project structure
It is roughly as follows.

```
msproject-be/
├── common/                     # Shared infrastructure and utilities
│   ├── logs/                   # Logging utilities
│   ├── errs/                   # Error handling
│   └── validate.go             # Validation utilities
│
├── user-service/               # User service
│   ├── cmd/                    # Application entry point
│   ├── config/                 # Configuration files
│   ├── internal/               # Internal code
│   │   ├── domain/             # Domain layer 
│   │   │   ├── model/          # Domain models (entities and value objects)
│   │   │   ├── repository/     # Repository interfaces
│   │   │   └── service/        # Domain services
│   │   ├── application/        # Application layer
│   │   │   ├── service/        # Application services
│   │   │   └── dto/            # Data Transfer Objects
│   │   ├── infrastructure/     # Infrastructure layer
│   │   │   ├── repository/     # Repository implementations
│   │   │   ├── persistence/    # Persistence mechanisms
│   │   │   └── cache/          # Cache implementations
│   │   └── interface/          # Interface layer
│   │       └── grpc/           # gRPC interfaces
│   └── proto/                  # Protocol definitions
│
├── project-service/            # Project service
│   ├── cmd/                    # Application entry point
│   ├── config/                 # Configuration files
│   ├── internal/               # Internal code
│   │   ├── domain/             # Domain layer
│   │   ├── application/        # Application layer
│   │   ├── infrastructure/     # Infrastructure layer
│   │   └── interface/          # Interface layer
│   └── proto/                  # Protocol definitions
│
├── api-gateway/                # API Gateway
│   ├── cmd/                    # Application entry point
│   ├── config/                 # Configuration files
│   ├── internal/               # Internal code
│   │   ├── application/        # Application layer
│   │   │   ├── service/        # Services
│   │   │   └── dto/            # Data Transfer Objects
│   │   ├── domain/             # Domain layer
│   │   │   └── model/          # Domain models
│   │   ├── infrastructure/     # Infrastructure layer
│   │   │   └── grpc/           # gRPC clients
│   │   └── interfaces/         # Interface layer
│   │       └── rest/           # REST API interfaces
│   └── proto/                  # Protocol definitions
│
├── go.work                     # Go workspace configuration
└── README.md                   # Project documentation
```
# msproject-be

这个原本是一个go训练营的文件，他的[前端](https://github.com/Wafer233/msproject-fe)是vue写的，我当时构建了之后发现他的前端意外的还不错，这里diss一下我以前那个连前端都没有硬写go test的项目，而且接口文档之类的写的挺全的，就决定使用这个项目了。

# 改进
原作者的结构写的不知所云，他在他的doc上写着自己是按照DDD设计的，其实基本上写的乱成一团。无法忍受之下，我就重新设计了项目结构，并且用之前在[webook](https://github.com/Wafer233/webook)学到的依赖注入重新设计了一下结构，
个人感觉基本上是严格的遵守了DDD的设计理念。当然由于我确实不是专业人士，很多情况下都请教了一下claude老师和gpt老师 ^^

# 项目结构
大体如下
```
msproject-be/
├── common/                     # 共享基础设施和工具
│   ├── logs/                   # 日志工具
│   ├── errs/                   # 错误处理
│   └── validate.go             # 验证工具
│
├── user-service/               # 用户服务
│   ├── cmd/                    # 应用入口
│   ├── config/                 # 配置文件
│   ├── internal/               # 内部代码
│   │   ├── domain/             # 领域层 
│   │   │   ├── model/          # 领域模型（实体与值对象）
│   │   │   ├── repository/     # 仓储接口定义
│   │   │   └── service/        # 领域服务
│   │   ├── application/        # 应用层
│   │   │   ├── service/        # 应用服务
│   │   │   └── dto/            # 数据传输对象
│   │   ├── infrastructure/     # 基础设施层
│   │   │   ├── repository/     # 仓储实现
│   │   │   ├── persistence/    # 持久化
│   │   │   └── cache/          # 缓存实现
│   │   └── interface/          # 接口层
│   │       └── grpc/           # gRPC接口
│   └── proto/                  # 协议定义
│
├── project-service/            # 项目服务
│   ├── cmd/                    # 应用入口
│   ├── config/                 # 配置文件
│   ├── internal/               # 内部代码
│   │   ├── domain/             # 领域层
│   │   ├── application/        # 应用层
│   │   ├── infrastructure/     # 基础设施层
│   │   └── interface/          # 接口层
│   └── proto/                  # 协议定义
│
├── api-gateway/                # API网关
│   ├── cmd/                    # 应用入口
│   ├── config/                 # 配置文件
│   ├── internal/               # 内部代码
│   │   ├── application/        # 应用层
│   │   │   ├── service/        # 服务
│   │   │   └── dto/            # 数据传输对象
│   │   ├── domain/             # 领域层
│   │   │   └── model/          # 领域模型
│   │   ├── infrastructure/     # 基础设施层
│   │   │   └── grpc/           # gRPC客户端
│   │   └── interfaces/         # 接口层
│   │       └── rest/           # REST API接口
│   └── proto/                  # 协议定义
│
├── go.work                     # Go工作区配置
└── README.md                   # 项目文档
```
