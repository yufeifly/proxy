This is the proxy of containerized IMDB(in memory db, eg.redis) services.
All requests access redis services through proxy. 

## usage:
### specify the cluster
> This migration system works in a cluster. So you must specify a cluster.
A cluster consists of a proxy node and several worker nodes. 

The config file locates at *./proxy/cluster/cluster.json*.

### How to run this program?
```
make run
```