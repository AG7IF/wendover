# Wendover

Wendover is a suite for processing applications to CAP activities.

## Dev Dependencies
### Go Code
- Go 1.22.1 or later
- A local instance of `postgres` is necessary the way things are currently configured (can be set up w/Docker)
- Docker is required for building and running docker images
- Run `go mod tidy`

### Frontend Code
- Node 20.11.1 or later
- Run `npm update` or equivalent command (i.e. `yarn`) in the `web/` directory to install dependencies

## Using the Makefile
Most of the development tasks you will need to do are spelled out in the Makefile. You can use them as-is, or you 
can use the Makefile recipes as a guide for setting up run configurations in your IDE (recommended for running tests).

### Running Tests

### Dev