# StatusDeviceBatteryHealth

The device’s battery health.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.4

## Properties

### device.power.battery-health

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `non-genuine`, `normal`, `service-recommended`, `unknown`, `unsupported`

The battery health status, which has the following values:

- `non-genuine`: The battery isn’t a genuine Apple battery.
- `normal`: The battery is operating normally.
- `service-recommended`: The system recommends battery service.
- `unknown`: The system couldn’t determine battery health information.
- `unsupported`: The device doesn’t support battery health reporting.

Available in iOS 17 and later on iPhone, iPadOS 18.4 and later on supported iPad models, and macOS 14.4 and later on a Mac with Apple silicon.

## Discussion

For more information about battery health, see the following support articles:

- [iPhone devices](https://support.apple.com/101575)
- [iPad devices](https://support.apple.com/117759)
- [macOS devices](https://support.apple.com/108376)

### Status item availability

