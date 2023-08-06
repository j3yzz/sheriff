default: help

define HELP_MESSAGE
  @echo "Usage:"
  @echo "  create-migration:"
  @echo "    - make create_migration table_name=<table_name>"
endef

export HELP_MESSAGE

help:
	$(HELP_MESSAGE)


run:
	rm -rf .env && chmod +x generate_env.sh && ./generate_env.sh && docker-compose up -d

down:
	docker-compose down

create-migration:
	migrate create -ext sql -dir migrations $(table_name)
