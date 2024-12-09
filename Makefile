run:
	go run cmd/api/main.go

mig-new:
	migrate create -ext sql -dir migrations -seq $(NAME)


MAKEFILE_DB_URL=postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_URL):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable

action=up
mig:
	@migrate -path migrations -database $(MAKEFILE_DB_URL) $(action)
