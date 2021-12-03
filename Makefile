start:
	cd backend/src && go run cmd/server/main.go -datastore ${PWD}/datastore -addr :9595

container:
	docker build -t memorytrace .

