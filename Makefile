include .env
export

export PROJECT_ROOT=${shell cd}

env-up: 
	docker compose up -d todoapp-postgres
env-down:
	docker compose down todoapp-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans; \
	if ["$$ans" = "y"]; then \
		docker compose down todoapp-postgres && \
		rm -rf otp/pgdata && \
		echo "Файлы окруженя очищены" \
	else \
		echo "Очистка окружения отменена";
	fi