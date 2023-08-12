#!/bin/bash
VERSION=v0.0.2
docker build --tag localhost:5001/auth-service:$VERSION --build-context src=../src/ .
docker push localhost:5001/auth-service:$VERSION