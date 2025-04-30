.SIlENT:

deploy:
	curl -X POST http://localhost:8088/featureflag --header "Content-Type=application/json" --data '{"key": "new-retriever", "enabled" : false }'
	go run cmd/main.go


baseline-test:
	echo "Running baseline test where we know that we are around 2 seconds..."
	curl -X PUT http://localhost:8088/featureflag/new-retriever/false
	k6 run --iterations 5 tests/performance-test.js

experiment:
	echo "Running experiment.  New retriever query performed at <100ms in pre-prod."
	curl -X PUT http://localhost:8088/featureflag/new-retriever/true
	k6 run --iterations 5 tests/performance-test.js