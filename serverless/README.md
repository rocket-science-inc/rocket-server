# Serverless

This is serverless version of The Rocker Server.

### Run with Yarn
```
yarn
yarn start
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