## IP 2 GEO Service

### How to:
1. `make run` - To start the project locally
2. `make build` - To build the project locally
3. `make test` - To run the unit tests
4. `make docker-build` - To build the docker image
5. `make docker-run` - To run the docker image
6. `make watch` - To run the project in watch mode

### Things missing for the service to be production ready:
1. A real cache service (like redis) that is not in the instance's memory and can be distributed, same for the logger, metrics and exceptions.
2. More unit + integration tests (only the geo controller test exists and not thorough enough).
3. Choosing between the external location services at the moment is only in a predefined order with fallback to the next on failure, but to reduce load a better solution should be implemented.
4. Making sure the graceful shutdown is effective for all running services.
5. Adding docker building and pushing + deployment jobs to the ci/cd.
6. Supporting different configs for staging and production environments.