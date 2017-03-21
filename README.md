# kontainer.ooo

|Master|Develop|
|---|---|
|[![Build Status](https://travis-ci.org/kontainerooo/kontainer.ooo.svg?token=pB7oDfqnVGNEsecqRD8R&branch=master)](https://travis-ci.org/kontainerooo/kontainer.ooo)|[![Build Status](https://travis-ci.org/kontainerooo/kontainer.ooo.svg?token=pB7oDfqnVGNEsecqRD8R&branch=develop)](https://travis-ci.org/kontainerooo/kontainer.ooo)|

## Building kontainer.ooo (*in progress*)
The build consists of mainly two build processes.

1. Building the go backend
1. Building the angular frontend

To set up the environment you need to have `docker`, `docker-compose` (and if you are not on linux `docker-machine`). Builing the backend requires `golang` and building the frontend `node` and `npm`.

1. Run `make all` to transpile the frontend and compile the backend.
1. Run `docker-compose start` to start the environment
1. If you are developing on the frontend run `make fe-watch`
