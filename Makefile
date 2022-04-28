GOPATH:=$(shell go env GOPATH)

.PHONY: init_env
init_env:
	docker run -d -p 9002:9002 --name hystrix-dashboard mlabouardy/hystrix-dashboard:latest
	docker run -d -p 8700:8500 --name silly_chebyshev  consul

.PHONY: run_env
run_env:
	docker start silly_chebyshev #consul
	docker start hystrix-dashboard	#hystrix

	docker exec -t  silly_chebyshev consul  kv put 'micro/config/mysql'  '{ "user":"root", "pwd":"12345678", "database":"micro", "port":"3306", "host":"127.0.0.1" }'

.PHONY:run_cart
run_cart:
	cd cart && micro run .

.PHONY:run_cart_api
run_cart_api:
	cd cartApi && micro run .

.PHONY: run_user
run_user:
	cd user && micro .

