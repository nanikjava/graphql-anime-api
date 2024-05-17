# anime-api
golang anime api with graphql

# Setup

## Postgres

Create Docker volume

```
docker volume create animeapi
```

Run postgres on port `15306`
```
docker run  -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=user -p 15306:5432 -v animeapi:/var/lib/postgresql/data postgres
```

Create `anime` DB in postgres  

```
create database anime  with owner "user";
```

Run DB migration, it will run all the DB script inside `db/migrations`

```
go run ./cmd/main.go  migrate up
```

Add dummy data into `anime` table in the DB

## GraphQL

Run `make gql` to generate the GraphQL code. All generated code will be inside `graph` directory.

Run the following command

```
go run ./cmd/main.go serve
```

GraphQL UI can be accessed using `http://localhost:3000/ui/playground`

Execute the following GraphQL

```
query t{
  anime(id:1){
    tags
    titleEn
  }
}
```