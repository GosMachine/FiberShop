run:
	@docker-compose up -d
	@npx tailwindcss build -i web/static/styles/styles.css -o web/static/styles/output.css --minify
	@templ generate
	@go run cmd/app/main.go -config=configs/config.yaml