# Lookup Service
This is a simple IP lookup service built in Go. I did this for an interview
process some time ago, and folks asked me to make it private once I got the feedback. 
However, I spoke with their CEO, and he turned out to be a cynical, and pretentious 
little prick. I don't work with pricks, and also, I don't work for free, and my time is valuable.
So I'm making this code available as a small sample of my work as a software developer.
It contains the code for the service itself, local development environment, and a
Helm setup for deploying in any Kubernetes cluster.

## Local environment setup
The setup is done by downloading and installing both Docker, and Docker Compose.
Just bring it by simply typing:

```bash
make localenv-start
```

The API's schema can be found in the `swagger` directory. It really queries for
IP addresses, given a domain name, nothing fancy, really.

## Kubernetes setup
To deploy it in any Kubernetes cluster, you need to build the image:

```bash
make build
```

And then, install the Helm release in the cluster:

```bash
helm install lookup-service ./helm
```

It'll bring up one pod for the service and another for the DB.

```
NAME                                READY   STATUS    RESTARTS   AGE
lookup-service-779f956d58-vhm74     1/1     Running   0          21s
lookup-service-db-d9d4b6588-mbv55   1/1     Running   0          21s
```