# StatusContentCacheInfo

The status item that reports information about the Content Cache service.

**Platforms:** macOS 27.0 (Beta)

## Properties

### content-cache.info

- **Type:** `StatusContentCacheInfoContentCacheInfoObject`
- **Required:** Yes

A dictionary that contains info about the usage of the Content Cache on the device

## Discussion

### Status item availability

### Status item example

```json
{
    "content-cache": {
        "info": {
            "cache-free": 10737418240,
            "cache-limit": 0,
            "cache-status": "OK",
            "cache-used": 5368709120,
            "max-cache-pressure-last-hour": 0.15,
            "personal-cache-free": 2147483648,
            "personal-cache-limit": 0,
            "personal-cache-used": 1073741824
        }
    }
}
```

## Topics

### Objects

- [StatusContentCacheInfoContentCacheInfoObject](/documentation/devicemanagement/statuscontentcacheinfocontentcacheinfoobject) - A dictionary that contains info about the usage of the Content Cache on the device

