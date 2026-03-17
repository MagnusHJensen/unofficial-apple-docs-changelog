# WebContentFilter

The payload that configures web content filters.

**Platforms:** iOS 7.0, iPadOS 7.0, macOS 10.15, visionOS 1.1

## Properties

### AllowListBookmarks

- **Type:** `[WebContentFilter.AllowListBookmarksItem]`
- **Required:** No

An array of dictionaries that define the pages that the user can bookmark or visit. Use when `FilterType` is `BuiltIn`.

### AutoFilterEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables automatic filtering. Use when `FilterType` is `BuiltIn`.

### BlacklistedURLs

- **Type:** `[string]`
- **Required:** No

Use `DenyListURLs` instead.

### ContentFilterUUID

- **Type:** `string`
- **Required:** No

A globally unique identifier for this content filter configuration. The content filter processes network traffic for managed apps with the same `ContentFilterUUID` in their app attributes. Use when `FilterType` is `Plugin`.This key must be present for unsupervised devices and user enrollment.

### DenyListURLs

- **Type:** `[string]`
- **Required:** No

An array of URLs that are inaccessible. Use when `FilterType` is `BuiltIn`. Limit the number of these URLs to no more than 500.

### FilterBrowsers

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables filtering WebKit traffic. Use when `FilterType` is `Plugin`.

> 

### FilterDataProviderBundleIdentifier

- **Type:** `string`
- **Required:** No

The bundle identifier string of the filter data provider system extension. This string identifies the filter data provider when the filter starts running. Required if `FilterSockets` is `true`.

### FilterDataProviderDesignatedRequirement

- **Type:** `string`
- **Required:** No

The designated requirement string that the system embeds in the code signature of the filter data provider system extension. This string identifies the filter data provider when the filter starts running. Required if `FilterSockets` is `true`.

### FilterGrade

- **Type:** `string`
- **Required:** No
- **Default:** `firewall`
- **Allowed Values:** `firewall`, `inspector`

The system uses this value to derive the relative order of content filters. Filters with a grade of `firewall` see network traffic before filters with a grade of `inspector`. However, the system doesn’t define the order of filters within a grade.

### FilterPacketProviderBundleIdentifier

- **Type:** `string`
- **Required:** No

The bundle identifier string of the filter packet provider system extension. This string identifies the filter packet provider when the filter starts running. Required if `FilterPackets` is `true`.

### FilterPacketProviderDesignatedRequirement

- **Type:** `string`
- **Required:** No

The designated requirement string that the system embeds in the code signature of the filter packet provider system extension. This string identifies the filter packet provider when the filter starts running. Required if `FilterPackets` is `true`.

### FilterPackets

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true` and `FilterType` is `Plugin`, the system enables filtering network packets. Use when `FilterType` is `Plugin`.

> 

### FilterSockets

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables the filtering of socket traffic. Use when `FilterType` is `Plugin`.

> 

### FilterType

- **Type:** `string`
- **Required:** No
- **Default:** `BuiltIn`
- **Allowed Values:** `BuiltIn`, `Plugin`

The type of filter, built-in or plug-in. In macOS, the system only supports the plug-in value.

### FilterURLs

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system filters URL requests. Use when `FilterType` is `Plugin`. Available in iOS 26 and macOS 26, and later.

### HideDenyListURLs

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device hides the `DenyListURLs` item in the profiles that display in Settings > General > VPN & Device Management.

### Organization

- **Type:** `string`
- **Required:** No

The organization string to pass to the third-party plug-in. Use when `FilterType` is `Plugin`.

### Password

- **Type:** `string`
- **Required:** No

The password for the service. Use when `FilterType` is `Plugin`.

### PayloadCertificateUUID

- **Type:** `string`
- **Required:** No

The UUID of the certificate payload within the same profile that the system uses to authenticate the user. Use when `FilterType` is `Plugin`.

### PermittedURLs

- **Type:** `[string]`
- **Required:** No

An array or URLs that are accessible whether or not the automatic filter allows access. Use when `FilterType` is `BuiltIn`. Requires that `AutoFilterEnabled` is `true`.

### PluginBundleID

- **Type:** `string`
- **Required:** No

The bundle ID of the plug-in that provides filtering service. Required when `FilterType` is `Plugin`. Otherwise, it ignores this value. Consult your filtering solution vendor to determine what to specify for this value. Required when `FilterType` is `Plugin`.

### SafariHistoryRetentionEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, this payload enforces a policy which requires retention of browsing history. This causes Safari to disable clearing of browsing history, and prevents the use of private browsing mode because that mode doesn’t keep browsing history.

### ServerAddress

- **Type:** `string`
- **Required:** No

The server address, which may be the IP address, hostname, or URL. Use when `FilterType` is `Plugin`.

### URLFilterParameters

- **Type:** `WebContentFilter.URLFilterParameters`
- **Required:** No

A dictionary containing URL filter parameters. Required when `FilterURLs` is `true`. Available in iOS 26 and macOS 26 and later.

### UserDefinedName

- **Type:** `string`
- **Required:** No

The display name for this filtering configuration. Required when `FilterType` is `Plugin`.

### UserName

- **Type:** `string`
- **Required:** No

The user name for the service. Use when `FilterType` is `Plugin`.

### VendorConfig

- **Type:** `WebContentFilter.VendorConfig`
- **Required:** No

The custom dictionary that the filtering service plug-in needs. Use when `FilterType` is `Plugin`.

### WhitelistedBookmarks

- **Type:** `[WebContentFilter.WhitelistedBookmarksItem]`
- **Required:** No

Use `AllowListBookmarks` instead.

## Discussion

Specify `com.apple.webcontent-filter` as the payload type.

The system matches URLs using string-based matching. A URL matches an allow list, deny list, or permitted list pattern if the exact characters of the pattern appear as a substring of the URL requested in the web browser. For example, if the system doesn’t allow `test.com/a`, it blocks `test.com/a`, `test.com/apple`, and `test.com/a/b`.

The system matches list entries that terminate with a `/` character explicitly; if the system blocks or allows `test.com/a/`, it blocks or allows `test.com/a` and `test.com/a/b`.

Matching discards a `www` subdomain prefix if present, so if the system doesn’t allow `www.test.com`, it also blocks `m.test.com`.

All filtering options are active simultaneously. The system only permits URLs and sites that pass all rules.

### Profile availability

### Profile example

```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>PayloadContent</key>
    <array>
        <dict>
            <key>AutoFilterEnabled</key>
            <true/>
            <key>DenylistURLs</key>
            <array>
                <string>https://notallowedname.company.com</string>
            </array>
            <key>FilterBrowsers</key>
            <true/>
            <key>FilterSockets</key>
            <true/>
            <key>FilterType</key>
            <string>BuiltIn</string>
            <key>PermittedURLs</key>
            <array>
                <string>https://example.company.com</string>
            </array>
            <key>PayloadIdentifier</key>
            <string>com.example.mywebcontentfilterpayload</string>
            <key>PayloadType</key>
            <string>com.apple.webcontent-filter</string>
            <key>PayloadUUID</key>
            <string>fb5d598f-0a96-4b77-9702-9edfc3417601</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Web Content Filter</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>b510e0c6-dc81-4b62-88d0-6a3ef82925e7</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [WebContentFilter.AllowListBookmarksItem](/documentation/devicemanagement/webcontentfilter/allowlistbookmarksitem) - The bookmark in the allow list of the web content filter.
- [WebContentFilter.URLFilterParameters](/documentation/devicemanagement/webcontentfilter/urlfilterparameters-data.dictionary) - A dictionary containing URL filter parameters.
- [WebContentFilter.VendorConfig](/documentation/devicemanagement/webcontentfilter/vendorconfig-data.dictionary) - A custom dictionary for the filtering service plug-in.
- [WebContentFilter.WhitelistedBookmarksItem](/documentation/devicemanagement/webcontentfilter/whitelistedbookmarksitem) - The bookmark in the allow list of the web content filter.

