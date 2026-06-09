# SecurityInfoResponse.SecurityInfo

A dictionary that contains security-related information.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### AuthenticatedRootVolumeEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the system booted using an Authenticated Root Volume.

Available: macOS 11+

### AutoLockTime

- **Type:** `integer`
- **Required:** No

The number of seconds before a device goes to sleep after being idle. This value is only available on Shared iPad in iOS 17 and later.

Available: iOS 17+ | iPadOS 17+

### BootstrapTokenAllowedForAuthentication

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `allowed`, `disallowed`, `not supported`

This value specifies whether the Secure Enclave Processor (SEP) supports and allows secure operations to use the Bootstrap Token. The value is automatically set for devices enrolled through Automated Device Enrollment (ADE). The user can also manually set this value in the RecoveryOS.

This value is available for a Mac with Apple silicon in macOS 11 and later. Not available for user enrollment.

Available: macOS 11+

### BootstrapTokenRequiredForKernelExtensionApproval

- **Type:** `boolean`
- **Required:** No

If `true`, the device can accept a Bootstrap Token from the MDM server instead of prompting for user authentication prior to enabling kernel extensions. This includes enabling kexts through the `com.apple.syspolicy.kernel-extension-policy` payload or triggering the `RestartDevice` command with `RebuildKernelCache` set to `true`. This only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the [SecurityInfoResponse.SecurityInfo](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary) response.

This value is available for a Mac with Apple silicon in macOS 11 and later. Not available for user enrollment.

Available: macOS 11+

### BootstrapTokenRequiredForSoftwareUpdate

- **Type:** `boolean`
- **Required:** No

If `true`, the device can accept a Bootstrap Token from the MDM server instead of prompting for user authentication prior to installation. This only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the [SecurityInfoResponse.SecurityInfo](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary) response.

This value is available for a Mac with Apple silicon in macOS 11 and later. Not available for user enrollment.

Available: macOS 11+

### FDE_Enabled

- **Type:** `boolean`
- **Required:** No

If `true`, the device has enabled FileVault full disk encryption (FDE).

Available: macOS 10.9+

### FDE_HasInstitutionalRecoveryKey

- **Type:** `boolean`
- **Required:** No

If `true`, FileVault FDE has an institutional recovery key.

Available: macOS 10.9+

### FDE_HasPersonalRecoveryKey

- **Type:** `boolean`
- **Required:** No

If `true`, FileVault FDE has a personal recovery key.

Available: macOS 10.9+

### FDE_PersonalRecoveryKeyCMS

- **Type:** `data`
- **Required:** No

If the FileVault personal recovery key has enabled escrow with a recovery key, this value contains the key. The certificate from the [FDERecoveryKeyEscrow](/documentation/devicemanagement/fderecoverykeyescrow) profile encrypts the key and wraps it as CMS data.

Available: macOS 10.13+

### FDE_PersonalRecoveryKeyDeviceKey

- **Type:** `string`
- **Required:** No

If the FileVault personal recovery key has enabled escrow with a recovery key, this value is the device serial number. This is the value that displays to the user at the EFI Login Window as part of the help message if they enter their password incorrectly three times. The server also uses this value as an index when saving the device personal recovery key. This replaces the `recordNumber` that the server returned in the previous escrow mechanism.

Available: macOS 10.13+

### FirewallSettings

- **Type:** `SecurityInfoResponse.SecurityInfo.FirewallSettings`
- **Required:** No

A dictionary that contains the firewall settings.

Available: macOS 10.12+

### FirmwarePasswordStatus

- **Type:** `SecurityInfoResponse.SecurityInfo.FirmwarePasswordStatus`
- **Required:** No

A dictionary that contains the status of the EFI firmware password.

Available: macOS 10.13+

### HardwareEncryptionCaps

- **Type:** `integer`
- **Required:** No

An integer that indicates the underlying hardware encryption capabilities of the device, which is one of the following values:

- `1`: Block-level encryption
- `2`: File-level encryption
- `3`: Both block-level and file-level encryption

> 

Available: iOS 4+ | iPadOS 4+ | tvOS 9+ | visionOS 1.1+ | watchOS 10+

### IsRecoveryLockEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, a password is required to enter recovery (see [SetRecoveryLockCommand](/documentation/devicemanagement/setrecoverylockcommand)). Available in macOS 11.5 and later and only on a Mac with Apple silicon.

Available: macOS 11.5+

### ManagementStatus

- **Type:** `SecurityInfoResponse.SecurityInfo.ManagementStatus`
- **Required:** No

A dictionary that contains the status of the device’s MDM enrollment.

Available: iOS 13+ | iPadOS 13+ | macOS 10.13.2+ | tvOS 13+ | visionOS 1.1+ | watchOS 10+

### PasscodeCompliant

- **Type:** `boolean`
- **Required:** No

If `true`, the user’s passcode is compliant with all requirements on the device, including Exchange and other accounts.

Available: iOS 4+ | iPadOS 4+ | tvOS 9+ | visionOS 1.1+ | watchOS 10+

### PasscodeCompliantWithProfiles

- **Type:** `boolean`
- **Required:** No

If `true`, the user’s passcode is compliant with requirements from profiles. This key doesn’t apply to User-Enrolled devices.

Available: iOS 4+ | iPadOS 4+ | tvOS 9+ | visionOS 1.1+ | watchOS 10+

### PasscodeLockGracePeriod

- **Type:** `integer`
- **Required:** No

The user preference for the number of seconds before a locked screen requires the device passcode to unlock it. This value is only available for Shared iPad.

Available: iOS 9.3.2+ | iPadOS 9.3.2+ | tvOS 9+ | visionOS 1.1+ | watchOS 10+

### PasscodeLockGracePeriodEnforced

- **Type:** `integer`
- **Required:** No

The enforced value for the number of seconds before a locked screen requires the device passcode to unlock it. If a device has a passcode, changing `PasscodeLockGracePeriod` to a larger value doesn’t take effect until the user logs out or removes the passcode. This value is only available for Shared iPad.

Available: iOS 9.3.2+ | iPadOS 9.3.2+ | tvOS 9+ | visionOS 1.1+ | watchOS 10+

### PasscodePresent

- **Type:** `boolean`
- **Required:** No

If `true`, the device has a passcode. This key doesn’t apply to User-Enrolled devices.

Available: iOS 4+ | iPadOS 4+ | tvOS 9+ | visionOS 1.1+ | watchOS 10+

### RemoteDesktopEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, Remote Desktop is active on the device.

Available: macOS 10.14.4+

### SecureBoot

- **Type:** `SecurityInfoResponse.SecurityInfo.SecureBoot`
- **Required:** No

A dictionary that contains the device’s Secure Boot settings.

Available: macOS 10.15+

### SystemIntegrityProtectionEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, System Integrity Protection (SIP) is active on the device.

Available: macOS 10.12+

## Topics

### Objects

- [SecurityInfoResponse.SecurityInfo.FirewallSettings](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary/firewallsettings-data.dictionary) - A dictionary that contains the firewall settings.
- [SecurityInfoResponse.SecurityInfo.FirmwarePasswordStatus](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary/firmwarepasswordstatus-data.dictionary) - A dictionary that contains the status of the EFI firmware password.
- [SecurityInfoResponse.SecurityInfo.ManagementStatus](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary/managementstatus-data.dictionary) - A dictionary that contains the status of the device’s MDM enrollment.
- [SecurityInfoResponse.SecurityInfo.SecureBoot](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary/secureboot-data.dictionary) - The response object for the secure boot settings.

