# APN.DefaultsData.ApnsItem

A dictionary that describes an APN configuration.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, Device Assignment Services , VPP License Management 

## Properties

### apn

- **Type:** `string`
- **Required:** Yes

The access point name.

### password

- **Type:** `data`
- **Required:** No

The password for the user. For obfuscation purposes, the system encodes the password. If missing, the device prompts for the password during profile installation.

### proxy

- **Type:** `string`
- **Required:** No

The IP address or URL of the APN proxy.

### proxyPort

- **Type:** `integer`
- **Required:** No

The port number of the APN proxy.

### username

- **Type:** `string`
- **Required:** No

The user name. If missing, the device prompts for it during profile installation.

