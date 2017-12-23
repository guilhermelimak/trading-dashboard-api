default:
	go run *.go

devenv:
	docker-compose run --rm web bash
