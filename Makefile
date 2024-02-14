
start: templ tailwind
	go run cmd/main.go
templ:
	templ generate
tailwind:
	npx tailwindcss -i ./input.css -o ./static/css/output.css 

.PHONY: start templ tailwind	
