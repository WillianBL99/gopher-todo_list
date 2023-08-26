<p align="center">
  <a href="https://github.com/WillianBL99/repo-provas">
    <img src="https://github.com/WillianBL99/gopher-todo_list/assets/65803142/fc32a68b-929e-4849-8f66-d5d875f5456f" width="180" >
  </a>

  <h1 align="center">
    Gopher ToDo List API
  </h3>
</p>
</br>

## :page_facing_up: About

Gopher ToDo List API is a REST API developed to manage a to-do list. The API was developed using Go and PostgreSQL. The API is available on [Render](https://gopher-todolist.onrender.com/).

## :bulb: Motivation

I developed this API with the purpose of putting into practice my studies of the Go language, as well as my knowledge of software architecture. In this specific API, I used the clean architecture to separate the responsibilities of each layer of the application and make it more decoupled. In this API, it's possible to replace, for example, the database without many side effects, since everything is interconnected through interfaces.

## :rocket: Technologies used
The project was developed using the following technologies:

- [<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />](https://golang.org/)
- [<img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />](https://www.postgresql.org/)
- [<img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />](https://www.docker.com/)

## :warning: Prerequisites
>To run the project locally, you must have installed:
- [Go](https://golang.org/) - v1.20
- [PostgreSQL](https://www.postgresql.org/) - (or use [Docker](https://www.docker.com/))

>To run the project into a container, you must have installed:
- [Docker](https://www.docker.com/)

## :cd: Usage
### How to run for development (Locally)

1. Clone this repository and install all dependencies.

    ```bash
    $ git clone https://https://github.com/WillianBL99/gopher-todo-list.git

    $ cd gopher-todo-list

    $ go mod download
    ```

2. Create and configure the `.env` file based on the `.env.example` file.

3. Create a PostgreSQL database with any name you like, or create a container with Docker. The repository contains the `create-tables.sql` file that is in `internal/infra/db/postgresql` for creating the tables. Remember to correctly configure the `.env` file.

    ```bash
    # Create tables with the create-tables.sql file
    $ psql -U postgres -d <database-name> -a -f create-tables.sql
    
    # Or create a container with Docker
    $ docker run --name <container-name> -e POSTGRES_PASSWORD=<password> -p 5432:5432 -d postgres
    ```

4. Run the API

    ```bash
    $ go run ./cmd/todolist
    ```

The API will display `Server running on port <port>`, if everything is correct.

### How to run tests for development (Locally)

1. Run the command below to run the tests.

    ```bash
    $ go test ./...
    ```

### How to run for development (Docker)

There are two ways to run the API using docker. The first is running the script `start.sh` that is in the root of the project. The second is running the docker-compose file.

#### Using the start.sh script
1. Run the command below to run the script.

    ```bash
    # Give permission to the script
    $ chmod +x start.sh
    # Run the script
    $ ./start.sh
    ```

#### Using docker-compose
1. Run the command below to run the docker-compose file.

    ```bash
    $ docker-compose up
    ```
2. Create the tables in the database.

    ```bash
    $ docker exec -it pg_db psql -U postgres -d todo_list -a -f create-tables.sql
    ```

### How to run tests for development (Docker)
1. Run the command below to run the tests.

    ```bash
    $ docker exec -it go_api go test ./...
    ```

### How to run for production (Docker)
For to run the API in production, you must have installed [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/).
1. Update the `.env` file with the production environment variables. Provide the connection variables to the PostgreSQL database.
2. Run the command below to run the docker-compose file.

    ```bash
    $ docker-compose -f docker-compose.prod.yml up
    ```

## :twisted_rightwards_arrows: Available routes in the API
Above are the available routes in the API. For more details, see the documentation available at [gopher-todolist.onrender.com](https://gopher-todolist.onrender.com/).

### API
- `GET /`: See the API documentation.
- `GET /health`: Check the API health.
### Auth
- `POST /auth/sign-up`: Create a new user.
- `POST /auth/sign-in`: Authenticate a user.

### Tasks
- `GET /tasks`: Get all tasks.
- `POST /task`: Create a new task.
- `GET /task/{id}`: Get a task by id.
- `PUT /task/{id}`: Update a task by id.
- `PATCH /task/{id}/done`: Mark a task as done.
- `PATCH /task/{id}/undone`: Mark a task as undone.
- `DELETE /task/{id}`: Delete a task by id.

## :star: Curiosities

The repository has some scripts to automate some processes.
- `start.sh`: [Linux] Script to start the API. It's possible to start the API and execute the cron jobs to generate the backups of the database.
  
  ```bash
    # Start the API
    $ ./start.sh

    # Or start the API and execute cron jobs
    $ ./start.sh --cron
    ```
- `cron.sh`: [Linux] Script to execute the cron jobs to generate the backups of the database.
  
  ```bash
    # Execute cron jobs
    $ ./cron.sh

    # Execute cron jobs and set the time interval
    $ ./cron.sh -t 3 # 3 hours

    # Or execute cron jobs only once
    $ ./cron.sh -o

## :memo: To do
- [x] Add unit tests for the use cases
- [ ] Add unit tests for the repository implementation
- [ ] Add unit tests for the controllers
- [ ] Add unit tests for the handlers
- [ ] Add integration tests
- [ ] Add OAuth2 authentication
- [x] Dockerize the API
- [ ] Add CI/CD
    
## :page_facing_up: License
This project is under the [MIT license](https://github.com/WillianBL99/gopher-todo_list/blob/main/LICENSE)

---
Desenvolvido por **Paulo Uilian Barros Lago**🧑🏻‍💻
