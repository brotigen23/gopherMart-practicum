max=1000
for i in `seq 2 $max`
do    
    echo "GET /api/orders/{number}"
    curl -i --location --request GET 'localhost:8080/api/orders/1789372997' \
    --header 'Content-Type: application/json' \
    --cookie 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzQ3MDQxNjIsIkxvZ2luIjoiMTIzIn0.wYMINaAHrZJ00aGsNkft_ndk9D7FZs2EQSliPXdm0Fo' 
done