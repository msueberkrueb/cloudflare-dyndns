# CloudFlare DynDNS

CloudFlare DynDNS is used for dynamic and configurable A record updates for networks without a static ip address.

It uses the go sdk and trace api from CloudFlare, as well as a yaml configuration for record and secret management.

# Configuration

The configuration is a yaml document specifying which records should be updates, the corrosponding settings and the api key that should be used.

> [!IMPORTANT]
> Each record must have a default or record specific api_key and zone_id!

## Example Configuration

```yaml
api_key: a1b2
zone_id: c3d4
ttl: 60 # Optional
comment: Managed by CloudFlare DynDNS # Optional

records:
  example.com:
    proxied: true # Defaults to false
  "*.example.com": # Sub-Domain wildcards must be quoted
    proxied: true
  other-example.com:
    api_key: e5f6 # Record configs are preferred
    zone_id: g7h8 # This can be used to manage multiple zones with different keys
```
