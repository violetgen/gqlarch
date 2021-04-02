# GraphQL Gateway + REST API Microservice Architecture

The project contains 2 RESTful microservice Add Microservice and Subtract Microservice and a Graphql Gateway. 

### To Run the project
cd into the addmicro, subtractmicro and gqlgateway run by doing `go run main.go` . The add microserice runs on port 8081, subtract microserive runs on port 8082 and gqlgateway runs on port 8080.
You can then open the browser based graphql client by visiting localhost:8080.

### Sample Queries
1. Sum Query

```
query{
  sum(req: {a: 1,b:2}){
    result
  }
}
```

2. Subtract Query

```
query{
  subtract(req: {a:1,b:2}){
    result
  }
}
```