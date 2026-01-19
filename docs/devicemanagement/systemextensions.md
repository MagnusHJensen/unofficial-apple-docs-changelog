# SystemExtensions

The payload that configures system extensions.

**Platforms:** macOS 10.15

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

