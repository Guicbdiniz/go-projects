# HTTP Gin REST API

For this project, I focused on learning as much as I could about the [Gin Web Framework](https://github.com/gin-gonic/gin). 

Go standard library definitely has all the tools necessary for the creation of a complete HTTP API. I believe, though, that Go has many third party libraries that can facilitate and speed up this process. Gin seemed like a good one to start with.

### Project specifications:

- It follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout), as I decided to test it.
- It uses [Docker](https://www.docker.com/) and [Compose](https://docs.docker.com/compose/) to run the application in containers.

### Usage 

- `Docker` must be installed with the `compose-plugin`. 

- Start the containers with:  
    ```
        docker compose -f ./deployments/compose.yaml up
    ```

- The API will be available on `localhost:3001`. Test it with:
    ```
        curl localhost:3001/ping
    ```

### Next steps

- [ ] Add logging to the API.
- [ ] Create a NoSQL service with a named volume connected to it.
- [ ] Connect the API to the NoSQL service.
- [ ] Create new routes.