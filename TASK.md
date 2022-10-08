The goal of this challenge is to create a backend for a simple movie gallery which can be used by other applications to
display e.g. a library with title information, cover images and alike. However focus should not be put as much on the
application itself but rather on how to run it.

1. Create a very simple service that exposes an API to store and get a movie. The data you collect is up to you.
2. Store the movie data to a database of your choice. Data should be read from the database too.
3. Provide a docker-compose setup of the application and database.
4. Deploy the application including the database to an orchestrated environment of your choice (e.g. any managed
Kubernetes environment, local K8s, ...). Think about different approaches, how you can achieve the goal and how
your work can be reviewed.
5. Document your work

Design

API:

CRUD - 
create, - POST movies
read, GET movies GET movie/id


Movie Data: 

Title 
Cover Image [Image processing and storage] 
Release date 
Genre 
Director
Language
Country of Origin
Actors

Docker compose
Dockerfile - lean as possible with multi-stage caching

Minikube
Docker eval/ Use github actions
Docker build

Deployment
StatefulSet - db
PVC - db

