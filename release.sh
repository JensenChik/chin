#!/usr/bin/env bash

cd frontend
npm run build

cd ../server
rm src/main/resources/static/*
cp ../frontend/dist/*  src/main/resources/static

mvn clean package






