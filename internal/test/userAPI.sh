# Auth test
echo "POST /api/user/register"
curl -i --location --request POST 'localhost:8080/api/user/register' \
--header 'Content-Type: application/json' \
--data '{
    "login":"11",
    "password": "11"
}'

#Orders test
echo "POST /api/user/orders"
curl -i --location --request POST 'localhost:8080/api/user/orders' \
--header 'Content-Type: application/json' \
--cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg3NTgxNjEsIkxvZ2luIjoiMTEifQ.vUgLE00-4_GCSXoa9qU0BzV-ZuXPnRp14meQDxjBGXQ' \
--data '1'


#Orders test
echo "POST /api/user/orders"
curl -i --location --request POST 'localhost:8080/api/user/orders' \
--header 'Content-Type: application/json' \
--cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg3NTgxNjEsIkxvZ2luIjoiMTEifQ.vUgLE00-4_GCSXoa9qU0BzV-ZuXPnRp14meQDxjBGXQ' \
--data '22'


#Orders test
echo "POST /api/user/orders"
curl -i --location --request POST 'localhost:8080/api/user/orders' \
--header 'Content-Type: application/json' \
--cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg3NTgxNjEsIkxvZ2luIjoiMTEifQ.vUgLE00-4_GCSXoa9qU0BzV-ZuXPnRp14meQDxjBGXQ' \
--data '333'


#Orders test
echo "POST /api/user/orders"
curl -i --location --request POST 'localhost:8080/api/user/orders' \
--header 'Content-Type: application/json' \
--cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg3NTgxNjEsIkxvZ2luIjoiMTEifQ.vUgLE00-4_GCSXoa9qU0BzV-ZuXPnRp14meQDxjBGXQ' \
--data '4444'


#Orders test
echo "GET /api/user/orders"
curl -i --location --request GET 'localhost:8080/api/user/orders' \
--header 'Content-Type: application/json' \
--cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg3NTgxNjEsIkxvZ2luIjoiMTEifQ.vUgLE00-4_GCSXoa9qU0BzV-ZuXPnRp14meQDxjBGXQ' \