# semi-task

Movie gallery API

#### Requirements

    - golang 1.19.1
    - mongo 6.0.2
    - helm 3.10.0
    - terraform 1.3.2
    - docker compose 2.11.1
    - minikube v1.27.0

### API endpoints

| Endpoint |            |     | body                                                                                                                                                                                      | response                                                                                                                                                                                                                                                                                                                                                                |     |
| -------- | ---------- | --- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- |
| GET      | /movies    |     |                                                                                                                                                                                           | `json [ { "id" : "634371b1a044e41b215f4262" , "title" : "Life" , "genre" : "SCI-FI" , "director" : "director guy" , "language" : "English" , "country_of_origin" : "USA" , "actors" : [ "Random actor-1" , "Random actor-2" ], "release_date" : "2012-1-02" }, { "id" : "634371c1a044e41b215f4263" , "title" : "Dead Dont Die" , "genre" : "Horror Comedy" , ".." } ] ` |     |
| GET      | /movie/:id |     |                                                                                                                                                                                           | `json { "id" : "634371b1a044e41b215f4262" , "title" : "Life" , "genre" : "Horror Comedy" , "director" : "director guy" , "language" : "English" , "country_of_origin" : "USA" , "actors" : [ "Random actor-1" , "Random actor-2" ], "release_date" : "2012-1-02" } `                                                                                                    |     |
| POST     | /movies    |     | `json { "title" : "string" , "genre" : "string" , "director" : "string" , "language" : "string" , "country_of_origin" : "string" , "actors" : [ "string" ], "release_date" : "string" } ` | `json { "id": "000000000000000000000000", "title": "Dead Dont Die", "genre": "Horror Comedy", "director": "director guy", "language": "English", "country_of_origin": "USA", "actors": ["Random actor-1", "Random actor-2"], "release_date": "2012-1-02" } `                                                                                                            |     |

#### Docker Compose

```
  cp .env.example .env
  docker compose up
```

 - API should be available at `http://localhost:$PORT`, `$PORT` is set `.env` file; default `3000`

#### Minikube (k8s)
- minikube `semi-task` profile is required

```
minikube config set memory 4092
minikube start -p semi-task --nodes 2
minikube addons enable ingress -p semi-task
echo $(minikube ip -p semi-task) semi-task.local | sudo tee -a /etc/hosts 
cat /etc/hosts
```

#### Deploy to minikube `semi-task`

- kube config path `~/.kube/config` is required 
- kube context is required to be `semi-task`

```
cd terraform
terraform init
terraform apply -auto-approve
```

- API available at semi-task.local being relayed by the ingress.

## Notes
- For docker compose .env file supplies the password used to intiate the root and app databases
- For minikube-terraform the random provider is used to create the passwords, then handled by the helm provider
- Helm charts are at /charts one for the application and for the database

## Potential Improvements
- Support replication for mongodb.
- Retry logic on app db connect and ping
- Terraform hard-coded contexts and config paths