# Network Reporter

The networkreporter responds to GET requests to `/network` with the JSON-encoded
set of network values under which it is running.

This is useful for reporting a Pod's external network characteristics to
interested parties.

Optional environment variables are:
  - `CLOUD` - the cloud environment in which the Pod is running.  See
    [CyCoreSystems/netdiscover](github.com/CyCoreSystems/netdiscover) for details about supported cloud
    environments.

