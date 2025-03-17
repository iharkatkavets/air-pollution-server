WEB_PORT=4000
API_PORT=4001
DB="api.db"
ENV="development"

## build: builds all binaries
build: clean build_front build_api
	@printf "All binaries built!\n"

## clean: cleans all binaries and runs go clean
clean:
	@echo "Cleaning..."
	@- rm -f dist/*
	@go clean
	@echo "Cleaned!"

## build_front: builds the front end
build_front:
	@echo "Building front end..."
	@go build -o dist/web-server ./cmd/web
	@echo "Front end built!"

## build_api: builds the back end
build_api:
	@echo "Building back end..."
	@go build -o dist/api-server ./cmd/api
	@echo "Back end built!"

release_api:
	@echo "Building api release..."
	@CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=musl-gcc go build -tags "sqlite_omit_load_extension" -ldflags "-extldflags -static" -o dist/api-server ./cmd/api/
	@echo "Back end built!"

## start: starts front and back end
start: start_front start_api
	
## start_front: starts the front end
start_front: build_front
	@echo "Starting the front end..."
	@env ./dist/web-server -port=${WEB_PORT} &
	@echo "Front end running!"

## start_api: starts the back end
start_api: build_api
	@echo "Starting the back end..."
	@env ./dist/api-server -port=${API_PORT} -db=${DB} -env=${ENV}&
	@echo "Back end running!"

## stop: stops the front and back end
stop: stop_front stop_api
	@echo "All applications stopped"

## stop_front: stops the front end
stop_front:
	@echo "Stopping the front end..."
	@-pkill -SIGTERM -f "web-server -port=${WEB_PORT}"
	@echo "Stopped front end"

## stop_api: stops the back end
stop_api:
	@echo "Stopping the back end..."
	@-pkill -SIGTERM -f "api-server -port=${API_PORT}" -db=${DB} -env=${ENV}
	@echo "Stopped back end"

