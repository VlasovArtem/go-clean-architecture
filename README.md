# GoLang Clean Architecture Example
This is a simple example of how to implement a clean architecture in GoLang.
![Diagram](./documentation/diagram.png)
## Idea
The main idea is to have a project that is easy to maintain and extend. The project is divided into 4 layers:
- **Domain**: This layer contains interfaces and structs that are used in the whole project. It is the core of the project.
- **Service**: This layer contains the business logic of the project. It is the layer that communicates with the controller layer and the repository layer.
- **Repository**: This layer contains the database logic of the project. It is the layer that communicates with the service layer and the database.
- **Controller**: This layer contains the API logic of the project. It is the layer that communicates with the service layer and the API.
### Examples
* [Domain](./internal/domain)
* [Service](./internal/service)
* [Repository](./internal/repository)
* [Controller](./internal/controller)
## Additional information
### Main.go
The application has only one entrypoint: [main.go](./cmd/main.go). The file is responsible to run the application.
### Internal package
The folder [./internal](./internal) contains the application logic.

The folder [./internal/app](./internal/app) contains the application configuration, dependencies manager and startup logic.

The folder [./internal/config](./internal/config) contains the application configuration.

The folder [./internal/domain](./internal/domain) contains the application domain logic.

The folder [./internal/service](./internal/service) contains the application service logic.

The folder [./internal/repository](./internal/repository) contains the application repository logic.

The folder [./internal/controller](./internal/controller) contains the application controller logic.

The folder [../internal/errors/](./internal/errors) contains the application errors.
### Build
The [./build](./build) folder contains the build scripts. For example Dockerfile
### Api Documentation
The [./api](./api) folder contains the api documentation.
### Deployment
The [./deployment](./deployment) folder contains the deployment scripts. For example Kubernetes deployment.
# How to use
Use [a sample](https://github.com/VlasovArtem/go-clean-architecture-sample) as starting point. 