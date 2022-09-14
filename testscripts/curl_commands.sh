#!/bin/bash

TOKEN=$(curl -s -X POST http://localhost:8000/auth/sign-in   -H 'Accept: application/json' -H 'Content-Type: application/json' -d '{"username":"UncleBob","password":"cleanArch"}' | jq -r '.token')

curl -X POST http://localhost:8000//api/bookmarks    -H 'Content-Type: application/json' -H "Authorization: Bearer ${TOKEN}"  -d '{"url": "https://github.com/zhashkevych/go-clean-architecture","title": "Go Clean Architecture example"}'
curl -X GET http://localhost:8000/api/bookmarks -H "Authorization: Bearer ${TOKEN}"

curl -X GET http://localhost:8000/api/tasks -H "Authorization: Bearer ${TOKEN}"
curl -X POST http://localhost:8000//api/tasks    -H 'Content-Type: application/json' -H "Authorization: Bearer ${TOKEN}"  -d '{"taskDetail": "myfirstTask","dueDate": "2022-09-14T06:50:08+00:00"}'
curl -X GET http://localhost:8000/api/tasks -H "Authorization: Bearer ${TOKEN}"

