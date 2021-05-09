## 快速入门

```
micro new user
make proto
micro run .
```

## 指导
1. 先创建 *.proto 文件
2. 在 domain 中创建model
3. 在 domain 中创建 repository,实现与model交互的方法(不含业务逻辑)
4. 在 domain 中创建 service,实现与repository的交互(包含业务逻辑,hook处理等)
5. 在 handler 中创建 *.micro.go的 *handle(如 UserHandler) 的接口实现
    并绑定 domain的service的实现