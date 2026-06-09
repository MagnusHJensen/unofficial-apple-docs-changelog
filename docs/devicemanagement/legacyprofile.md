# LegacyProfile

The declaration to configure a legacy profile.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### ProfileAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration containing a reference to the profile data. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset’s “ContentType” value set to match the data type.

One of `ProfileURL` or `ProfileAssetReference` needs to be present.

Available: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+ | visionOS 27+ | watchOS 27+

### ProfileURL

- **Type:** `string`
- **Required:** No

The URL of the profile to download and install, which needs to start with `https://`. The request uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`.

One of `ProfileURL` or `ProfileAssetReference` needs to be present.

## Discussion

Specify `com.apple.configuration.legacy` as the declaration type.

This declaration specifies an MDMv1 profile for the device to download and install.

The profile may contain any payload type other than the following:

- `com.apple.mdm`
- `com.apple.declarations`

If a user enrollment triggers this configuration: in macOS the system silently ignores any MDMv1 payloads in macOS where the User Enrollment Mode setting is `forbidden`; in iOS, tvOS, watchOS and visionOS, the system rejects the entire profile if any MDMv1 payload has its User Enrollment Mode setting set to `forbidden`.

### Transition profiles from MDM

A declarative device management (DDM) legacy profile can take control of profiles installed via MDM. This avoids the need to first remove the MDM profile, before installing the DDM equivalent. DDM cannot take over control of non-MDM-installed profiles.

The rules for transitioning profiles are:

1. An existing MDM-installed profile is present (installed by MDM using the [Install Profile](/documentation/devicemanagement/install-profile-command)).
2. DDM is enabled on the device.
3. The server sends a legacy profile configuration to the device and ensures it is “activated”.

When DDM takes control of the MDM profile, the following occurs:

1. The system doesn’t reinstall the profile. Instead, the MDM profile’s existing system state remains unchanged. Thus system state won’t include any differences between the MDM and DDM profiles (other than the structural items outlined above that must match).
2. Any attempt to install, update, or remove the profile using MDM commands fails (using the usual identifier and UUID matching rules). This holds true while the DDM profile is active.
3. Updates to the DDM configuration result in the system reapplying the profile which updates the system state with any new or changed settings.

### Configuration availability

### Configuration examples

