# Auth test
echo "POST /api/user/register"
curl -i --location --request POST 'localhost:8080/api/user/register' \
--header 'Content-Type: application/json' \
--data '{
    "login":"33",
    "password": "33"
}'

#Orders test
echo "POST /api/user/orders"
curl -i --location --request POST 'localhost:8080/api/user/orders' \
--header 'Content-Type: application/json' \
--cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg4MDc0NzgsIkxvZ2luIjoiMzMifQ.4Ly8BxMolpf7W7_wi-niLWalPo1iq81jtpmZzl3fqf8' \
--data '131416880329'


#Orders test
echo "POST /api/user/orders"
curl -i --location --request POST 'localhost:8080/api/user/orders' \
--header 'Content-Type: application/json' \
--cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg4MDc0NzgsIkxvZ2luIjoiMzMifQ.4Ly8BxMolpf7W7_wi-niLWalPo1iq81jtpmZzl3fqf8' \
--data '326383880704'


# Auth test
echo "POST /api/user/register"
curl -i --location --request POST 'localhost:8080/api/user/register' \
--header 'Content-Type: application/json' \
--data '{
    "login":"44",
    "password": "44"
}'
#Orders test
echo "POST /api/user/orders"
curl -i --location --request POST 'localhost:8080/api/user/orders' \
--header 'Content-Type: application/json' \
--cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg4MDg3OTgsIkxvZ2luIjoiNDQifQ.492NPmAbCYoMMk_QkdausFM44WWH_1uooaUnFZHwykU' \
--data '326383880704'