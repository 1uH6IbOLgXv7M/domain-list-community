# Domain list community

This project manages a list of domains, to be used as geosites for routing purpose in Project V.

## Purpose of this project

This project is not opinionated. In other words, it does NOT endorse, claim or imply that a domain should be blocked or proxied. It can be used to generate routing rules on demand.

## Download links

- **dlc.dat**：[https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat](https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat)
- **dlc.dat.sha256sum**：[https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat.sha256sum](https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat.sha256sum)
- **dlc.dat_plain.yml**：[https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat_plain.yml](https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat_plain.yml)
- **dlc.dat_plain.yml.sha256sum**：[https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat_plain.yml.sha256sum](https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat_plain.yml.sha256sum)

## Notice

- Rules with `@!cn` attribute has been cast out from cn lists. `geosite:geolocation-cn@!cn` is no longer available. Check [#390](https://github.com/v2fly/domain-list-community/issues/390), [#3119](https://github.com/v2fly/domain-list-community/pull/3119) and [#3198](https://github.com/v2fly/domain-list-community/pull/3198) for more information.
- Dedicated non-category ad lists like `geosite:xxx-ads` has been removed. Use `geosite:xxx@ads` instead. `geosite:category-ads[-xx]` is not affected.

Please report if you have any problems or questions.

## Usage example

Each file in the `data` directory can be used as a rule in this format: `geosite:filename`.

> **Personal note:** I use this config as my baseline. The outbound tags (Reject, Direct, Proxy-1, etc.) need to match whatever you have defined in your outbounds section.

```json
"routing": {
  "domainStrategy": "IPIfNonMatch",
  "rules": [
    {
      "type": "field",
      "outboundTag": "Reject",
      "domain": [
        "geosite:category-ads-all",
        "geosite:category-porn"
      ]
    },
    {
      "type": "field",
      "outboundTag": "Direct",
      "domain": [
        "domain:icloud.com",
        "domain:icloud-content.com",
        "domain:cdn-apple.com",
        "geosite:cn",
        "geosite:private"
      ]
    },
    {
      "type": "field",
      "outboundTag": "Proxy-1",
      "domain": [
        "geosite:category-anticensorship",
        "geosite:category-media",
        "geosite:category-vpnservices"
      ]
    },
    {
      "type": "field",
      "outboundTag": "Proxy-2",
      "domain": [
        "geosite:category-dev"
      ]
    },
    {
      "type": "field",
      "outboundTag": "Proxy-3",
      "domain": [
        "geosite:geolocation-!cn"
      ]
    }
  ]
}
```

## Generate `dlc.dat` manually

- Install `golang` (version 1.21 or later recommended) and `git`
- Clone project code: `git clone https://github.com/v2fly/domain-list-community.git`
- Navigate to proje
