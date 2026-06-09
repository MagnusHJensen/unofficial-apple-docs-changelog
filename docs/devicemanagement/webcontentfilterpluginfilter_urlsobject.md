# WebContentFilterPluginFilter_URLsObject

Settings that control the URL filter. If not present, the system doesn’t use URL filtering.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta)

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

