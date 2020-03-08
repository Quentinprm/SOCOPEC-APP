# SOCOPEC-AP
Application web en go , react et postgre
## Lancemement de l'application

```bash
$ docker-compose build
$ docker-compose up
```

1. url du frontend : localhost:3000
2. url du backend : localhost:5000


## Routes d'api (localhost:5000):
- GET 		"/user"
- GET 		"/user/:userId"
- POST		"/user" avec l’objet Usersans PK  en json dans le body
- PUT		"/user/:userId" avec l’objet User en json dans le body
- DELETE	"/user/:userId"

- GET		"/car?statusId&agencyId&model&minYear&maxYear&minPower&maxPower&length&width&height"
- GET		"/car/:carId"
- GET   "/car/:carId/carstate
- GET   "/car/:carId/carstate/:date
- POST		"/car" avec l’objet car sans pk en json dans le body
- PUT		"/car/:carId" avec l’objet car en json dans le body
- DELETE	"/car/:carId"

- GET 		"/agency"
- GET 		"/agency/:agencyId"
- POST 		"/agency" avec dans le body l’objet agency sans pk en json
- PUT 		"/agency/:agencyId" avec dans le body l’objet agency en json
- DELETE	"/agency/:agencyId"

- GET		"/status"
- GET		"/status/:statusId"
- POST		"/status" avec l’objet status sans PK en json dans le body 
- PUT		"/status/:statusId" avec l’objet status en json dans le body
- DELETE	"/status/:statusId"

- GET		"/carstate"
- GET		"/carstate/:carStateId"
- POST		"/carstate" avec l’objet cartate sans la pk en json dans le body
- PUT		"/carstate/:carStateId" avec l’objet cartate en json dans le body
- DELETE	"/carstate/:carStateId"




