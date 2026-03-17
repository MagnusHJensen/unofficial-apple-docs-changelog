# ExtensibleSingleSignOn

The payload that configures an app extension that performs single sign-on (SSO).

**Platforms:** iOS 13.0, iPadOS 13.0, macOS 10.15, visionOS 1.1

## Properties

### AuthenticationMethod

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Password`, `UserSecureEnclaveKey`

The Platform SSO authentication method the extension uses. Requires that the SSO Extension also supports the method. Available in macOS 13 and later, and deprecated in macOS 14.

### DeniedBundleIdentifiers

- **Type:** `[string]`
- **Required:** No

An array of bundle identifiers of apps that don’t use SSO provided by this extension. Available in iOS 15 and later, and macOS 12 and later.

### ExtensionData

- **Type:** `ExtensibleSingleSignOn.ExtensionData`
- **Required:** No

A dictionary of arbitrary data passed through to the app extension.

### ExtensionIdentifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the app extension that performs SSO for the specified URLs.

### Hosts

- **Type:** `[string]`
- **Required:** No

An array of host or domain names that apps can authenticate through the app extension.

Required for `Credential` payloads. Ignored for `Redirect` payloads.

The system:

- Matches host or domain names case-insensitively
- Requires that all the host and domain names of all installed Extensible SSO payloads are unique

> 

### PlatformSSO

- **Type:** `ExtensibleSingleSignOn.PlatformSSO`
- **Required:** No

The dictionary to configure Platform SSO. Requires `Type` to be set to `Redirect`.

### Realm

- **Type:** `string`
- **Required:** No

The realm name for `Credential` payloads. Use proper capitalization for this value. Ignored for `Redirect` payloads.

### RegistrationToken

- **Type:** `string`
- **Required:** No

The token this device uses for registration with Platform SSO. Use it for silent registration with the Identity Provider. Requires that `AuthenticationMethod` in `PlatformSSO` isn’t empty. Available in macOS 13 and later.

### ScreenLockedBehavior

- **Type:** `string`
- **Required:** No
- **Default:** `Cancel`
- **Allowed Values:** `Cancel`, `DoNotHandle`

If set to `Cancel`, the system cancels authentication requests when the screen is locked. If set to `DoNotHandle`, the request continues without SSO instead. This doesn’t apply to requests where `userInterfaceEnabled` is `false`, or for background [URLSession](/documentation/Foundation/URLSession) requests. Available in iOS 15 and later, and macOS 12 and later.

### TeamIdentifier

- **Type:** `string`
- **Required:** No

The team identifier of the app extension. This key is required on macOS and ignored elsewhere.

### Type

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Credential`, `Redirect`

The type of SSO.

### URLs

- **Type:** `[string]`
- **Required:** No

An array of URL prefixes of identity providers where the app extension performs SSO.

Required for `Redirect` payloads. Ignored for `Credential` payloads.

The URLs need to begin with `http://` or `https://`.

The system:

- Matches scheme and host name case-insensitively
- Doesn’t allow query parameters and URL fragments
- Requires that the URLs of all installed Extensible SSO payloads are unique

## Discussion

Specify `com.apple.extensiblesso` as the payload type.

The system supports user channel installation in macOS 11 and later.

### Profile availability

### Profile example

```plist
<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>PayloadContent</key>
    <array>
        <dict>
            <key>ExtensionData</key>
            <dict>
                <key>useSiteAutoDiscovery</key>
                <true/>
            </dict>
            <key>ExtensionIdentifier</key>
            <string>com.example.com</string>
            <key>TeamIdentifier</key>
            <string>RandomTeamID</string>
            <key>Hosts</key>
            <array>
                <string>.com.example.com</string>
            </array>
            <key>Realm</key>
            <string>COM.URL.COM</string>
            <key>Type</key>
            <string>Credential</string>
            <key>PayloadIdentifier</key>
            <string>com.example.myessopayload</string>
            <key>PayloadType</key>
            <string>com.apple.extensiblesso</string>
            <key>PayloadUUID</key>
            <string>dbed949d-39a2-440d-a84b-e0c825cdcb2e</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Extensible SSO</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>da3bbbec-a753-4aa7-aeae-a74b7a65c0b5</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [ExtensibleSingleSignOn.ExtensionData](/documentation/devicemanagement/extensiblesinglesignon/extensiondata-data.dictionary) - The additional data to pass to the app extension.
- [ExtensibleSingleSignOn.PlatformSSO](/documentation/devicemanagement/extensiblesinglesignon/platformsso-data.dictionary) - The dictionary to configure Platform SSO. Requires `Type` to be set to `Redirect`.

