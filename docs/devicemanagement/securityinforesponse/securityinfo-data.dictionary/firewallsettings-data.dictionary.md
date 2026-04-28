# SecurityInfoResponse.SecurityInfo.FirewallSettings

A dictionary that contains the firewall settings.

**Platforms:** macOS 10.12, Device Assignment Services , VPP License Management 

## Properties

### Applications

- **Type:** `[SecurityInfoResponse.SecurityInfo.FirewallSettings.ApplicationsItem]`
- **Required:** No

An array of dictionaries that describes the allowed applications.

### BlockAllIncoming

- **Type:** `boolean`
- **Required:** No

If `true`, the firewall blocks all incoming connections.

### FirewallEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the firewall is on.

### LoggingEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, logging is enabled.

### LoggingOption

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `throttled`, `brief`, `detail`

The type of logging emitted by the firewall.

### StealthMode

- **Type:** `boolean`
- **Required:** No

If true, stealth mode is active for the firewall.

## Topics

### Objects

- [SecurityInfoResponse.SecurityInfo.FirewallSettings.ApplicationsItem](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary/firewallsettings-data.dictionary/applicationsitem) - A dictionary that describes the allowed apps.

