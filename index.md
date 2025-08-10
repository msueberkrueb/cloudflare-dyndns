# CloudFlare DynDNS

CloudFlare DynDNS is used for dynamic and configurable A record updates for networks without a static ip address.

## Add Helm repository

```bash
helm repo add cloudflare-dyndns https://msueberkrueb.github.io/cloudflare-dyndns/charts
helm repo update
```

## Install chart

Using config from a file:

```bash
helm install --generate-name cloudflare-dyndns/cloudflare-dyndns
```
