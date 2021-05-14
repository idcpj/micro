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
