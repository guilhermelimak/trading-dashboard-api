default:
	go run main.go

devenv:
	docker-compose run -p 8888:8888 --rm web bash
