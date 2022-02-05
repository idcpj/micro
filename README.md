## goimooc

## user
1. 人员的基本操作
2. 只涉及最基础的微服务操作

## category
1. 对商品分类的操作
2. 并添加了配置中心与注册中心的代码

## product
1. 对产品的基本操作
2. 涉及到一对多表的 orm 操作

## cart 
购物车

1. 在 cart 中  提出共同code 到 common
2. 添加链路追踪
3. 添加限流

微服务API网关
路径说明
- 通过网关请求`/greeter/say/hello`这个路径,网关会将请求转发到go.micro.api.greeter服务的Say.Hello 方法处理
- go.micro.api 是网关的默认服务器前缀
- 路径 `/cartApi/cartApi/findAll` 也可以写成`/cartApi/findAll`