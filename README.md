wx3
===

wx3 is just a simple (but eventually configurable) web server in go.

### Install & Configuration
* `cp wx3-sample.toml /etc/wx3.toml` - All that's configurable, for now, is the port the that's listened on.
* At the moment, the idea is that you would create a file that imports your libs that work with net/http. These libs can supply their own handlers for URIs in their init() functions.
