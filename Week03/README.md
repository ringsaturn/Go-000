# Concurrency

- 锁
    - Mutex：性能开销会很大
    - RWMutex：相对提升一些
    - Atmic.Value：很「清真」
- Redis COW 实现
- 临界区的代码不应当很复杂，比如只做 +1 的操作
- select 的顺序不要做假设
- ErrGroup
- CPU 密集型，没有做超时控制
- IO 密集，用 context 来托管超时
    - 某些特殊场景下 IO 密集型高频系统调用应当避免，比如 zap 时间戳的获取
        - <https://blog.laisky.com/p/go-fluentd/>
- 空 interface 判断的坑

