# ExtensibleSSO

The declaration to configure Extensible Single Sign-On.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### DeniedBundleIdentifiers

- **Type:** `[string]`
- **Required:** No

An array of bundle identifiers of apps that don’t use SSO provided by this extension.

### ExtensionComposedIdentifier

- **Type:** `string`
- **Required:** Yes

The identifier of the provider to use for this configuration. Useful for apps that contain more than one DNS proxy extension.

In iOS and visionOS, the identifier is a bundle ID, for example, “com.example.app.sso-extension”.

In macOS, the identifier is a composed identifier. The format of the composed identifier is “Bundle-ID (Team-ID)”. “Bundle-ID” is the bundle identifier string of the app extension. “Team-ID” is the team identifier from the app extension’s code signature. For example, “com.example.app.sso-extension (ABCD1234)”.

### ExtensionData

- **Type:** `ExtensibleSSOExtensionDataObject`
- **Required:** No

A dictionary of arbitrary data passed through to the app extension.

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

- **Type:** `ExtensibleSSOPlatformSSOObject`
- **Required:** No

The dictionary to configure Platform SSO.

Available: macOS 27+

### Realm

- **Type:** `string`
- **Required:** No

The realm name for `Credential` payloads. Use proper capitalization for this value. Ignored for `Redirect` payloads.

### ScreenLockedBehavior

- **Type:** `string`
- **Required:** No
- **Default:** `Cancel`
- **Allowed Values:** `Cancel`, `DoNotHandle`

If set to `Cancel`, the system cancels authentication requests when the screen is locked. If set to `DoNotHandle`, the request continues without SSO instead. This doesn’t apply to requests where `userInterfaceEnabled` is `false`, or for background [URLSession](/documentation/Foundation/URLSession) requests.

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

Specify `com.apple.configuration.extensible-sso` as the declaration type.

### Configuration availability

### Configuration Examples

## Topics

### Objects

- [ExtensibleSSOExtensionDataObject](/documentation/devicemanagement/extensiblessoextensiondataobject) - A dictionary of arbitrary data passed through to the app extension.
- [ExtensibleSSOPlatformSSOObject](/documentation/devicemanagement/extensiblessoplatformssoobject) - The dictionary to configure Platform SSO.

