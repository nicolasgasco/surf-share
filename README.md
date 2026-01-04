# SurfShare

A small Go application that will help me deepend Backend concepts.

## Execution

Start the application with Docker Compose:

```bash
  docker-compose up
```

⚠️ In order to run the project, you need the `db/password` file containing the password to access the database. It is not commited to the repository.
## Folder structure

- `app`: contains the Go application code. It's available at http://localhost:8080/.
- `db`: contains the database initialization scripts and password file.
- `ui`: contains the React UI. It's available at http://localhost:3000/.

## Useful commands

- Access database container inside psql shell:
  ```bash
    docker exec -it surf-share-db-1 psql -U postgres -d surf-share 
  ```
- Access application container shell:
  ```bash
    docker exec -it surf-share-app-1 sh
  ```
- Access UI container shell:
  ```bash
    docker exec -it surf-share-ui-1 sh
  ```