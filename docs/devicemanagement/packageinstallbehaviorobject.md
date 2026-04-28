# PackageInstallBehaviorObject

Specifies the install behavior of the package.

**Platforms:** macOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### Install

- **Type:** `string`
- **Required:** No
- **Default:** `Optional`
- **Allowed Values:** `Optional`, `Required`

A string that specifies when the system installs the package:

- `Optional`: The user can install the package after the system activates the configuration.
- `Required`: The system installs the package after it activates the configuration.

