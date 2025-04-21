compare_steam_libs:
	go run cmd/compare_steam_libs/main.go

get_steam_api_list:
	go run cmd/get_steam_api_list/main.go

list_steam_apps:
	go run cmd/list_steam_apps/main.go

test:
	go test ./...

mod_tidy:
	go mod tidy