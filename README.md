## IP 2 GEO Service

### Things missing for the service to be production ready:
1. A real cache service (like redis) that is not in the instance's memory and can be distributed, same for the logger, metrics and exceptions.
2. More unit + integration tests (only the geo controller test exists and not thorough enough).
3. Choosing between the external location services at the moment is only in a predefined order with fallback to the next on failure, but to reduce load a better solution should be implemented.
4. Making sure the graceful shutdown is effective for all running services.
5. Adding docker building and pushing + deployment jobs to the ci/cd.
6. Supporting different configs for staging and production environments.