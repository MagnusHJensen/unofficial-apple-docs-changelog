# AudioAccessorySettingsTemporaryPairing_Configuration_UnpairingTimeObject

A dictionary that describes when the device automatically unpairs temporarily paired audio accessories.

**Platforms:** iOS 26.0, iPadOS 26.0

## Properties

### Hour

- **Type:** `integer`
- **Required:** No

The local time hour (24-hour clock) when the device automatically unpairs temporarily paired audio accessories. Required when setting the `Policy` key to `Hour`.

### Policy

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `None`, `Hour`

A string that specifies the device’s unpairing policy.

- `None`: The device doesn’t automatically unpair. Use this only with a return to service device that you erase and reenroll when assigning it from one user to another.
- `Hour`: The device automatically unpairs temporarily paired audio accessories at the local time that the `Hour` key specifies.

