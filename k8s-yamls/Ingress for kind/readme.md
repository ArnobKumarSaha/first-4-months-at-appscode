First create the multinode cluster.

As,  to use ingress ,  we need a ingress controller , make it the controller using manifest file taken from kubernetes's ingress-nginx github repo. (https://github.com/kubernetes/ingress-nginx/tree/main/deploy/static/provider/kind)

wait for everything being created. 2-3 minutes required.

Now, it is ok to deploy pods, svcs & ingress.

-- 
It is requires to add "127.0.0.1  HOST_PATH(in ingress rules)" in /etc/hosts file. to use ingress in kind.



NOTE:  after adding annotations: 
    nginx.ingress.kubernetes.io/rewrite-target: "/" ,
  Every requests that start with arnob.com/foo/ are responding "Hello world". (Hitting the '/' path of my container.)

Giving path: "/foo" or "/foo/" doesn't matter. same.. 

-- Everything is just working fine with single node cluster, as described in kind doc.
