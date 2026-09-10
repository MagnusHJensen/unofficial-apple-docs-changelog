# WebContentFilterPluginFilter_URLsObject

Settings that control the URL filter. If not present, the system doesn’t use URL filtering.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0

## Properties

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, the system filters URL requests.

### Parameters

- **Type:** `WebContentFilterPluginFilter_URLs_ParametersObject`
- **Required:** No

A dictionary containing URL filter parameters. Required when `Enabled` is `true`.

## Topics

### Objects

- [WebContentFilterPluginFilter_URLs_ParametersObject](/documentation/devicemanagement/webcontentfilterpluginfilter_urls_parametersobject) - A dictionary containing URL filter parameters. Required when `Enabled` is `true`.

