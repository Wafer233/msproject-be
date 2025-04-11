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
│   ├── infrastructure/         # 基础设施代码
│   ├── domain/                 # 共享领域对象
│   └── application/            # 共享应用服务
│
├── user-service/               # 用户服务
│   ├── cmd/                    # main函数放的地方
│   ├── internal/               # 核心代码
│   │   ├── domain/             # DDD设计之领域层 
│   │   │   ├── model/          # 模型，存值对象和实体，但是感觉这两个东西基本上是一样的，唯一的区别似乎是不带ID
│   │   │   ├── repository/     # infra里头的repo的接口
│   │   │   ├── service/        # 用不到dto的service
│   │   │   └── event/          # 
│   │   ├── application/        # DDD设计之应用层
│   │   │   ├── service/        # 用了dto的service
│   │   │   ├── dto/            # 简单的理解就是带json tag的结构体
│   │   │   └── command/        # 
│   │   ├── infrastructure/     # DDD设计之基础设施层
│   │   │   ├── repository/     # 这个重名了，所以调用的时候我一般用impl，即上面的domain的实现
│   │   │   ├── persistence/    # 持久化，存mysql
│   │   │   └── messaging/      # 
│   │   └── interfaces/         # DDD设计之接口层
│   │       ├── grpc/           # gRPC接口
│   │       ├── rest/           # REST接口
│   │       └── event/          # 
│   ├── proto/                  # 协议定义
│   └── config/                 # 配置
│
├── api-gateway/                # API网关
│   ├── cmd/                    # 应用入口
│   ├── internal/               # 内部代码结构
│   │   ├── application/        # 应用层
│   │   │   ├── service/        # 聚合服务
│   │   │   └── dto/            # 数据传输对象
│   │   ├── infrastructure/     # 基础设施层
│   │   │   ├── grpc/           # gRPC客户端
│   │   │   └── middleware/     # 中间件
│   │   └── interfaces/         # 接口层
│   │       └── rest/           # REST接口
│   └── config/                 # 配置
│
├── go.work                     # Go工作区配置
└── docker/                     # Docker配置
```
