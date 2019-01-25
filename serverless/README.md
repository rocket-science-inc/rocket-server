Serverless
===============================

This is Functions as a Service (FaaS) serverless container which is used as a temprorary server-side application while the Rocker Server is main server-side logic and still been writting by the application developer unlike main architecture, it’s run in stateless compute container.

![image](https://user-images.githubusercontent.com/1122708/51774779-253b4700-20fc-11e9-89c9-37ccbadbf5f6.png)

### Build with Docker
```
docker-compose build
```
### Run with Docker
```
docker-compose up
```
### Browse
```
http://localhost:4000/graphql
```
### Query example
```
{ 
  # Get user's first name with id equal 200:
  user(id: 200) {
    firstName  
  }
  # Get title and all members of event with id equal 100:
  event(id: 100) {
    title
    members {
      firstName
    }
  }
  # Get all users id and first name:
  users {
    id
    firstName
  }
  # Get all events id and title:
  events {
    id
    title
  }  
}
```
### Mutation example
```
mutation {
  # Create new user and get his id and email:
  createUser(email: "peter@gmail.com") {
    id
    email
  }
  # Create new event and get its id and title:
  createEvent(title: "Data Science Course") {
    id
    title
  }
  # Remove user with id equal 200:
  deleteUser(id: 200)
  # Remove event with id equal 300:
  deleteEvent(id: 300)
}
```

## Development

### Resolve dependencies
```
yarn
```
### Run with Yarn
```
yarn start
```
### Watch with Yarn
```
yarn watch
```
### Debug with VSCode
```
yarn debug

Start Debugging in VScode
```
