run:
	rm -rf .env && chmod +x generate_env.sh && ./generate_env.sh && docker-compose up -d

down:
	docker-compose down