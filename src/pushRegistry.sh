#!/bin/bash

docker build . -t jonas27/hello:v*
docker push jonas27/hello:v*