# EnrollmentSSODocument

Enrollment SSO streamlines the MDM enrollment process, reduces sign-ins, and improves security.

**Platforms:** iOS 16.0, iPadOS 16.0, Mac Catalyst 16.0, visionOS 2.0, Device Assignment Services , VPP License Management 

## Properties

### AppIDs

- **Type:** `[string]`
- **Required:** No

An array of App IDs that specify apps that Enrollment SSO developer mode can use. In Enrollment SSO documents delivered through the developer endpoint, this key must be present and contain at least one value. In Enrollment SSO documents delivered by the standard Enrollment SSO endpoint, this key must not be present.

### AssociatedDomains

- **Type:** `[string]`
- **Required:** No

An array of associated domains that the device uses with the Enrollment SSO extension.

### AssociatedDomainsEnableDirectDownloads

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true,` allows the domain to directly verify site association, instead of at Apple’s servers. Use this verification only with domains that are inaccessible on the public Internet.

### ConfigurationProfile

- **Type:** `data`
- **Required:** No

The profile containing an [ExtensibleSingleSignOn](/documentation/devicemanagement/extensiblesinglesignon) payload that specifies the SSO extension in the downloaded app prior to enrollment. This profile may contain certificate payloads.

One of `ConfigurationProfile` and `Declarations` must be present.

### Declarations

- **Type:** `[data]`
- **Required:** No

An array of base64-encoded JSON formatted Declarative Device Management declarations that specify the managed app and its configuration, including any certificates or identities.

The set of declarations must include one `com.apple.configuration.app.managed` configuration, and one activation declaration that references the configuration. Asset declarations may be present if required by the app config.

The app configuration must include `AppStoreID` when developer mode is not being used, or it must include `BundleID` when developer mode is used.

One of `ConfigurationProfile` and `Declarations` must be present.

### iTunesStoreID

- **Type:** `integer`
- **Required:** No

The iTunes Store ID of the app to download prior to enrollment, to support Enrollment SSO during enrollment. Using developer mode ignores this key.

