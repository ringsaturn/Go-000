# Error 处理

<https://u.geekbang.org/lesson/68?article=314039>

- [`main.go`](main.go)
- 蜜桃成熟时33D
- 喜爱夜蒲
- wrap 时机

  1. 系统级别异常
  2. 基础库
  3. 第三方库

- 业务代码做日志
- 野生 goroutine 尽量额外包一层

  ```go
  func Go(trueFunc func()) {
  	go func() {
  		defer func() {
  			if err := recover(); err != nil {
  				fmt.Println(err)
  			}
  		}()
  		trueFunc()
  	}()
  }
  ```

- panic 时机
  - 初始化
  - 配置解析之类的
  - 读多写少的任务能接受数据库不可用而缓存可用的情况（ MySQL + Redis）
- 业务代码尽量不 panic，而是处理 error
- 数据库查询是没找到，返回 err 还是 nil
  1. 需要在内存里做进一步处理的，如处理订单的，直接抛 error（反正也处理不了）
  2. 查询后直接返回的，返回空 **也许** 可以
