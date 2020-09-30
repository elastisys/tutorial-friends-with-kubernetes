all: srp-server-deploy

deploy-srp-server: build-srp-server
	kubectl apply -f srp-server/deploy

build-srp-server:
	docker build -t srp-server srp-server

deploy-srp-server-redis: build-srp-server-redis
	kubectl apply -f srp-server-redis/deploy

build-srp-server-redis:
	docker build -t srp-server-redis srp-server-redis

deploy-srp-server-redis-graceful: build-srp-server-redis-graceful
	kubectl apply -f srp-server-redis-graceful/deploy

build-srp-server-redis-graceful:
	docker build -t srp-server-redis-graceful srp-server-redis-graceful

load:
	docker build -t srp-client srp-client
	kubectl run \
			-ti \
			--rm \
			--generator=run-pod/v1 \
			--image srp-client \
			--image-pull-policy Never \
			srp-client \
			http://$$(minikube ip)

scale-to-3:
	kubectl scale --replicas=3 deployment/srp-server

scale-to-2:
	kubectl scale --replicas=2 deployment/srp-server

scale-to-1:
	kubectl scale --replicas=1 deployment/srp-server

redeploy:
	kubectl patch deployment srp-server -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"$$(date +'%s')\"}}}}}"
