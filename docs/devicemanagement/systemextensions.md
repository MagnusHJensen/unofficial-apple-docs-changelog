# SystemExtensions

The payload that configures system extensions.

**Platforms:** macOS 10.15

## Properties

### AllowedSystemExtensions

- **Type:** `SystemExtensions.AllowedSystemExtensions`
- **Required:** No

A dictionary of approved system extensions on the computer. The dictionary maps the team identifiers (keys) to arrays of bundle identifiers, where the bundle identifier defines the system extension to install.

To avoid requiring an administrator to authorize the operation, you can activate system extensions that this key specifies using [activationRequest(forExtensionWithIdentifier:queue:)](/documentation/SystemExtensions/OSSystemExtensionRequest/activationRequest(forExtensionWithIdentifier:queue:)).

It’s an error for the same team identifier to appear in both the `AllowedTeamIdentifiers` array and as a key in this dictionary.

### AllowedSystemExtensionTypes

- **Type:** `SystemExtensions.AllowedSystemExtensionTypes`
- **Required:** No

A dictionary that maps a team identifier to an array of strings, where each string is a type of system extension that you can install for that team identifier. The allowed extension types are `DriverExtension`, `NetworkExtension`, and `EndpointSecurityExtension`.

If there’s no entry for a specified team identifier in the dictionary, the system allows all extension types.

### AllowedTeamIdentifiers

- **Type:** `[string]`
- **Required:** No

An array of team identifiers that defines valid, signed system extensions that are allowable to load. Approved system extensions are those signed with any of the specified team identifiers.

To avoid requiring an administrator to authorize the operation, you can activate system extensions that this key specifies using [activationRequest(forExtensionWithIdentifier:queue:)](/documentation/SystemExtensions/OSSystemExtensionRequest/activationRequest(forExtensionWithIdentifier:queue:)).

It’s an error for the same team identifier to appear in both this array and as a key in the `AllowedSystemExtensions` dictionary.

### AllowUserOverrides

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, restricts users from approving additional system extensions that configuration profiles don’t explicitly allow.

### NonRemovableFromUISystemExtensions

- **Type:** `SystemExtensions.NonRemovableFromUISystemExtensions`
- **Required:** No

A dictionary of system extensions on the computer. The dictionary maps the team identifiers (keys) to arrays of bundle identifiers, where the bundle identifier defines the system extension which can’t be disabled or uninstalled from System Settings or Finder. The set of system extensions between `RemovableSystemExtensions` and `NonRemovableFromUISystemExtensions` can to overlap.

### NonRemovableSystemExtensions

- **Type:** `SystemExtensions.NonRemovableSystemExtensions`
- **Required:** No

A dictionary of system extensions on the computer. The dictionary maps the team identifiers (keys) to arrays of bundle identifiers, where the bundle identifier defines the system extension which can’t be disabled or uninstalled when SIP is enabled. It’s an error for the same mapping to appear in the dictionary values corresponding to `RemovableSystemExtensions` and `NonRemovableSystemExtensions` keys.

### RemovableSystemExtensions

- **Type:** `SystemExtensions.RemovableSystemExtensions`
- **Required:** No

A dictionary of system extensions that are allowed to remove themselves from the machine. The dictionary maps team identifiers (keys) to arrays of bundle identifiers, where the bundle identifier defines the system extension. An application using the `OSSystemExtensionDeactivationRequest` API can deactivate the specified system extensions without requiring an administrator to authorize the operation.

Available in macOS 12 and later.

## Discussion

Specify `com.apple.system-extension-policy` as the payload type.

When there are multiple installed profiles, the keys combine as follows:

- `AllowUserOverrides` is `false` if any profile sets it to `false`.
- All the other values combine together.

Beginning in macOS 11.3, installing or removing this payload can change the state of system extensions on the computer. If a containing application activates a system extension, and the system extension is in a pending state, installing a payload that allows the extension completes the activation process. If a system extension is active, removing a payload that allows the extension deactivates that extension.

### Profile availability

### Profile example

```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>PayloadContent</key>
    <array>
        <dict>
            <key>AllowedExtensions</key>
            <array>
                <string>com.apple.share.System.send</string>
            </array>
            <key>PayloadIdentifier</key>
            <string>com.example.mysystemextensionpayload</string>
            <key>PayloadType</key>
            <string>com.apple.nsextension</string>
            <key>PayloadUUID</key>
            <string>3b378001-8cff-4bde-a6af-843fdb02f36d</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>System Extensions</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>2e880ad8-e49e-47af-b416-7e6c3fba9df2</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [SystemExtensions.AllowedSystemExtensionTypes](/documentation/devicemanagement/systemextensions/allowedsystemextensiontypes-data.dictionary) - A dictionary that maps team identifiers to system extensions.
- [SystemExtensions.AllowedSystemExtensions](/documentation/devicemanagement/systemextensions/allowedsystemextensions-data.dictionary) - A dictionary that maps team identifiers to bundle identifiers that are allowed.
- [SystemExtensions.NonRemovableFromUISystemExtensions](/documentation/devicemanagement/systemextensions/nonremovablefromuisystemextensions-data.dictionary) - A dictionary that maps team identifiers to bundle identifiers of extensions that are non-removable.
- [SystemExtensions.NonRemovableSystemExtensions](/documentation/devicemanagement/systemextensions/nonremovablesystemextensions-data.dictionary) - A dictionary that maps team identifiers to bundle identifiers of extensions that are non-removable.
- [SystemExtensions.RemovableSystemExtensions](/documentation/devicemanagement/systemextensions/removablesystemextensions-data.dictionary) - A dictionary that maps team identifiers to bundle identifiers of extensions that are removable.

