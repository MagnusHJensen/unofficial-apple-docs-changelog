# AirPrint.AirPrintItem

A dictionary of AirPrint printer details.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.10, visionOS 2.0, Device Assignment Services , VPP License Management 

## Properties

### ForceTLS

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, AirPrint connections are secured by Transport Layer Security (TLS). Available only in iOS 11 and later.

### IPAddress

- **Type:** `string`
- **Required:** Yes

The IP address or hostname of the AirPrint destination.

### Port

- **Type:** `integer`
- **Required:** No

The listening port of the AirPrint destination. Available only in iOS 11 and later.

### ResourcePath

- **Type:** `string`
- **Required:** Yes

The resource path associated with the printer. This path corresponds to the `rp` parameter of the `_ipps.tcp` Bonjour record. For example:

- `printers/Canon_MG5300_series`
- `printers/Xerox_Phaser_7600`
- `ipp/print`
- `Epson_IPP_Printer`

