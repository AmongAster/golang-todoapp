include .env
export

export PROJECT_ROOT=${shell pwd}

env-up: 
	@docker compose up -d todoapp-postgres
env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		rm -rf ./out/pgdata && \
		echo "Файлы окруженя очищены"; \
	else  \
		echo "Очистка окружения отменена"; \
	fi

env-portf:
	@docker compose up -d port-forwarder

env-portf-cloase:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутсвует необходимый парамитр seq. пример:make migrate-create seq=init" \
		echo "1"; \
	fi; \
	docker compose run --rm todo-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"


migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутсвует необходимый парамитр action. пример:make migrate-action action=up" \
		echo "Все успешно"; \
	fi; \
		docker compose run --rm todo-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"


todoapp-run:
	@export LOGGER_FLOADER=${PROJECT_ROOT}/out/logs && \
	go run cmd/todoapp/main.go 