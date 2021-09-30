# mastering-kubernetes


#install kubectl on linux (A Command line tool to control kubernetes)
1) curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
2) sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
3) kubectl version --client

#install go
0) Download golang .deb file from official website.
1) rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.1.linux-amd64.tar.gz
2) export PATH=$PATH:/usr/local/go/bin  --> add this to ~/.bashrc
3) go version


#install kind
1) GO111MODULE="on" go get sigs.k8s.io/kind@v0.11.1
