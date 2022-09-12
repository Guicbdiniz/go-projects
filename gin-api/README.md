# HTTP Gin REST API

For this project, I focused on learning as much as I could about the [Gin Web Framework](https://github.com/gin-gonic/gin). 

Go standard library definitely has all the tools necessary for the creation of a complete HTTP API. I believe, though, that Go has many third party libraries that can facilitate and speed up this process. Gin seemed like a good one to start with.

### Project specifications:

- It follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout), as I decided to test it.
- It uses [Docker](https://www.docker.com/) to run the application in a container.

### Installation 

- Build the Docker image with:  
    ```
        docker build . -f build/Dockerfile --tag gin-server
    ```

- Run the image in a Docker container with:  
    ```
        docker run -p 3001:3001 -e GIN_MODE='release' gin-server
    ```