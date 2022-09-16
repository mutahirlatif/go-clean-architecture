#!/bin/bash

#curl -X POST http://localhost:8000/auth/sign-up    -H 'Content-Type: application/json'   -d '{"username":"UncleBob","password":"cleanArch"}'
TOKEN=$(curl -s -X POST http://localhost:8000/auth/sign-in   -H 'Accept: application/json' -H 'Content-Type: application/json' -d '{"username":"UncleBob","password":"cleanArch"}' | jq -r '.token')
echo ${TOKEN}
#curl -X POST http://localhost:8000//api/bookmarks    -H 'Content-Type: application/json' -H "Authorization: Bearer ${TOKEN}"  -d '{"url": "https://github.com/zhashkevych/go-clean-architecture","title": "Go Clean Architecture example"}'
#curl -X GET http://localhost:8000/api/bookmarks -H "Authorization: Bearer ${TOKEN}" | jq

curl -X GET http://localhost:8000/api/tasks -H "Authorization: Bearer ${TOKEN}" | jq
curl -X POST http://localhost:8000//api/tasks    -H 'Content-Type: application/json' -H "Authorization: Bearer ${TOKEN}"  -d '{"taskDetail": "myfirstTask","dueDate": "2022-09-14T06:50:08+00:00"}'
curl -X GET http://localhost:8000/api/tasks -H "Authorization: Bearer ${TOKEN}" | jq

echo "Starting update"
ID=$(curl -s -X GET http://localhost:8000/api/tasks -H "Authorization: Bearer ${TOKEN}" | jq '.tasks[0].id')
#ID=632179ff7e43ed8cf3136fd6
echo "Doing put with ID ${ID}"
curl -X PUT http://localhost:8000//api/tasks    -H 'Content-Type: application/json' -H "Authorization: Bearer ${TOKEN}"  -d '{"id": '"${ID}"',"taskDetail": "myfirstTaskUpdate","dueDate": "2022-10-14T06:50:08+00:00"}'
# -d "{\"id\": \"${ID}\", \"taskDetail\": \"updated Task\",\"dueDate\": \"2022-10-14T06:50:08+00:00\"}"
curl -X GET http://localhost:8000/api/tasks -H "Authorization: Bearer ${TOKEN}" | jq

echo "Doing delete ${ID}"
curl -X DELETE http://localhost:8000//api/tasks    -H 'Content-Type: application/json' -H "Authorization: Bearer ${TOKEN}"  -d '{"id": '"${ID}"'}'
# -d "{\"id\": \"${ID}\", \"taskDetail\": \"updated Task\",\"dueDate\": \"2022-10-14T06:50:08+00:00\"}"
curl -X GET http://localhost:8000/api/tasks -H "Authorization: Bearer ${TOKEN}" | jq
