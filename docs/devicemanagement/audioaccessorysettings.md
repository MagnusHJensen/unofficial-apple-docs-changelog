# AudioAccessorySettings

The declaration to configure audio accessory settings.

**Platforms:** iOS 26.0, iPadOS 26.0

## Discussion

Specify `com.apple.configuration.audio-accessory.settings` as the declaration type.

Setting `TemporaryPairing` to `false` disables only the temporary pairing feature, without impacting any other use of audio accessories, so users can still:

- Pair and use audio accessories - the device records the pairing and synchronizes it to their iCloud account.
- Use the audio accessory AirPods Sharing feature.

### Configuration availability

### Configuration example

This configuration enables temporary pairing and sets an unpairing time of 6 pm.

```json
{
    "Type": "com.apple.configuration.audio-accessory.settings",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "TemporaryPairing": {
            "Configuration": {
                "UnpairingTime": {
                    "Policy": "Hour",
                    "Hour": 18
                }
            }
        }
    }
}
```

## Topics

### Objects

- [AudioAccessorySettingsTemporaryPairingObject](/documentation/devicemanagement/audioaccessorysettingstemporarypairingobject) - A dictionary that describes audio accessory temporary pairing behavior. The device enables temporary pairing when this key is present and the `Disabled` key isn’t `false`. The device doesn’t synchronize pairing information with iCloud when temporary pairing is active.

