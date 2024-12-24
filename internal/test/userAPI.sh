# Auth test
echo "POST /api/user/register"
curl -i --location --request POST 'localhost:8080/api/user/register' \
--header 'Content-Type: application/json' \
--data '{
    "login":"123",
    "password": "asd"
}'

#Orders test
echo "POST /api/user/orders"
curl -i --location --request POST 'localhost:8080/api/user/orders' \
--header 'Content-Type: application/json' \
--cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzUwNTI2ODgsIkxvZ2luIjoiMTIzIn0.UjnXBcXKQEJzMZg8SwPepqivagD80e8LVHHl7aDWXzc' \
--data '1111'
