buildModel:
	rm -rf src/backend/dbmodels/*
	sqlboiler psql
postgres:
	docker run --name postgresdb -p 5450:5432 -v ./pgdata:/var/lib/postgresql/data -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d --restart always postgres:12-alpine