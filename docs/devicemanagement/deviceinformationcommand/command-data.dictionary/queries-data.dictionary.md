# DeviceInformationCommand.Command.Queries

An array of query dictionaries to get information about a device.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### AccessibilitySettings

- **Type:** `string`
- **Required:** No

The key to get the current state of settable accessibility settings. Available in iOS 16 and later.

### ActiveManagedUsers

- **Type:** `string`
- **Required:** No

The key to get an array of directory GUIDs for logged-in managed users. Requires the Device Information access right. Available in macOS 10.11 and later.

### AppAnalyticsEnabled

- **Type:** `string`
- **Required:** No

The key to determine whether the device is sharing app analytics. Requires the Device Information access right. Available in iOS 4 and later, and macOS 10.7 and later.

### AutoSetupAdminAccounts

- **Type:** `string`
- **Required:** No

The key to get the contents of [DeviceInformationResponse.QueryResponses.AutoSetupAdminAccountsItem](/documentation/devicemanagement/deviceinformationresponse/queryresponses-data.dictionary/autosetupadminaccountsitem), which Setup Assistant automatically creates during enrollment. Requires the Device Information access right. Available in macOS 10.11 and later.

### AvailableDeviceCapacity

- **Type:** `string`
- **Required:** No

The key to get the available capacity. Requires the Device Information access right. Available in iOS 4 and later, and macOS 10.7 and later.

### AwaitingConfiguration

- **Type:** `string`
- **Required:** No

The key to determine whether the device is waiting for a [Device Configured](/documentation/devicemanagement/device-configured-command) command or [User Configured](/documentation/devicemanagement/user-configured-command) command to continue through Setup Assistant on the device channel or user channel, respectively.

### BatteryLevel

- **Type:** `string`
- **Required:** No

The key to get the battery level. Requires the Device Information access right. Available in iOS 5 and later.

### BluetoothMAC

- **Type:** `string`
- **Required:** No

The key to get the Bluetooth media access control (MAC) address. Requires the Network Information access right.

### BuildVersion

- **Type:** `string`
- **Required:** No

The key to get the operating system version. Requires the Device Information access right.

### CellularTechnology

- **Type:** `string`
- **Required:** No

The key to get the cellular technology type. Requires the Device Information access right. Available in iOS 4.2.6 and later.

### DataRoamingEnabled

- **Type:** `string`
- **Required:** No

The key to determine whether the system enabled data roaming on the device. Requires the Network Information access right. Available in iOS 5 and later.

### DeviceCapacity

- **Type:** `string`
- **Required:** No

The key to get the device’s total capacity. Requires the Device Information access right. Available in iOS 4 and later, and macOS 10.7 and later.

### DeviceID

- **Type:** `string`
- **Required:** No

The key to get the device ID. Requires the Device Information access right. Available in tvOS 6 and later.

### DeviceName

- **Type:** `string`
- **Required:** No

The key to get the device name. Requires the Device Information access right.

### DevicePropertiesAttestation

- **Type:** `string`
- **Required:** No

The key to get an attestation of the device’s properties. Available in iOS 16 and later, macOS 14 and later, tvOS 16 and later, and watchOS 10 and later. The hardware requirements for attestation are described below.

### DiagnosticSubmissionEnabled

- **Type:** `string`
- **Required:** No

The key to determine whether the system enabled the diagnostic submission setting on the device. Requires the Device Information access right. Available in iOS 9.3 and later.

### EACSPreflight

- **Type:** `string`
- **Required:** No

The key to determine whether the device can perform an [EraseDeviceCommand](/documentation/devicemanagement/erasedevicecommand) using Erase All Content and Settings (EACS).

### EASDeviceIdentifier

- **Type:** `string`
- **Required:** No

The key to get the device identifier for Exchange ActiveSync (EAS). Requires the Device Information access right. Available in iOS 7 and later.

### EstimatedResidentUsers

- **Type:** `string`
- **Required:** No

The key to get the estimated number of users that can use this Shared iPad device, according to the available space of the device and each user’s quota. Requires the Device Information access right. Available in iOS 14 and later.

### EthernetMAC

- **Type:** `string`
- **Required:** No

The key to get the primary Ethernet MAC address. Requires the Network Information access right. Available in macOS 10.7 and later.

### HasBattery

- **Type:** `string`
- **Required:** No

The key to determine whether the device has an internal battery.

### HostName

- **Type:** `string`
- **Required:** No

The key to get the hostname. Available in macOS 10.11 and later.

### IsActivationLockEnabled

- **Type:** `string`
- **Required:** No

The key to determine whether the system enabled Activation Lock on the device. Requires the Device Information access right. Available as of iOS 7 and macOS 10.15, and deprecated in iOS 16 and macOS 13.

### IsActivationLockSupported

- **Type:** `string`
- **Required:** No

The key to determine whether the device supports Activation Lock. Also see `IsActivationLockManageable` in [SecurityInfoResponse.SecurityInfo.ManagementStatus](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary/managementstatus-data.dictionary). Available in macOS 10.9 and later.

### IsAppleSilicon

- **Type:** `string`
- **Required:** No

The key to determine whether the device is a Mac with Apple silicon (for example, an Apple M1 chip). Available in macOS 12 and later.

### IsCloudBackupEnabled

- **Type:** `string`
- **Required:** No

The key to determine whether the system enabled iCloud Backup on the device. Requires the Device Information access right. Available in iOS 7.1 and later.

### IsDeviceLocatorServiceEnabled

- **Type:** `string`
- **Required:** No

The key to determine whether the system enabled a device locator service such as Find My on the device. Requires the Device Information access right. Available in iOS 7 and later.

### IsDoNotDisturbInEffect

- **Type:** `string`
- **Required:** No

The key to determine whether the device is in Do Not Disturb (DND) mode. Requires the Device Information access right. Available in iOS 7 and later.

### IsMDMLostModeEnabled

- **Type:** `string`
- **Required:** No

The key to determine whether the system enabled Managed Lost Mode on the device. Requires the Device Information access right. Available in iOS 9.3 and later.

### IsMultiUser

- **Type:** `string`
- **Required:** No

The key to determine whether the device is a Shared iPad. Requires the Device Information access right. Available in iOS 9.3 and later.

### IsNetworkTethered

- **Type:** `string`
- **Required:** No

The key to determine whether the device is network-tethered. Requires the Network Information access right. Available in iOS 10.3 and later.

### IsSupervised

- **Type:** `string`
- **Required:** No

The key to determine whether the device is supervised. Requires the Device Information access right. Available in iOS 6 and later, macOS 10.15 and later, and tvOS 9 and later.

### iTunesStoreAccountHash

- **Type:** `string`
- **Required:** No

The key to get a hash of the logged-in iTunes Store account. Also see [GetVppUserRequest](/documentation/devicemanagement/getvppuserrequest). Requires the App Installation access right.

### iTunesStoreAccountIsActive

- **Type:** `string`
- **Required:** No

The key to determine whether the iTunes Store account is active. Requires the App Installation access right.

### LastCloudBackupDate

- **Type:** `string`
- **Required:** No

The key to get the date of the most-recent iCloud backup. Available in iOS 8 and later.

### LocalHostName

- **Type:** `string`
- **Required:** No

The key to get the local hostname from Bonjour. Available in macOS 10.11 and later.

### ManagedAppleIDDefaultDomains

- **Type:** `string`
- **Required:** No

The key to get the list of domains that the device suggests on the Shared iPad login screen. Available in iOS 16 and later.

### MaximumResidentUsers

- **Type:** `string`
- **Required:** No

The key to get the maximum number of users that can use this Shared iPad device. In iOS 13.4 and later, this value is always `32`. Requires the Device Information access right. Available in iOS 9.3 and later.

### MDMOptions

- **Type:** `string`
- **Required:** No

The key to get the contents of [SettingsCommand.Command.Settings.MDMOptions.MDMOptions](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/mdmoptions-data.dictionary/mdmoptions-data.dictionary).

### Model

- **Type:** `string`
- **Required:** No

The key to get the model. Requires the Device Information access right.

### ModelName

- **Type:** `string`
- **Required:** No

The key to get the model name, such as **. Requires the Device Information access right.

### ModelNumber

- **Type:** `string`
- **Required:** No

The key to get the device’s hardware model number, including region info, such as `MK1A3LL/A`. Requires the Device Information access right. Requires a Mac with Apple silicon for macOS.

### ModemFirmwareVersion

- **Type:** `string`
- **Required:** No

The key to get the modem firmware version. Requires the Device Information access right. Available in iOS 4 and later.

### OnlineAuthenticationGracePeriod

- **Type:** `string`
- **Required:** No

The key to get the grace period for Shared iPad online authentication (in days). Available in iOS 16 and later.

### OrganizationInfo

- **Type:** `string`
- **Required:** No

The key to get the contents of [SettingsCommand.Command.Settings.OrganizationInfo.OrganizationInfo](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/organizationinfo-data.dictionary/organizationinfo-data.dictionary).

### OSUpdateSettings

- **Type:** `string`
- **Required:** No

The key to get the contents of [DeviceInformationResponse.QueryResponses.OSUpdateSettings](/documentation/devicemanagement/deviceinformationresponse/queryresponses-data.dictionary/osupdatesettings-data.dictionary). Requires the Device Information access right. Available in macOS 10.11 and later.

### OSVersion

- **Type:** `string`
- **Required:** No

The key to get the operating system version. Requires the Device Information access right.

### PersonalHotspotEnabled

- **Type:** `string`
- **Required:** No

The key to determine whether the system enabled Personal Hotspot on the device, which isn’t available for all carriers. Requires the Network Information access right. Available in iOS 7 and later.

### PINRequiredForDeviceLock

- **Type:** `string`
- **Required:** No

The key to determine whether the [DeviceLockCommand](/documentation/devicemanagement/devicelockcommand) requires a PIN. Available in macOS 11 and later.

### PINRequiredForEraseDevice

- **Type:** `string`
- **Required:** No

The key to determine whether the [EraseDeviceCommand](/documentation/devicemanagement/erasedevicecommand) requires a PIN. Available in macOS 11 and later.

### ProductName

- **Type:** `string`
- **Required:** No

The key to get the product name, such as **. Requires the Device Information access right.

### ProvisioningUDID

- **Type:** `string`
- **Required:** No

The key to get the device identifier for provisioning profiles. This value differs from the UDID for a Mac with Apple silicon. Available in macOS 11.3 and later.

### PushToken

- **Type:** `string`
- **Required:** No

The key to get the push token for the current user-channel connection. The MDM server ignores this query for the device channel. Requires the Device Information access right. Available in iOS 9.3 and later, and macOS 10.12 and later.

### QuotaSize

- **Type:** `string`
- **Required:** No

The key to get the quota size for each user on this Shared iPad device. Requires the Device Information access right. Available in iOS 13.4 and later.

### ResidentUsers

- **Type:** `string`
- **Required:** No

The key to get the number of users currently on this Shared iPad device. Requires the Device Information access right. Available in iOS 13.4 and later.

### SerialNumber

- **Type:** `string`
- **Required:** No

The key to get the serial number. Requires the Device Information access right.

### ServiceSubscriptions

- **Type:** `string`
- **Required:** No

The key to get the contents of [DeviceInformationResponse.QueryResponses.ServiceSubscriptionProperty](/documentation/devicemanagement/deviceinformationresponse/queryresponses-data.dictionary/servicesubscriptionproperty). Requires the Network Information access right.

### SkipLanguageAndLocaleSetupForNewUsers

- **Type:** `string`
- **Required:** No

The key to determine whether the system skips the language and country/region panes for new users on Shared iPad.

### SoftwareUpdateDeviceID

- **Type:** `string`
- **Required:** No

The key to get the device identifier to look up available OS updates through [https://gdmf.apple.com/v2/pmv](https://gdmf.apple.com/v2/pmv). Available in iOS 15 and later, and macOS 12 and later.

### SoftwareUpdateSettings

- **Type:** `string`
- **Required:** No

The key to get the device settings that control which updates appear in the Software Update pane in Settings. Available in iOS 14.5 and later.

### SupplementalBuildVersion

- **Type:** `string`
- **Required:** No

The key to get the build version for the currently installed Background Security Improvement. If there’s no installed Background Security Improvement, this value is the same as `BuildVersion`. Requires the Device Information access right.

### SupplementalOSVersionExtra

- **Type:** `string`
- **Required:** No

The key to get the OS update Background Security Improvement version letter, if a Background Security Improvement update is installed. Requires the Device Information access right.

### SupportsiOSAppInstalls

- **Type:** `string`
- **Required:** No

The key to determine whether the macOS device supports iOS or iPadOS app installs. Available in macOS 11 and later.

### SupportsLOMDevice

- **Type:** `string`
- **Required:** No

The key to determine whether the device can receive `PowerON`, `PowerOFF`, and `Reset` commands from a lights-out management (LOM) controller. Available in macOS 11 and later.

### SystemIntegrityProtectionEnabled

- **Type:** `string`
- **Required:** No

The key to determine whether the system enabled System Integrity Protection on the device. Requires the Device Information access right, and is available in macOS 10.12 and later.

### TemporarySessionOnly

- **Type:** `string`
- **Required:** No

The key to determine whether the device allows only temporary sessions.

### TemporarySessionTimeout

- **Type:** `string`
- **Required:** No

The key to get the timeout interval for the temporary session.

### TimeZone

- **Type:** `string`
- **Required:** No

The key to get the current Internet Assigned Numbers Authority (IANA) time zone database name. Requires the Device Information access right. Available in macOS 26 and later, iOS 14 and later, tvOS 14 and later, and visionOS 2 and later.

### UDID

- **Type:** `string`
- **Required:** No

The key to get the unique identifier of the device.

### UserSessionTimeout

- **Type:** `string`
- **Required:** No

The key to get the timeout interval for the user session.

### WiFiMAC

- **Type:** `string`
- **Required:** No

The key to get the Wi-Fi MAC address. Requires the Network Information access right.

