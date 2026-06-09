# StatusContentCacheService

The status item that reports the status of the Content Cache service.

**Platforms:** macOS 27.0 (Beta)

## Properties

### content-cache.status

- **Type:** `StatusContentCacheServiceContentCacheStatusObject`
- **Required:** Yes

The basic set of AssetCache status items

## Discussion

### Status item availability

### Status item example

```json
{
    "content-cache": {
        "status": {
            "server-guid": "C3D4E5F6-A7B8-9012-CDEF-012345678912",
            "activated": true,
            "active": true,
            "cache-status": "OK",
            "private-addresses": [
                "192.168.1.5"
            ],
            "public-address": "203.0.113.5",
            "port": 51194,
            "registration-status": 1,
            "startup-status": "OK",
            "tetherator-status": 0,
            "sending-reports": false,
            "version": "2.0"
        }
    }
}
```

## Topics

### Objects

- [StatusContentCacheServiceContentCacheStatusObject](/documentation/devicemanagement/statuscontentcacheservicecontentcachestatusobject) - The basic set of AssetCache status items

