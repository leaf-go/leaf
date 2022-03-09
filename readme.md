# leaf-go
---

## 目录结构 ##

```
.
├── apps              # 应用目录 
├── boot              # 启动初始化、注册应用。
├── config            # 配置文件。
├── data              # 数据文件、model、logic、cache... 自行添加。
│
├── e/                # errors 简写命名防止与errors冲突
├── middleware/       # 中间件。
├── mounts/           # 挂载目录 包括应用结构与自定义结构 
│   └── service.conf  # 开发环境配置文件，所有的配置都应该放在里面
├── params/           # 参数结构体 用于获取参数验证参数。
├── routes/           # 路由。
├── utils/            # 工具类 后续单独封装
└── vendor/           # 由 glide 管理的 vendor 目录。
```

