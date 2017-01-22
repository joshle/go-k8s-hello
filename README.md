# go-k8s-hello

[![CircleCI](https://circleci.com/gh/alexharveyeu/go-k8s-hello.svg?style=svg)](https://circleci.com/gh/alexharveyeu/go-k8s-hello)

"HELLO WORLD" ~ Golang, Kubernetes, Microservice, GCP, CircleCI

```
$ go run main.go 
2017/01/21 13:07:07 starting, will listen on port 8080
2017/01/21 13:07:18 GET: /
2017/01/21 13:07:33 POST: /
^Csignal: interrupt
```

## Circle CI ~ circle.yml
- PROJECT_NAME = go-k8s-hello
- GCP_CLUSTER_NAME = my k8s cluster name on gcp
- GCP_PROJECT = my project on google cloud platform, example: something-123
- GCP_NAMESPACE = k8s namespace, normally just "default"
- GCP_SERVICE_KEY = make a key with "project.edit" privilege on gcp and base64 encode it
- GCP_ZONE = example: europe-west1-d