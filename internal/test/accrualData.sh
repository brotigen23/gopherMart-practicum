# Auth test
echo "Goods"
curl -v --location --request POST 'localhost:9090/api/goods' \
--header 'Content-Type: application/json' \
--data '{
    "match":"Bork",
    "reward":10,
    "reward_type":"%"

}'

echo "Orders"
curl -i --location --request POST 'localhost:9090/api/orders' \
--header 'Content-Type: application/json' \
--data '{
    "order":"131416880329",
    "goods":[
        {
            "description":"Чайник Bork",
            "price": 111
        }
    ]
}'
curl -i --location --request POST 'localhost:9090/api/orders' \
--header 'Content-Type: application/json' \
--data '{
    "order":"326383880704",
    "goods":[
        {
            "description":"Чайник Bork",
            "price": 555
        }
    ]
}'
