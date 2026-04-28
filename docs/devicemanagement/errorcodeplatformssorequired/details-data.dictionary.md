# ErrorCodePlatformSSORequired.Details

A dictionary that contains additional data about the error code.

**Platforms:** macOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### AuthURL

- **Type:** `string`
- **Required:** Yes

The URL the device uses to create an `ASWebAuthenticationSession` to trigger Platform SSO authentication, once the profile and app are installed.

### Package

- **Type:** `ErrorCodePlatformSSORequired.Details.Package`
- **Required:** Yes

A dictionary that specifies the package that the device uses to install an app with the SSO app extension used for Platform SSO.

### ProfileURL

- **Type:** `string`
- **Required:** Yes

The URL of the profile containing an [ExtensibleSingleSignOn](/documentation/devicemanagement/extensiblesinglesignon) profile payload that the device uses to configure the SSO extension for Platform SSO.

## Topics

### Objects

- [ErrorCodePlatformSSORequired.Details.Package](/documentation/devicemanagement/errorcodeplatformssorequired/details-data.dictionary/package-data.dictionary) - A dictionary that specifies the package that the device uses to install an app with the SSO app extension used for Platform SSO.

