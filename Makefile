# Binaries
GO=go

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  dev             Run the development environment"
	@echo "  dev-%           Run a specific command in the development environment"
	@echo "  prod            Run the production environment"
	@echo "  prod-%          Run a specific command in the production environment"

dev:
	docker compose \
		-f compose.yml \
		-f compose.dev.yml \
		--profile dev \
		-p ember \
		up -d --build

dev-down:
	docker compose -p ember down

dev-%:
	docker compose \
		-f compose.yml \
		-f compose.dev.yml \
		--profile dev \
		$* $(ARGS)

prod:
	docker compose \
		-f compose.yml \
		--profile prod \
		-p ember-prod \
		up -d --build

prod-down:
	docker compose -p ember-prod down

prod-%:
	docker compose \
		-f compose.yml \
		--profile prod \
		$* $(ARGS)