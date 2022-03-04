
# leaf-go
---
## 目录结构 ##
```
.
├── app               # 应用
├── bootstraps        # 不同应用启动服务不同。
├── main.go           # 入口文件，做一些初始化工作，详见代码注释。
├── Makefile          # 编译脚本，一般不需要改。
├── Makefile.in       # 这个项目特殊的 Makefile 规则。
├── README.md         # 这个文档。
│
├── build/            # GOPATH 指向的目录，Makefile 中使用，不要进行任何修改。
├── client/           # 外部 client。
├── conf/             # 放所有的配置，线上配置应该在文件名中带上shell环境变量GIN_MODE的值，比如 service.conf.release，线上、测试环境分别执行 export GIN_MODE=release 、export GIN_MODE=test
│   └── service.conf  # 开发环境配置文件，所有的配置都应该放在里面
├── dao/              # 持久存储的业务封装。
├── deploy-meta/      # deploy 系统需要用到的配置。
├── model/            # 业务 model 抽象。
├── output/           # 构建的输出目录。
├── controller/       # 访问控制器
└── vendor/           # 由 glide 管理的 vendor 目录。
```





### TODO
- api x
- admin x
- crontab x
- services x