# WebContentFilterPlugin

The declaration to configure a WebContent Filter that uses a plugin.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Authentication

- **Type:** `WebContentFilterPluginAuthenticationObject`
- **Required:** No

Settings that control authentication.

### ContentFilterUUID

- **Type:** `string`
- **Required:** No

A globally unique identifier for this content filter configuration. The content filter processes network traffic for managed apps with the same `ContentFilterUUID` in their app attributes. This key must be present for unsupervised devices and user enrollment.

Available: iOS 27+ | iPadOS 27+ | visionOS 27+

### Filter

- **Type:** `WebContentFilterPluginFilterObject`
- **Required:** No

Settings that control authentication.

### Organization

- **Type:** `string`
- **Required:** No

The organization string to pass to the third-party plug-in.

### PluginBundleID

- **Type:** `string`
- **Required:** Yes

The bundle ID of the plug-in that provides filtering service. Consult your filtering solution vendor to determine what to specify for this value.

### ServerAddress

- **Type:** `string`
- **Required:** No

The server address, which may be the IP address, hostname, or URL.

### VendorConfig

- **Type:** `WebContentFilterPluginVendorConfigObject`
- **Required:** No

The custom dictionary that the filtering service plug-in needs.

### VisibleName

- **Type:** `string`
- **Required:** Yes

The name of the web content filter that the system displays on the device.

## Discussion

Specify `com.apple.configuration.webcontent-filter.plugin` as the declaration type.

### Configuration availability

### Configuration example

This configuration sets up a web content filter using a Network Extension plugin bundle.

```json
{
    "Type": "com.apple.configuration.webcontent-filter.plugin",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "VisibleName": "Content Filter",
        "PluginBundleID": "com.example.contentfilter",
        "ServerAddress": "filter.example.com",
        "ContentFilterUUID": "A1B2C3D4-E5F6-7890-ABCD-EF1234567890"
    }
}
```

## Topics

### Objects

- [WebContentFilterPluginAuthenticationObject](/documentation/devicemanagement/webcontentfilterpluginauthenticationobject) - Settings that control authentication.
- [WebContentFilterPluginFilterObject](/documentation/devicemanagement/webcontentfilterpluginfilterobject) - Settings that control authentication.
- [WebContentFilterPluginVendorConfigObject](/documentation/devicemanagement/webcontentfilterpluginvendorconfigobject) - The custom dictionary that the filtering service plug-in needs.

