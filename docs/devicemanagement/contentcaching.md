# ContentCaching

The payload that configures the Content Caching service.

**Platforms:** macOS 10.13.4

## Properties

### AllowCacheDelete

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If true, the system purges content from the cache automatically when it needs disk space for other apps when free disk space runs low on the computer. Set to `false` to maximize effectiveness of Content Caching. Available in macOS 10.15 and later.

### AllowPersonalCaching

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system caches the user’s iCloud data. Changes to this value don’t have an immediate effect. Clients may take some time, such as hours or days, to react to changes.

> 

### AllowSharedCaching

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system caches non-iCloud content, such as apps and software updates. Changes to this value don’t have an immediate effect. Clients may take some time, such as hours or days, to react to changes.

> 

### AutoActivation

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system automatically activates the content cache when possible and prevents disabling it. If `allowContentCaching` is `false`, `AutoActivation` is also `false`.

Removing a profile that set `AutoActivation` to `true` doesn’t deactivate the Content Cache.

### AutoEnableTetheredCaching

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system automatically enables Internet connection sharing when possible and prevent disabling Internet connection sharing. `DenyTetheredCaching` overrides `AutoEnableTetheredCaching`. Tethered caching requires Content Caching.

Available in macOS 10.15.4 and later.

### CacheLimit

- **Type:** `integer`
- **Required:** No
- **Default:** `0`

The maximum number of bytes of disk space to use for the content cache. Set to `0` for unlimited disk space.

### DataPath

- **Type:** `string`
- **Required:** No
- **Default:** `/Library/Application Support/Apple/AssetCache/Data`

The path to the directory used to store cached content. Changing this setting manually doesn’t automatically move cached content from the old location to the new one. To move content automatically, use the Sharing preference’s Content Caching pane. The value must be (or end with) `/Library/Application Support/Apple/AssetCache/Data`.

The system creates a directory and its intermediates for the given data path if it doesn’t already exist. The directory is owned by `_assetcache:_assetcache` and has mode 0750. Its immediate parent directory (`.../Library/Application Support/Apple/AssetCache`) is owned by `_assetcache:_assetcache` and has mode `0755`.

### DenyTetheredCaching

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables tethered caching.

### DisplayAlerts

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, Content Caching displays exceptional conditions (alerts) as system notifications in the upper corner of the screen. Alerts were automatically displayed starting in macOS 10.13. In macOS 10.15 the alerts are off by default, but still available through this setting. Available in macOS 10.15 and later.

### KeepAwake

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents the computer from sleeping as long as Content Caching is on (System Preferences > Sharing > Content Caching is on). Customers who want Content Caching to be as available as much as possible should turn this setting on. Available in macOS 10.15 and later.

### ListenRanges

- **Type:** `[ContentCaching.Ranges]`
- **Required:** No

An array of dictionaries that describe a range of client IP addresses to serve.

### ListenRangesOnly

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the content cache provides content to the clients in the `ListenRanges`.

### ListenWithPeersAndParents

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the content cache provides content to the clients in the union of the `ListenRanges`, `PeerListenRanges` and `Parents`.

### LocalSubnetsOnly

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the content cache offers content to clients only on the same immediate local network only. No content is offered to clients on other networks reachable by the content cache. If `LocalSubnetsOnly` is `true`, the system ignores `ListenRanges`.

### LogClientIdentity

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the Content Cache logs the IP address and port number of the clients that request content.

### Parents

- **Type:** `[string]`
- **Required:** No

An array of the local IP addresses of other content caches that this cache should download from or upload to, instead of downloading from or uploading to Apple directly. The system ignores invalid addresses and addresses of computers that aren’t content caches. The system skips Parent caches that become unavailable. If all parent content caches become unavailable, the content cache downloads from or uploads to Apple directly, until a parent content cache becomes available again.

### ParentSelectionPolicy

- **Type:** `string`
- **Required:** No
- **Default:** `round-robin`
- **Allowed Values:** `first-available`, `url-path-hash`, `random`, `round-robin`, `sticky-available`

The policy to implement when choosing among more than one configured parent content cache. With every policy, the system skips parent caches that are temporarily unavailable. Allowed values:

### PeerFilterRanges

- **Type:** `[ContentCaching.Ranges]`
- **Required:** No

An array of dictionaries describing a range of peer IP addresses that the content cache uses to filter its list of peers to query for content. The content cache only queries peers in `PeerFilterRanges`. When `PeerFilterRanges` is an empty array, the content cache doesn’t query any peers.

### PeerListenRanges

- **Type:** `[ContentCaching.Ranges]`
- **Required:** No

An array of dictionaries describing a range of peer IP addresses the content cache responds to. When `PeerListenRanges` is an empty array, the content cache responds with an error to all cache queries.

### PeerLocalSubnetsOnly

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the content cache only peers with other content caches on the same immediate local network, rather than with content caches that use the same public IP address as the device. When `PeerLocalSubnetsOnly` is `true`, it overrides the configuration of `PeerFilterRanges` and `PeerListenRanges`. If the network changes, the local network peering restrictions update appropriately. If `false`, the content cache defers to `PeerFilterRanges` and `PeerListenRanges` for configuring the peering restrictions.

### Port

- **Type:** `integer`
- **Required:** No
- **Default:** `0`

The TCP port number on which the content cache accepts requests for uploads or downloads. Set to `0` to pick a random, available port.

### PublicRanges

- **Type:** `[ContentCaching.Ranges]`
- **Required:** No

An array of dictionaries describing a range of public IP addresses that the cloud servers should use for matching clients to content caches.

## Discussion

Specify `com.apple.AssetCache.managed` as the payload type.

### Profile availability

### Profile example

```plist
<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>PayloadContent</key>
    <array>
        <dict>
            <key>AllowPersonalCaching</key>
            <false/>
            <key>AllowSharedCaching</key>
            <true/>
            <key>AutoActivation</key>
            <true/>
            <key>CacheLimit</key>
            <integer>180000000000</integer>
            <key>DenyTetheredCaching</key>
            <false/>
            <key>ListenRangesOnly</key>
            <false/>
            <key>LocalSubnetsOnly</key>
            <true/>
            <key>LogClientIdentity</key>
            <false/>
            <key>ParentSelectionPolicy</key>
            <string>round-robin</string>
            <key>PeerLocalSubnetsOnly</key>
            <true/>
            <key>Port</key>
            <integer>0</integer>
            <key>PayloadIdentifier</key>
            <string>com.example.mycontentcachingpayload</string>
            <key>PayloadType</key>
            <string>com.apple.AssetCache.managed</string>
            <key>PayloadUUID</key>
            <string>c7d8c850-e4ef-0135-0bd6-0c4de9ce4c04</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Content Caching</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>a430c370-b5d5-4e5b-b46b-137931e8b230</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [ContentCaching.Ranges](/documentation/devicemanagement/contentcaching/ranges) - A range of IP addresses to cache.

