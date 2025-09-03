# ellie-layout

├── application             // 应用编排层，调用domain层执行业务逻辑
│   ├── convert             // 数据转换，转换dto和entity
│   └── dto                 // 数据传输对象定义
├── domain                  // 领域层
│   ├── biz                 // 业务逻辑
│   ├── cache               // 缓存接口定义
│   ├── entity              // 领域实体
│   ├── event               // 领域事件接口定义
│   └── repo                // 领域仓储接口定义
├── iface                   // 接口层，定义对外接口，负责内外部数据转换，调用application层
├── infra                   // 基础设施层，包含对外部资源的访问实现，如数据库、缓存、消息队列等
│   ├── impl                // 基础设施实现，实现领域层定义的接口
│   │   ├── cacheimpl       // 缓存实现
│   │   ├── eventimpl       // 领域事件实现
│   │   └── repoimpl        // 领域仓储实现
│   │       ├── migration   // 数据库迁移脚本
│   │       └── model       // 数据库模型定义
│   └── rpc                 // RPC客户端实现，封装对外部服务的调用
└── server                  // 启动服务相关代码，如HTTP服务器、GRPC服务器等
