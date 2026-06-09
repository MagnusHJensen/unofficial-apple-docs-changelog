# WebContentFilterPluginFilterObject

Settings that control authentication.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Browsers

- **Type:** `WebContentFilterPluginFilter_BrowsersObject`
- **Required:** No

Settings that control the browser filter. If not present, the system doesn’t use browser filtering.

Available: iOS 27+ | iPadOS 27+ | visionOS 27+

### Grade

- **Type:** `string`
- **Required:** No
- **Default:** `firewall`
- **Allowed Values:** `firewall`, `inspector`

The system uses this value to derive the relative order of content filters. Filters with a grade of `firewall` see network traffic before filters with a grade of `inspector`. However, the system doesn’t define the order of filters within a grade.

Available: macOS 27+

### Packets

- **Type:** `WebContentFilterPluginFilter_PacketsObject`
- **Required:** No

Settings that control the packet filter. If not present, the system doesn’t use packet filtering.

Available: macOS 27+

### Sockets

- **Type:** `WebContentFilterPluginFilter_SocketsObject`
- **Required:** No

Settings that control the socket filter. If not present, the system doesn’t use socket filtering.

### URLs

- **Type:** `WebContentFilterPluginFilter_URLsObject`
- **Required:** No

Settings that control the URL filter. If not present, the system doesn’t use URL filtering.

Available: iOS 27+ | iPadOS 27+ | macOS 27+

## Topics

### Objects

- [WebContentFilterPluginFilter_BrowsersObject](/documentation/devicemanagement/webcontentfilterpluginfilter_browsersobject) - Settings that control the browser filter. If not present, the system doesn’t use browser filtering.
- [WebContentFilterPluginFilter_PacketsObject](/documentation/devicemanagement/webcontentfilterpluginfilter_packetsobject) - Settings that control the packet filter. If not present, the system doesn’t use packet filtering.
- [WebContentFilterPluginFilter_SocketsObject](/documentation/devicemanagement/webcontentfilterpluginfilter_socketsobject) - Settings that control the socket filter. If not present, the system doesn’t use socket filtering.
- [WebContentFilterPluginFilter_URLsObject](/documentation/devicemanagement/webcontentfilterpluginfilter_urlsobject) - Settings that control the URL filter. If not present, the system doesn’t use URL filtering.

