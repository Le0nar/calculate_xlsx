### For run docker postgresql:

Pull docker image of postgres docker pull postgres
Run container `docker run --name=calculate_xlsx -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d postgres`

### Postresql via docker:

1. docker exec -it {container_id} /bin/bash 
2. psql -U postgres 
3. \d