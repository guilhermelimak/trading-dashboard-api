default:
	go run main.go

devenv:
	docker-compose run --rm web bash
