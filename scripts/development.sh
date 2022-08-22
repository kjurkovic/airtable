#!/bin/bash

# build services
make -C ..

# run containers
docker compose up
