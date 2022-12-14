build:
	docker build -t wahyubrr/sae:1.0.0-be .

run:
	docker run -d -p 8000:8000 wahyubrr/sae:1.0.0-be

build-jenkins:
	docker build -t myjenkins-blueocean:2.375.1-1 .

run-jenkins:
	docker run --name jenkins-blueocean --restart=on-failure --detach --network jenkins --env DOCKER_HOST=tcp://docker:2376 --env DOCKER_CERT_PATH=/certs/client --env DOCKER_TLS_VERIFY=1 --volume jenkins-data:/var/jenkins_home --volume jenkins-docker-certs:/certs/client:ro --publish 8080:8080 --publish 50000:50000 myjenkins-blueocean:2.375.1-1
