# rproxy
A simple HTTP2 proxy written in go.

# Abstract
I have to manage many, many web services and some of them are difficult or impossibile to upgrade.
As software evolves, and cryptography more so, we find often difficult to connect new clients to old webservers.

This is often due to old SSL/TLS methods offered by the servers that are not accepted by the client.

The simpliest solution is an https proxy, written in go and statically linked, that can be put in the middle. Being statically linked, it doesn't depend on old openssl/gnutls/whateverssl libraries and can offer modern ciphers (and HTTP2) to the client.

I haven't found something simple, so I wrote one mysql.

# Usage

The software is fully configured via command line flags:
```
./rproxy -help
Usage of ./rproxy:
  -pidfile string
        PID file (default "/var/run/rproxy.pid")
  -port int
        Listen port (default 8080)
  -reverse_url string
        Reverse URL (default "http://127.0.0.1/")
  -tls_cert string
        TLS Certificate file
  -tls_key string
        TLS Certificate key
```

Just give it a certicate file and key, the URL to reverse to and a port to bind.
