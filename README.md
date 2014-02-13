barycenter
==========

Barycenter serves a JSON configuration file over HTTP using basic authentication (so run it over SSL).

Run an endpoint as follows:

```barycenter -c config.json -a DEFAULT -p 8080```

You can then make a request against the endpoint.

```curl -u DEFAULT: 127.0.0.1:8080```
