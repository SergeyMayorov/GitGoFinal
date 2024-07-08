COMPOSE_FILE=docker-compose.yml

all: 
	@echo "Build and Start Docker images..."
	docker-compose -f $(COMPOSE_FILE) --project-name="final_web" up -d --build	