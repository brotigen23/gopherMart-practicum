# Auth test
echo "Goods"
curl -i -X POST 'localhost:9090/api/goods' \
--header 'Content-Type: application/json' \
--data '{
    "match":"Bork",
    "reward":10,
    "reward_type":"%"

}'

echo
echo
echo
echo "Orders"
curl -i -X POST 'localhost:9090/api/orders' \
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

echo
echo
echo
curl -i -X POST 'localhost:9090/api/orders' \
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

echo
echo
echo
curl -i -X POST 'localhost:9090/api/orders' \
--header 'Content-Type: application/json' \
--data '{
    "order":"87051263033",
    "goods":[
        {
            "description":"Чайник Bork",
            "price": 555
        }
    ]
}'