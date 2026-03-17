# StatusPackageList

The client’s declarative packages.

**Platforms:** macOS 26.0

## Properties

### package.list

- **Type:** `[StatusPackageListPackageObject]`
- **Required:** Yes

An array of dictionaries that describe the device’s declarative packages.

## Discussion

### Status item availability

### Reason Codes

- `Error.DownloadFailed`: The package download failed.
- `Error.InstallFailed`: The package install failed.

## Topics

### Objects

- [StatusPackageListPackageObject](/documentation/devicemanagement/statuspackagelistpackageobject) - A dictionary that describes a declarative package.

