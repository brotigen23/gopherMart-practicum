.PHONY: run
run:
	go run cmd/gophermart/main.go

.PHONY: test
test:
	go test ./... -v -count 1 -cover

.PHONY: testCover 
testCover: coverage.out
	go tool cover -html=coverage.out -o cover.html

coverage.out:
	go test ./... -coverprofile coverage.out

.PHONY: clean
clean:
	rm coverage.out
	rm cover.html
	
.PHONY: mock
mock:
	~/go/bin/mockgen -destination=internal/repository/mocks/mockUserRepository.go -package=mocks github.com/brotigen23/gopherMart/internal/repository UserRepository
	~/go/bin/mockgen -destination=internal/repository/mocks/mockOrderRepository.go -package=mocks github.com/brotigen23/gopherMart/internal/repository OrderRepository
	~/go/bin/mockgen -destination=internal/repository/mocks/mock_user.go -package=mocks github.com/brotigen23/gopherMart/internal/repository UserRepository
