runpostgres:
	docker run -it -e POSTGRES_PASSWORD=postgres -p 5432:5432 --name postgresdocker  postgres
stoppostgres:
	docker stop postgresdocker

