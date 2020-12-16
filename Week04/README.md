# Pkg

- 建立统一的 kit 包
- Package oriented design
- `/api` 存储所有的 Probuf 文件，公用
- YAML/JSON/TOML -> Probuf 作为配置
- Go 编译会忽略 `.`/`_` 开头的文件
- 服务树，service tree
- DNS 形式的 AppID `account.service.vip` 用于服务注册
  - `cmd` 目录支部则启动，初始化，关闭，listen 这种逻辑
- dao 面向表，数据库访问层，返回 model 结构体
- 关注 kratos v2 的代码
- model
  - 失血模型
  - 贫血模型（比较好）
  - 充血模型
- 控制反转
- 依赖倒置
- service 关注 gRPC 服务发现，编排
- gRPC 尽量不用 empty 对象，不好扩展
- Golang 的默认值实现
- B 站 App 埋点上报用的是 probuf 形式
- **全局错误码不可靠**
- 不要将业务的错误代码和 gRPC 耦合，deepcopy 到对象中
  - DTO 复制思路
  - service error -> gRPC error -> service error
- Probuf field mask 标志哪些字段更新
- 配置：环境变量、静态配置、动态配置、全局配置
- `sync.Once` 保证测试用的数据库初始化只执行一次，但不利用清理；Golang 后来已经补上了测试后清理的功能
- YAPI 测试

Books:

- [x] Google 测试之道
- [ ] 微软的软件测试之道
- [ ] 架构整洁之道
