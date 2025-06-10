# Makefile
.PHONY: build

BINARY_NAME=taxi-pudasjarvi

# Build the app in webserver mode
build/server:
	go mod tidy && \
	templ generate && \
	go generate && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}

build/static:
	go run github.com/a-h/templ/cmd/templ@latest generate && \
	go mod tidy && \
	go generate && \
	npx --yes @tailwindcss/cli -i ./tailwind.css -o ./assets/tailwind.css && \
	go build -ldflags="-w -s" -o ${BINARY_NAME} && \
	./${BINARY_NAME} static

# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
dev/templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

# run air to detect any go file changes to re-build and re-run the server.
dev/server:
	air \
	-d \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
dev/tailwind:
	npx --yes @tailwindcss/cli -i ./tailwind.css -o ./assets/tailwind.css --minify --watch

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
dev/sync_assets:
	air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "assets" \
	--build.include_ext "css"

# start all 5 watch processes in parallel.
dev: 
	make -j4 dev/tailwind dev/server dev/templ dev/sync_assets