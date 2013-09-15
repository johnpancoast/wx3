wx3
===

wx3 is just a simple (but eventually configurable) web server in go.

### Install & Configuration
* `cp extras/conf/wx3.toml /etc/wx3.toml` - All that's configurable, for now, is the port that's listened on.
* At the moment, the idea is that you would create a file that imports your libs that work with net/http. These libs can supply their own handlers for URIs in their init() functions. Eventually, these libs *may* be configured in the toml file.
