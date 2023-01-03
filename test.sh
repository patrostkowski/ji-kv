#!/bin/bash

curl -X POST "http://localhost:9991/v1/api?key=LOL&value=XD" &

curl -X GET "http://localhost:9991/v1/api?key=LOL" &

curl -X DELETE "http://localhost:9991/v1/api?key=LOL" &
