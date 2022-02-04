# Product Service

This is the Product service

Generated with

```
micro new product
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```


## docker

使用 启动服务
```
docker pull consul
docker run -d -p 8500:8500 consul   // web 访问 127.0.0.1:8500

我的 mac 本机测试
docker start  silly_chebyshev
```

## consule 

registry config

```
key:
/micro/config/mysql

value:
{
	"user":"root",
  "pwd":"12345678",
  "database":"micro",
  "port":"3306",
  "host":"127.0.0.1"
}
```
