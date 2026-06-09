# StatusSoftwareUpdateInstallState

The status item that reports the device’s software update install state.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, tvOS 18.4, visionOS 26.0

## Properties

### softwareupdate.install-state

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `none`, `downloading`, `prepared`, `installing`, `failed`

The software update install status, which has the following values:

- `none`: There’s no software update pending, and any previous software update succeeded.
- `downloading`: The system is downloading data for a software update.
- `prepared`: The system prepared the software update and it’s ready for installation.
- `installing`: The system is installing the software update.
- `failed`: The software update failed.

## Discussion

### Status item availability

### Status item example

```json
{
    "softwareupdate": {
        "install-state": "none"
    }
}
```

