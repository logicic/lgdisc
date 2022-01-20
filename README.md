# lgdisc
## design:
#### datastruct:

- registry: map string => application struct
key-value to save the service and host relationship.

- applications: struct
  - appid: application unique id
  - instances: map string => instance struct
  - latestTimestamp: the latest time to update
key to value which is the applications.

- instance: struct
  - env: service enviroment symbol. dev、test ...ect
  - appID: service unique id
  - address: service instance address. http or rpc
  - status: service status stated service online or downline

### api:
use the `gin` to make a web server.

server:
- service registry
```
POST http://$HOST/discory/register
```shell
request params:
- 
client:
- service discory
