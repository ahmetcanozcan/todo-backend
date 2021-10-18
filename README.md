# todo-backend



## Run project
```bash
go run .
```

## Run unit and integration tests
```bash
go test ./...
```

## Deployment

Deployment is automated using GitHub actions. whenever a push comes to the `master` branch, It triggers a GitHub action. that action build project a container and runs all tests. if tests are not failed, Create a docker image and deploy this image into the Heroku container registry, after that Heroku deploys this image as a running container


## Development Process

- Initialize project using `go mod` and create initial folder structre.
- Write unit tests for service and repository
- Write interface, not any business logic implemented untill this step
- Write integration test for controller layer, using mock http request and responses.
- Implement whole project
- Update code in TDD recycle
- Dockerization
- Create Github actions and CI/CD pipeline
