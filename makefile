up:
	docker compose -f docker/compose.yml up
stop:
	docker compose -f docker/compose.yml down
	docker system prune -a -f
