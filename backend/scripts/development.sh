#!/bin/bash

# build services
make -C ../auth
make -C ../workspace
make -C ../notifications

# run containers
docker compose up
