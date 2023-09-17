# Load Balancer

A load balancer sits in front of a group of servers and routes client requests across all of the servers that are 
capable of fulfilling those requests. The intention is to minimise response time and maximise utilisation whilst 
ensuring that no server is overloaded. If a server goes offline the load balance redirects the traffic to the remaining 
servers and when a new server is added it automatically starts sending requests to it.

## Challenge

- Distributes client requests/network load efficiently across multiple servers