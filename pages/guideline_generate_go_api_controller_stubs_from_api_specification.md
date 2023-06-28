# Generate Go API Controller Stubs from OpenAPI API Specification

The C&M microservice engineering approach defines, that artifacts are built incrementally upon each other.
The API specification serves as the foundation for developing the Go API controller code artifacts.

It is possible to generate Go structs representing OpenAPI component schemas and Go interface methods representing OpenAPI endpoints from the API specification.
For this, the [oapi-codegen](https://github.com/deepmap/oapi-codegen) plugin can be used.
The automatic generation of API relevant Go structs and interfaces helps to implement the API controllers a lot.

This guideline describes the steps required to use the plugin for generating API relevant Go structs and interfaces:

1. Ensure Go is properly installed via the command
   ```
   go version
   ```
2. Install the oapi-codegen plugin via the command
   ```
   go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
   ```
3. Open the microservice repository
4. Create directory `scr/api/stubs` and navigate to the directory in the terminal
5. Generate structs representing OpenAPI component schemas via the command
   ```
   oapi-codegen -package stubs -generate types <relative-path-to-api-specification> > server-model.go
   ```
6. Generate interface methods representing OpenAPI endpoints via the command
   ```
   oapi-codegen -package stubs -generate server <relative-path-to-api-specification> > server.go
   ```
7. Now, the directory `scr/api/stubs` contains the two files `server-models.go` and `server.go` with the generated API relevant Go structs and interfaces.
8. Open each generated file and sync dependencies
