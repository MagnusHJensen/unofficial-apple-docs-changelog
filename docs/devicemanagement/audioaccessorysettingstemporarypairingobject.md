# AudioAccessorySettingsTemporaryPairingObject

A dictionary that describes audio accessory temporary pairing behavior. The device enables temporary pairing when this key is present and the `Disabled` key isn’t `false`. The device doesn’t synchronize pairing information with iCloud when temporary pairing is active.

**Platforms:** iOS 26.0, iPadOS 26.0, Mac Catalyst 26.0, Device Assignment Services , VPP License Management 

## Properties

### Configuration

- **Type:** `AudioAccessorySettingsTemporaryPairing_ConfigurationObject`
- **Required:** No

A dictionary providing configuration for temporary pairing. Required if `Disabled` isn’t present or is `false`.

### Disabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, temporary pairing of audio accessories is disabled.

## Topics

### Objects

- [AudioAccessorySettingsTemporaryPairing_ConfigurationObject](/documentation/devicemanagement/audioaccessorysettingstemporarypairing_configurationobject) - A dictionary providing configuration for temporary pairing. Required if `Disabled` isn’t present or is `false`.

