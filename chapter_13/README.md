# SQL and Databases
- In this section, we will learn how to use the Go programming language to connect to and and wok SQL databases.
- We will connect to databases; create tables, inser data into those tables, retrieve data from those tables.
- In the end, we will get to know how to delete tables, truncate them or drop these tables altogether.

## Introduction
- As a developer (me, lol), it is best to have knowledge/or a proper understanding of persistent data storage and databases.
- The applications we produce collect input, process them and produce outputs. Most of the time, a database is involved in this process.
- The nice things about a database is that, it can be in these forms:
  - an in-memory database OR a database where data is kept in the RAM
  - an file-based database (storing data in a file somewhere)
  - a local or remote storage.

## The Database
- We will be working with `PostgreSQL`
- I am running a `Postgres` container on my machine with Docker; the best way (since the world is moving to heavily containerized workloads anyways).
- I will be making my work more fancy by having a compose file with my database and a tool called `pgAdmin`. This tool is a GUI frontend for my database.
- Here is the compose (I grabbed it off the internet because I was too lazy to write this. Besides, there's a ton of these on the internet to copy from if you already know how to use Docker):
```yaml
version: "3.8"
services:
  db:
    image: postgres:16.3-alpine
    container_name: local_pgdb
    restart: always
    ports:
      - "5432:5432"
    environment:
      POSTGRES_USER: taskmaster
      POSTGRES_PASSWORD: 4@strong-password
    volumes:
      - local_pgdata:/var/lib/postgresql/data
  pgadmin:
    image: dpage/pgadmin4:8
    container_name: pgadmin4_container
    restart: always
    ports:
      - "8888:80"
    environment:
      PGADMIN_DEFAULT_EMAIL: user-name@domain-name.com
      PGADMIN_DEFAULT_PASSWORD: strong-password-123@
    volumes:
      - pgadmin-data:/var/lib/pgadmin

volumes:
  local_pgdata:
  pgadmin-data:
```

## Database API and Drivers
- In order to wok with databases, Go uses an approach that is considered "pure". This means that Go has an API that allows a user to use different drivers to connect to databases
- The API comes from the `database/sql` package in the Go packages.
- The drivers are of two types:
  - natively supported; ie done it such a way that you do not have to perform a lot of configuration to connect the database and the application.
  - third-party drivers; these need additional packages installed to function.
- The idea behind the API and driver approach is to present a Unified interface that can be plugged into an application that the developers can utilize to communicate with different types of databases.
  - All that you (the developer) needs to do is import the API and the necessary database driver and you will be able to talk to the database.
  - There is no need to learn driver-specific implementations as the API provides the abstraction layer that helps to accelerate development.
  - With this approach, you will write scripts that will work even if you swap out the underlying database because the API will work to translate the requests to specific driver.
- When working with databases in Go, we differentiate these types of databases:
  - Relational databases
  - NoSQL databases
  - Search and Analytic databases


## Connecting to Databases
- Apparently, connecting to databases is the easiest ting to do withe regards to Go so far.
- To connect to a database, you need at least four(4) things to be in place. A host to connect to; a database to connect to that is running on a port; a username and password.
- This is because a user needs to have privileges that can allow it to perform specific operations such as query, insert or remove data, create or delete databases, manage users and views.
- For our database, we need to use a pure Postgres driver using a module that we can pull from the internet. Using Go, download the module with this command `go get github.com/lib.pq` or `go get -u github.com/lib/pq`.
- Once the package has been installed, we can start initializing the script:
```go
package main

import (
  "fmt"
  "database/sql"
  _ "github.com/lib.pq" // the _ <package name> is an import statement that tells GO to import the package solely for its side effects.
)
```
- When the script has been initialized, we can connect to the database now:
```go
db, err := sql.Open("postgres", "user=taskmaster password=4@strongpassword host=0.0.0.0 port=5432 dbname=postgres sslmode=disable")
```
  - the API provides an `Open()` function which takes in a variety of arguments as seen in the above script. There is alos a shorthand way of doing this though.
  - the first argument, `postgres` tells the function to use the PostgreSQL driver to make the connection.
  - the second argument, is called the connection string which holds the user, password, host, port, dbname and sslmode arguments. `sslmode` is disabled here for this example BUT for production environments, SSL should be enforced for encrypted traffic to the database server.
- The `Open()` function returns two values, one for the database connection and another for the error. How do we check if the connection was successful or not?
```go
if err != nil {
  panic(err)
} else {
  fmt.Println("The connection to the database was successfully initialized.")
}
```
  - the `panic` is used to show that we are not interested in gracefully managing any error that comes from a worngful connection. We just want to cut the connection.
- For a long running application, we can implement a way to check if the connection to the database is still alive/reachable. This is because we may have intermittent network issues that may cause the connection to be lost and prevent an execution on that database. We can check the reachability with this script:
```go
connectivity := db.Ping()
if connectivity != nil {
  panic(err)
} else {
  fmt.Println("Good to go!")
}
```
  - the panic here is used to indicate that the connection has been lost
- Once the executions in the database are done, we need to terminate the connection in order to remove user sessions and free up resources.
  - in big enterprise environments with users hitting the same database, it is wise to aonly connect to the database when needed and terminate the connection once the work is done.
```go
db.Close()
defer db.Close()
```
  - the difference in the above lines of code is the scope. The first one terminates the connection once the code arrives at that line. The other executes once the function in which the execution is happening goes out of scope. The second is te idiomatic way.


## Creating Tables
- The act of creating tables aims to make logical containers that persistently hold data together.
- The database engines then control who can access what data. There are two approaches to get this done:
  - first is by use of `Access Control Lists, ACLs`. The ACL security logic tells us which userhas which permissions, like CREATE, UPDATE, DELETE etc.
  - the second approach involves inheritances and roles. This is by far more robust and better suited for enterprise applications.
- Postgres uses the second approach.
- The general syntax for creating tables in SQL-compliant databases is like this:
```SQL
CREATE TABLE table_name {
  column1 datatype constrain,
  column2 datatype constrain,
  column3 datatype constrain,
  ...
};
```
  - This is `SQL` ie `Structured Query Language` that is a standard used in SQL-complaint databases aka relational databases.
  - The common datatypes that we can use are:
    - INT
    - DOUBLE
    - FLOAT
    - VARCHAR (a string with a specific lenght)
  - The constraints, in a way, imbue our columns with special powers. Constrains also depend on the database engine, some of them are as follows:
    - NOT NULL
    - PRIMARY KEY
    - Named function (a named function is executed every time a new record is inserted or an old one is updated and, based on the evaluation of the transaction, is either allowed or denied).
- We can not only create tables, we can also empty a table, or remove a table.
  - to empty a table, we do:
  ```SQL
  TRUNCATE TABLE table_name
  ```
  - to remove a table from the database, we do:
  ```SQL
  DROP TABLE table_name
  ```
- the script builds on `db_init.go` to create a table using a statement.
```go
package main

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

func main() {
  db, err := sql.Open("postgres", "user=taskmaster password=4@strongpassword host=0.0.0.0 port=5432 dbname=postgres sslmode=disable")
  if err != nil {
    panic(err)
  } else {
    fmt.Println("The connection to the database has been successfully initialized.")
  }

  connectivity := db.Ping()
  if connectivity != nil {
    panic(err)
  } else {
    fmt.Println("Good to go!")
  }

  DBCreate := `
  CREATE TABLE public.test
  (
    id integer,
    name character varying COLLATE pg_ccatalog."default"
  )
  WITH (
    OIDS = FALSE
  )
  TABLESPACE pg_default;
  ALTER TABLE public.test
    OWNER to postgres;
  `

  _, err = db.Exec(DBCreate)
  if err != nil {
    panic(err)
  } else {
    fmt.Println("The table was successfully created!")
  }

  defer db.Close()
}
```
  - if we run this script for the first time, we get the messages indicating that the connection to the DB is successful and the table being created.
  - If we rerun the script, we get the error that the table already exists.


## Inserting Data
- In order to guard against SQL injection attacks, Go utilizes the `Prepare()` statement which provides protection against these attacks.
- Go has two(2) types os substitutions; we either use `WHERE col = $1` in the case of queries, or `VALUES($1,$2)` in the case of updates or inserts.
- Let us update the tables:
```go
// normal init statements
  insert, err := db.Prepare("INSERT INTO test VALUES ($1, $2)")
  if err != nil {
    panic(err)
  }

  _, err := insert.Exec(2, "second")
  if err != nil {
    panic(err)
  } else {
    fmt.Println("The value was sucessfully inserted!")
  }

  db.Close()
}
```
- Essentially, the `db.Prepare()` statement takes an SQL statement and imbues it with protection against SQL injection attacks.


## Retrieving Data
- SQL injection does not only concern data that is being inserted. It also concerns any data that is being manipulated in the database.
- Retrieving data and doing it carefully is something we must prioritize and handle with proper caution.
- When we query data, our results depend on these things:
  - the database we connect to
  - the table we want to retrieve the data from
  - security privileges, by the database engine, for authentication and authorization.
- A successful query depends on whether the user has the appropriate permissions.
- Queries to the database take two forms:
  - those that do not take an argument eg  `SELECT * FROM table`
  - those that require the user to specify the filtering criteria
- Go provides two functions that we can use to query data. One is the `Query()` function and the other is the `QueryRow()` function.
  - the availability of these functions depends on the database engine we are interacting with.
  - most of the time, the `Query()` function will work. And this can be wrapped with the `Prepare()` function against malicious SQL injections.


## Updating Existing Tables
- In Go, there is no `Update()` function that we can use to update a row in a table in the database.
- We only have the `Exec()` function, which serves as a universal executor for our queries. With this, we can `UPDATE`, `SELECT`, `DELETE` or whatever SQL manipulation you want to run.
- In this session, we will look at a safe way through which we can accomplish this.
- Check `database/db_update` for the example.
- A nice additional resource to use is this: [updating and deleting tables](https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/)


## Deleting Data
- We delete data because of many reasons; maybe we do not need it anymore, we are migrating to another database solution or maybe we are changing/replacing the current value in the table.
- Just like with updating, Go does not provide and explicit `Delete()` but we can formulate one based off using SQL.
- The delete statement is in the `database/db_delete` script.


## Truncating and Deleting Tables
- A more elegant way to empty a table is to use the `TRUNCATE TABLE <table_name>` SQL statement with the `Exec()` function.
```go
emptyTable, err := db.Exec("TRUNCATE TABLE <table_name>")
if err != nil {
  panic(err)
}
```
- To get rid of the table competely, we use the `DROP TABLE <table_name>` SQL statement with the `Exec()` function
```go
dropTable, err := db.Exec("DROP TABLE <table_name>")
if err != nil {
  panic(err)
}
```