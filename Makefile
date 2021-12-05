test-backend:
	$(MAKE) -C backend test

test-frontend:
	yarn
	yarn test

test: test-backend test-frontend

start:
	cd backend/src && go run cmd/server/main.go -data ${PWD}/datastore -addr :9595 -static frontend/build

container:
	docker build -t memorytrace .

