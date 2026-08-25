# ScheduleOSUpdateResponse.UpdateResultsItem

The response dictionary that describes the result of processing an operating-system update.

**Platforms:** iOS 9.0, iPadOS 9.0, Mac Catalyst 9.0, macOS 10.11, tvOS 12.0

## Properties

### ErrorChain

- **Type:** `[ScheduleOSUpdateResponse.UpdateResultsItem.ErrorChainItem]`
- **Required:** No




Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### InstallAction

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Error`, `DownloadOnly`, `InstallASAP`, `NotifyOnly`, `InstallLater`, `InstallForceRestart`



     


Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### ProductKey

- **Type:** `string`
- **Required:** Yes




Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### Status

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Idle`, `Downloading`, `DownloadFailed`, `DownloadRequiresComputer`, `DownloadInsufficientSpace`, `DownloadInsufficientPower`, `DownloadInsufficientNetwork`, `Installing`, `InstallInsufficientSpace`, `InstallInsufficientPower`, `InstallPhoneCallInProgress`, `InstallFailed`



           


Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

## Topics

### Objects

- [ScheduleOSUpdateResponse.UpdateResultsItem.ErrorChainItem](/documentation/devicemanagement/scheduleosupdateresponse/updateresultsitem/errorchainitem) - A dictionary that describes an error chain item.

