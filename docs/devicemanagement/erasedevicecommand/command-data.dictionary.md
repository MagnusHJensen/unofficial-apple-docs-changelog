# EraseDeviceCommand.Command

The command to remotely and immediately erase a device.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### DisallowProximitySetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, disable Proximity Setup on the next reboot and skip the pane in Setup Assistant. This value is available in iOS 11 and later. Prior to iOS 14, don’t use this option with any other option.

### ObliterationBehavior

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Default`, `DoNotObliterate`, `ObliterateWithWarning`, `Always`

This key defines the fallback behavior for erasing a device.

In macOS 12 and later, this command uses Erase All Content and Settings (EACS) on Mac computers with the Apple M1 chip or the Apple T2 Security Chip. On those devices, if EACS can’t run, the device can use obliteration (macOS 11.x behavior). This key has no effect on machines prior to the T2 chip. For a list of supported macs, see [Mac models with the Apple T2 Security Chip](https://support.apple.com/en-us/HT208862).

Upon receiving this command, the device performs preflight checks to determine if the device is in a state that allows EACS. The `status` of the [EraseDeviceResponse](/documentation/devicemanagement/erasedeviceresponse) is either `Acknowledged` or `Error`.

The following values define the device’s fallback behavior:

- `DoNotObliterate`: If EACS preflight fails, the device responds to the server with an `Error` status and doesn’t attempt to erase itself. If EACS preflight succeeds, but EACS fails, the device doesn’t attempt to erase itself.
- `ObliterateWithWarning`: If EACS preflight fails, the device responds with an `Acknowledged` status and then attempts to erase itself. If EACS preflight succeeds, but EACS fails, the device attempts to erase itself.
- `Always`: The system doesn’t attempt EACS. T2 and later devices always obliterate.
- `Default`: If EACS preflight fails, the device responds to the server with an `Error` status and then attempts to erase itself. If EACS preflight succeeds, but EACS fails, the device attempts to erase itself.

### PIN

- **Type:** `string`
- **Required:** No

The six-character PIN for Find My. This value is available in macOS 10.8 and later.

### PreserveDataPlan

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, preserve the data plan on an iPhone or iPad with eSIM functionality, if one exists. This value is available in iOS 11 and later.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `EraseDevice`

The request type for this command.

### ReturnToService

- **Type:** `EraseDeviceCommand.Command.ReturnToService`
- **Required:** No

The configuration settings for return to service. This value is available in iOS 17 and later, with Shared iPad, in tvOS 18 and later, and in visionOS 26 and later.

## Topics

### Objects

- [EraseDeviceCommand.Command.ReturnToService](/documentation/devicemanagement/erasedevicecommand/command-data.dictionary/returntoservice-data.dictionary) - The configuration settings for return to service.

