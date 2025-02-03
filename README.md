# Golang
1) Run server
- go run cmd/server/main.go
2) Create excel file:
- go run cmd/create_test_file/main.go

# Postgres 

### For run docker postgresql:

Pull docker image of postgres docker pull postgres
Run container `docker run --name=calculate_xlsx -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d postgres`

### Postresql via docker:

1. docker exec -it {container_id} /bin/bash 
2. psql -U postgres 
3. \d

# Redis 

docker run --name redis -p 6379:6379 -d redis

# TODO:
сделать структуру данных хранения такой:

CREATE TABLE portfolios (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    capital DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);