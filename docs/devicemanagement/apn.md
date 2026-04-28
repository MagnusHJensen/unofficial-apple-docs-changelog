# APN

The payload that configures access point names.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, Device Assignment Services , VPP License Management 

## Properties

### DefaultsData

- **Type:** `APN.DefaultsData`
- **Required:** Yes

The list of access point names (APNs).

### DefaultsDomainName

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `com.apple.managedCarrier`

The domain name.

## Discussion

Specify `com.apple.apn.managed` as the payload type.

This profile is deprecated. Use the [Cellular](/documentation/devicemanagement/cellular) profile instead.

### Profile availability

## Topics

### Objects

- [APN.DefaultsData](/documentation/devicemanagement/apn/defaultsdata-data.dictionary) - An array of access point name dictionaries.

