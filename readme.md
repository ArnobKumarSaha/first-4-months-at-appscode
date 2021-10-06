This repo are actually the steps that I have followed when trying to learn golang & kubernetes on my first months in AppsCode.

# Installation
## golang installation on linux (https://golang.org/doc/install)
1) sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.1.linux-amd64.tar.gz
2) add /usr/local/go/bin in $HOME/.profile => export PATH=$PATH:/usr/local/go/bin
3) go version

## docker installation on linux (https://docs.docker.com/engine/install/ubuntu/)

## kubectl installation on linux (https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)

1) `curl -LO "[https://dl.k8s.io/release/$](https://dl.k8s.io/release/$)(curl -L -s [https://dl.k8s.io/release/stable.txt](https://dl.k8s.io/release/stable.txt))/bin/linux/amd64/kubectl"`

2) `sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl`

3) check : `kubectl version --client`


## kind installation on linux (https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
`GO111MODULE="on" go get [sigs.k8s.io/kind@v0.11.1](http://sigs.k8s.io/kind@v0.11.1)`

The above command will add kind in $HOME/go/bin



# Codebase structure


## Golang
1) basics (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/basics)
2) web-extras -> without-framework, learn-auth (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/web-extras)
3) chi (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/chi)
4) web-extras -> others..  (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/web-extras)
5) concurrency (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/concurrency)

## Kubernets
6) k8s-doc (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/k8s-doc)
7) k8s-yamls (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/k8s-yamls). 
8) crud-client-go (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/crud-client-go)
9) Makefiles (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/Makefiles)
10) explore-repos (https://github.com/ArnobKumarSaha/golang-and-kubernetes/tree/main/explore-repos)

