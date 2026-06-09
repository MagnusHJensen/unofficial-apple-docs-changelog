# StatusPackageListPackageObject

A dictionary that describes a declarative package.

**Platforms:** macOS 26.0

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system removed the package and only this key and the `identifier` key are present in the status item object.

### declaration-identifier

- **Type:** `string`
- **Required:** No

The identifier of the declaration that controls the package.

### identifier

- **Type:** `string`
- **Required:** Yes

The package’s unique identifier. This is the package identifier value of the package file.

### name

- **Type:** `string`
- **Required:** No

The name of the package.

### reasons

- **Type:** `[StatusPackageListStatusReasonObject]`
- **Required:** No

An array that contains additional details about the package state, including errors.

### state

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `optional`, `queued`, `prompting-for-consent`, `downloading`, `installing`, `installed`, `failed`

The status of the package, which has the following possible values:

- `optional`: The package is optional and the user has to trigger its installation.
- `queued`: Installation of the package has started.
- `prompting-for-consent`: The system is displaying a prompt to the user to proceed with package installation.
- `downloading`: The system is downloading the package.
- `installing`: The system is installing the package.
- `installed`: The package is installed.
- `failed`: The package install failed.

### version

- **Type:** `string`
- **Required:** No

The version of the package. This will be the package version value of the package file.

## Topics

### Objects

- [StatusPackageListStatusReasonObject](/documentation/devicemanagement/statuspackageliststatusreasonobject) - Information about a status error.

