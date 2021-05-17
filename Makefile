psql:
	psql -h localhost -p 5432 -W todos -U admin -E

run:
	docker-compose start
