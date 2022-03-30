###toruk  （托鲁克） 阿凡达电影里面最后一只上古神兽




## 目录结构

```
├── Dockerfile  
├── Makefile  
├── README.md
├── api // 可以放swagger 校验工具// 对外暴露接口
│  
├── cmd  // 整个项目启动的入口文件
│   └── server
│       ├── main.go
├── configs  // 这里通常维护一些本地调试用的样例配置文件
│   └── config.ini
├── go.mod
├── go.sum
├── internal  // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│   ├── biz   // 业务逻辑的组装层，类似 DDD 的 domain (php 里面的controllers  ，zh-ser-ugc里面的manager) 层，data 类似 DDD 的 repo
│   │   ├── README.md
│   │   └── biz.go
│   ├── conf  // 内部使用的config的结构定义，使用proto格式生成
│   │   ├── conf.pb.go
│   │   └── conf.proto
│   ├── data  // 业务数据访问，包含 cache、db 等封装， (php 里面的modle   zh-ser-ugc 里面  dao 和model 的结合)
│   │   ├── README.md
│   │   └──data.go
│   ├── server  // http和grpc实例的创建和配置
│   │   
│   └── service  // 实现了 api 定义的服务层，这里不应处理复杂逻辑（活动小说游戏的endpoint层） （zongheng-server-ugc   里面的remote）
│       └── README.md
└── third_party  // api 依赖的第三方proto
└── README.md
```