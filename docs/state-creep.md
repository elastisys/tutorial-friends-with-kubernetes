Imagine the following situation. You are working in a regulated environment, such as [healthcare](https://elastisys.com/hipaa-compliance-kubernetes-privacy-rule/). You want to make sure that your user's passwords are safe, hence you use [Secure Remote Password protocol (SRP)](https://en.wikipedia.org/wiki/Secure_Remote_Password_protocol) for authentication. In brief, SRP does not require the client to send the password to the server, nor the server to store the password. Instead, a sophisticated exchange allows the client to prove to the server that it knows the password. Similarly, the client can verify that the server knows the password.

## Running on Kubernetes

The SRP server of your company was coded years ago. Let us run it on top of Kubernetes:

```
git clone https://github.com/elastisys/tutorial-friends-with-kubernetes
cd tutorial-friends-with-kubernetes/code
```

!!!note
    To simplify this tutorial, the `srp-server` includes a hard-coded database with a single test user.

Thanks to the magic of ready-made tutorials, the Dockerfile and Kubernetes resources are already written.

!!!note
    To simplify this tutorial, you will build the container image directly inside the Docker Daemon of Minikube. Usually, you should push container images to a registry.

```
eval $(minikube docker-env)
docker build -t srp-server srp-server
kubectl apply -f srp-server/deploy

kubectl get pods
```

Awesome! The SRP server seems to work. That was really easy. Let's check if it can actually perform logins:

```
eval $(minikube docker-env)
docker build -t srp-client srp-client
kubectl run \
    -ti \
    --rm \
    --generator=run-pod/v1 \
    --image srp-client \
    --image-pull-policy Never \
    srp-client \
    http://$(minikube ip)
```

You should see a screen full of green checkboxes. Great success!
