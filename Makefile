build:
	mkdir -p bin
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go

deploy_prod: build
	npx serverless deploy --stage prod
