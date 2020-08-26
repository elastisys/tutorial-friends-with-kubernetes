# Friends with Kubernetes

HackerNews comments on any Kubernetes-related article tend to fall in three categories:

1. People that did not bump into the problems that Kubernetes solves and consider Kubernetes an unnecessary complication. (This is a very fair point!)

2. People that did bump into the problems that Kubernetes solves and found it an invaluable tool.

3. People that did bump into the problems that Kubernetes solves, but didn't manage to effectively use Kubernetes to solve said problems.

This tutorial is for the latter.

## Goals

After completing this tutorial:

* You will understand common pitfalls when porting an application to Kubernetes.
* You will know how to re-engineer an application to run on Kubernetes.

## Non-Goals

* Teaching Docker/Dockerfile/containerization basics: Please head to [Docker Get Started](https://docs.docker.com/get-started/).
* Teaching Kubernetes basics: Please head to [Kubernetes Tutorials](https://kubernetes.io/docs/tutorials/).
* Teaching Helm basics: Please head to [Helm Tutorial](https://helm.sh/docs/intro/).


## Preparations

You will need basic shell scripting tools, such as `bash`, `curl` and `make`. Here is how you install them in Ubuntu:

```
sudo apt-get install bash curl make
```

Furthermore, you will need:

* [docker](https://docs.docker.com/get-docker/)
* [helm](https://helm.sh/docs/intro/install/)
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
* [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)

Start a Minikube cluster:

```
minikube start
```

We will need the Ingress addons:

```
minikube addons enable ingress
```

Check that everything works, as follows:

```
kubectl get -n kube-system deploy ingress-nginx-controller > /dev/null && echo "==> kubectl and minikube work"

eval $(minikube docker-env)
docker ps > /dev/null && echo "==> minikube's docker works"

helm version > /dev/null && echo "==> helm works"
```

You are ready to get started!
