# CRUD Kubernetes specs

#### A library for writing apps that CRUD near copies of the same Kubernetes Specs

##### For example to run many copies of the same applications for different clients

See example files

Run with
`GOOS=linux go build example.go && docker run -it --rm -v $(pwd):/code -v ~/.kube:/root/.kube --workdir /code --entrypoint ./example snarlysodboxer/yq-kubectl:1.7.3`

### Notes to turn into docs
A `SpecGroup` is a group of Kubernetes object specifications, such a `Deployment`, `Service`, `PersistentVolume`, etc. A SpecGroup can easily represent an instance of your apps.

