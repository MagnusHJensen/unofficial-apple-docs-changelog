# Installing packages

Use declarative package management to install and remove packages in macOS.

## Overview

Device management services can install and monitor packages using declarative management on macOS devices. The services apply a declarative management configuration to the device and subscribe to a declarative management status item. Devices support packages with any type of content. If a package installs an app, a service can manage the app using the [AppManaged](/documentation/devicemanagement/appmanaged) configuration with an `AppComposedIdentifier` key.

Devices can install packages only on supervised devices in the system scope (the device management service device channel).

The [Package](/documentation/devicemanagement/package) configuration type is `com.apple.configuration.package`. Use this configuration to install a package.

## Install packages

To install packages, the device downloads a manifest document from the URL in the configuration’s `ManifestURL` key. The manifest needs to match the [ManifestURL](/documentation/devicemanagement/manifesturl) format and reference a single package file. It also needs to contain a `bundle-version` key for the package, which the device uses to detect package updates when updating configurations.

The device can install a package immediately after it applies a configuration, or it can wait for the user to trigger the install. The `Install` key in the [PackageInstallBehaviorObject](/documentation/devicemanagement/packageinstallbehaviorobject) object controls this behavior:

A ** displays required and optional packages to the user, provides details about packages, and allows the user to trigger installs of optional packages. For more information, see [Displaying managed apps and packages](/documentation/devicemanagement/displaying-managed-apps-and-packages).

## Update packages

The device doesn’t automatically update packages. To update a package, the device management service needs to:

- Change the `bundle-version` key in the manifest document.
- Change the package configuration’s `ServerToken` key.

The device checks the manifest during configuration updates, and updates packages if needed.

## Remove packages

The device removes a package when it deactivates or removes the configuration for a package, or when it unenrolls. However, the device doesn’t track installed package items, so it doesn’t remove those items. Instead, it only removes the system metadata for the package.

## Get the status of packages

The device reports the status of each declarative managed package in the [StatusPackageList](/documentation/devicemanagement/statuspackagelist) status item, which device management servers can subscribe to. When the package status changes, the device reports incremental updates to the [StatusPackageList](/documentation/devicemanagement/statuspackagelist) status item elements.

## Process the status item

The [StatusPackageList](/documentation/devicemanagement/statuspackagelist) status item type is `package.list`, which is the device’s list of declarative packages.

The status item’s value is an array of objects, where each object represents a declarative managed package. The device uses normal status-item array value reporting behavior to report changes incrementally to the device management server; for more information, see [StatusReport](/documentation/devicemanagement/statusreport). The status includes required and optional packages, whether installed or not.

The `identifier` key provides the unique identifier of the package for incremental reporting in the overall status item (see [StatusReport](/documentation/devicemanagement/statusreport)). The `declaration-identifier` key contains the `Identifier` of the package configuration. Additional properties show the package’s name and version identifiers.

The `state` key indicates the package’s management state, for example, `queued`, `downloading`, `installing`, and `installed`. If an error occurs during installation, the device reports the `failed` state and includes the `reasons` key to provide more details.

You can use this status item in declarative device management activation predicates using the same format for expressions that [Processing status for managed apps](/documentation/devicemanagement/processing-status-for-managed-apps) describes for managed app status.

## Handle errors

The device reports errors to the device management service in two ways:

- It reports an error in a configuration using the declarative management [StatusManagementDeclarations](/documentation/devicemanagement/statusmanagementdeclarations) status item. The status item shows the configuration as active but invalid, with a reason code. Examples include syntax errors or missing required keys.
- It reports an error during package installation using the declarative management [StatusPackageList](/documentation/devicemanagement/statuspackagelist) status item, described above.

## Build a package

The device only installs signed packages and only when it can verify the signature. The package signature needs to use an appropriate certificate that the device can verify, such as a TLS certificate with signing usage. Only sign the package; you don’t need to sign the app because Gatekeeper doesn’t check apps that MDM installs.

Use the following command-line invocation to build your own signed package for an app:

```bash
$ sudo pkgbuild --component <APP> --install-location /Applications --sign <CERT> /tmp/MyPackage.pkg
```

Set the `<APP>` argument to the path of the app to install, and the `<CERT>` argument to the signing certificate identifier. Set the install location to `/Applications` because the device only manages apps in that location.

In addition to the `.pkg` file, create a [ManifestURL](/documentation/devicemanagement/package/manifesturl) document that the device management service uses to install the package, as the following example shows:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>items</key>
  <array>
    <dict>
      <key>assets</key>
      <array>
        <dict>
          <key>kind</key>
          <string>software-package</string>
          <key>url</key>
          <string>https://mdm.example.com/packages/MyPackage.pkg</string>
          <key>sha256</key>
          <string>49f6554726ae98521b02d89a86f2a7eea5611295fa2f67bf8bc44f679c121a2d</string>
        </dict>
      </array>
      <key>metadata</key>
      <dict>
        <key>bundle-identifier</key>
        <string>com.test.MyPackage.pkg</string>
        <key>bundle-version</key>
        <string>1</string>
        <key>kind</key>
        <string>software</string>
        <key>title</key>
        <string>MyApp.app</string>
      </dict>
    </dict>
  </array>
</dict>
</plist>
```

The [ManifestURL](/documentation/devicemanagement/package/manifesturl) needs to adhere to the following requirements:

- Set the `items/assets/kind` key to `software-package`.
- Set the `items/assets/url` key to the URL the device uses to download the package file.
- Set the `items/assets/sha256` key to the SHA-256 hash value of the package file. For large package files, you can instead use the `sha256s` key to provide an array of hashes for each “chunk” of the package file, with the chunk size specified by the `sha256-size` key.
- Set the `items/metadata/bundle-identifier` key to a unique value for the package. The device uses this value to determine whether a package is an update to an existing package, or a new package.
- Set the `items/metadata/bundle-version` key to a version for the package. The device uses this value to determine whether a package needs to be updated.
- Set the `items/metadata/kind` key to `software`.
- Set the `items/metadata/title` key to the title of the package.

To determine the SHA-256 hash of the package file, use the following command-line invocation:

```bash
$ shasum -a 256 /tmp/MyPackage.pkg
```

You can build more complex packages — for example, ones that contain multiple apps or other types of files — using the `pkgbuild` and `productbuild` command-line tools.

