# Package

The declaration to configure a package.

**Platforms:** macOS 26.0

## Properties

### InstallBehavior

- **Type:** `PackageInstallBehaviorObject`
- **Required:** No

A dictionary that describes how and when to install the package.

### ManifestURL

- **Type:** `string`
- **Required:** Yes

The URL of the manifest document for the package that the device downloads. The manifest is returned as a [ManifestURL](/documentation/devicemanagement/manifesturl) property list. The `url` property of the manifest must point to the package (.pkg) file to install.

### UninstallBehavior

- **Type:** `PackageUninstallBehaviorObject`
- **Required:** No

A dictionary that describes how to uninstall the package.

Available: macOS 27+

## Discussion

Specify `com.apple.configuration.package` as the declaration type.

This declaration installs a package on a device. Packages can contain apps, fonts, documents, and other items. Apps that a package installs aren’t automatically managed; you can manage them using the [AppManaged](/documentation/devicemanagement/appmanaged) declaration.

### Configuration availability

### Configuration example

This configuration installs a required package.

```json
{
    "Type": "com.apple.configuration.package",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "ManifestURL": "https://example.com/files/packages/TestPackage.plist",
        "InstallBehavior": {
            "Install": "Required"
        }
    }
}
```

## Topics

### Objects

- [PackageInstallBehaviorObject](/documentation/devicemanagement/packageinstallbehaviorobject) - A dictionary that describes how and when to install the package.
- [PackageUninstallBehaviorObject](/documentation/devicemanagement/packageuninstallbehaviorobject) - A dictionary that describes how to uninstall the package.

