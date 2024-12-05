# Backend Basics

 This repository is a collection of notes and code examples for backend development, focusing on Go, Docker, Kubernetes, gRPC, AWS, and CI/CD. This will have a little bit information on postgres database too. For ease of understanding, we will creating a simple banking system, where we can create and manage accounts, record all balance changes, and record transactions, also perform money transfer between accounts.

## Table of Contents

### Containerization

A short note about containers and the commands.

To check for running containers: `docker ps`
To check for all containers: `docker ps -a`
To stop a container: `docker stop <container_id>`
To remove a container: `docker rm <container_id>`

To check for images: `docker images`
To remove an image: `docker rmi <image_id>`

To pull an image from Docker Hub: `docker pull <image_name>:<tag>`

To run a container from an image: `docker run --name <container_name> -e <environment_variable_name>=<environment_variable_value> -p <host_port>:<container_port> -d <image_name>:<tag>`

### Database

#### Database Design

We must design the database schema before we start, this is called design first approach. We can use many tools to design the database schema, but for now we will use [dbdiagram.io](https://dbdiagram.io/). With this tool, we can design the database schema and export it as a SQL file to be imported into our postgres database (can be done for MsSQL and MySQL too). We can also export it as a PNG image to so that the relationship between tables can be visualized and documented.

![Entity Relationship Diagram](https://i.ibb.co/sRqkXqc/simple-bank-vis.png)

We pull a postgres image from Docker Hub and run a container from it. (We will use alpine version of postgres, which is a minimal version of postgres.)

```sh
docker pull postgres:alpine
```

To run a container from the image:

```sh
docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:alpine
```

To check if the container is running, we can use `docker ps` command.

To connect to the postgres container, we can use:

```sh
docker exec -it <container_id> or <container_name> psql -U <username> -d <database_name>
# for example
docker exec -it postgres psql -U root
```

You can quit the psql session by typing `\q`.
To view the logs of a container:

```sh
docker logs <container_id> or <container_name> or <container_id>
# for example
docker logs postgres
```

We can use multiple tools like `data_grip`, `pgadmin`, `dbeaver` to connect to the postgres database engine. I personally use `data_grip` but it is paid.

#### Database Schema Migration using Go

We use [golang-migrate](https://github.com/golang-migrate/migrate) to manage the database schema migrations. It has multiple database drivers, so we can use it to migrate databases like postgres, mysql, mongodb etc.

To install golang-migrate:`brew install golang-migrate`

To create a new migration file: `migrate create -ext sql -dir db/migration -seq <migration_name>`

```sh
migrate create -ext sql -dir db/migration -seq init_schema
```

To migrate or build schema we can use:

```sh
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

To undo a migration:

```sh
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
```

#### CRUD operations

There are multiple ways to perform CRUD operations on the database. We can use raw SQL queries, or use a ORM like `Gorm`, or use a query builder like `sqlc` or `sqlx`.

For this project, we will use `sqlc` to generate the code for the CRUD operations.

To install sqlc: `brew install kyleconroy/sqlc/sqlc`

##### Transaction Summary

1. Create a transfer record with amount 10.
2. Create an entry for account1 with amount -10.
3. Create an entry for account2 with amount 10.
4. Subtract the transfer amount from account1 balance.
5. Add the transfer amount to account2 balance.

Reason to use database transaction:

1. To provide a reliable and consistent unit of work, even in case of system failure.
2. To provide isolation between programs that access the same data concurrently.

Goals to satisfy the ACID properties:

1. Atomicity: Each transaction is all or nothing.
2. Consistency: Transactions ensure that the database is in a consistent state before and after the transaction.
3. Isolation: Transactions are isolated from each other, meaning that the changes made by one transaction are not visible to other transactions until it is committed.
4. Durability: Once a transaction is committed, it remains in the database even if there is a system failure.

We bind all the transactions inside.

```sql
BEGIN;
-- your statements here
COMMIT;
```

If something fails, we can rollback the transaction using `ROLLBACK;` statement.
