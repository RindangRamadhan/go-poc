run:
	go run main.go start
	# @sh ./scripts/go-run.sh
swag-init:
	swag init --parseDependency --parseInternal --parseDepth 1 --overridesFile .swaggo
swag-fmt:
	swag fmt --exclude docs,scripts -g main.go
mock-repositories:
	mockgen -build_flags=--mod=mod -destination=internal/app/repositories/${DESTINATION} -package=mock marketplace/internal/app/repositories ${INTERFACE}
mock-usecase:
	mockgen -build_flags=--mod=mod -destination=internal/app/usecases/${DESTINATION} -package=mock marketplace/internal/app/usecases/${PKG_PATH} ${INTERFACE}                         