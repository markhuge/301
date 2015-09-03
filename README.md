# 301 

> A super simple HTTP redirection server

This is a Go port of a [project I built with Node.js a while back](https://github.com/markhuge/301-node).

## Usage
```

A super simple HTTP redirection server

-h, --help    print this message
-p, --port    port 301 will listen on (default 8080)
-d, --domain  domain requests will redirect to (default 127.0.0.1)
    --proto   protocol HTTP or HTTPS (default HTTP)
    --health  optional port on which to listen for a health check
              Handy for load balancers that will only accept a "200"
              response to keep 301 instance(s) in load.
```
