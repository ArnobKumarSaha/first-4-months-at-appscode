1) As nodeAffinity is used in the new-host-pv.yaml, you have to label the node using " kc label nodes first-control-plane hello=world"
2) Also, create a path using docker exec on the node , "/home/arnob/Desktop/Hello/"
3) Now apply everything.
4) Make a post request in NODE_IP:30012/story with request body = {"text": "abcd"}.
   other available paths are GET /story,  /error
