# route-policy

Simple example how to retrieve route policies on devices
via GRPC. This could then be leveraged for a route-policy store etc.



Running Against device:

```bash
(base) ➜  metagrpc git:(main) ✗ go run examples/main.go

rpl store via GRPC cfg from 192.168.88.3:57344
 {
 "Cisco-IOS-XR-policy-repository-cfg:routing-policy": {
  "route-policies": {
   "route-policy": [
    {
     "route-policy-name": "TELIA",
     "rpl-route-policy": "route-policy TELIA\n  if community is-empty then\n    set community 9999:999\n  else\n    set community 46489:1299\n  endif\n  if origin is incomplete then\n    drop\n  else\n    prepend as-path 19108\n  endif\nend-policy\n"
    },
    {
     "route-policy-name": "LEVEL-3",
     "rpl-route-policy": "route-policy LEVEL-3\n  if community is-empty then\n    set community 9999:999\n  else\n    set community 46489:3356\n  endif\n  if origin is incomplete then\n    drop\n  else\n    prepend as-path 19108\n  endif\nend-policy\n"
    }
   ]
  }
 }
}

TELIA
-------------
route-policy TELIA
  if community is-empty then
    set community 9999:999
  else
    set community 46489:1299
  endif
  if origin is incomplete then
    drop
  else
    prepend as-path 19108
  endif
end-policy

LEVEL-3
-------------
route-policy LEVEL-3
  if community is-empty then
    set community 9999:999
  else
    set community 46489:3356
  endif
  if origin is incomplete then
    drop
  else
    prepend as-path 19108
  endif
end-policy



```