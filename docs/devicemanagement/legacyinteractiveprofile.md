# LegacyInteractiveProfile

The declaration to configure an interactive legacy profile.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1

## Properties

### ProfileAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration containing a reference to the profile data. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset’s “ContentType” value set to match the data type.

One of `ProfileURL` or `ProfileAssetReference` needs to be present.

Available: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+ | visionOS 27+

### ProfileURL

- **Type:** `string`
- **Required:** No

The URL of the profile to download and install, which needs to start with `https://`. The request uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`.

One of `ProfileURL` or `ProfileAssetReference` needs to be present.

### VisibleName

- **Type:** `string`
- **Required:** Yes

The visible name of the configuration. This name needs to indicate the nature of the profile.

## Discussion

Specify `com.apple.configuration.legacy.interactive` as the declaration type.

This declaration specifies an MDMv1 profile to present to the user, who may choose to download and install the profile.

The profile may contain any payload type other than the following:

- `com.apple.mdm`
- `com.apple.declarations`

If a user enrollment triggers this configuration: in macOS the system silently ignores any MDMv1 payloads in macOS where the User Enrollment Mode setting is `forbidden`; in iOS, tvOS, watchOS and visionOS, the system rejects the entire profile if any MDMv1 payload has its User Enrollment Mode setting set to `forbidden`.

### Configuration availability

### Configuration examples

