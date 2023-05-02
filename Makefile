start: 
	docker compose down
	docker compose up --build -d
	
stop:
	docker compose down