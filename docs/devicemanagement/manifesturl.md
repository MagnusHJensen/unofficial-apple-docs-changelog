# ManifestURL

The URL to the app manifest.

**Platforms:** iOS 7.0, iPadOS 7.0, macOS 10.9, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Discussion

Use SHA-256 hashes instead of MD5 because SHA-256 has stronger security. If both SHA-256 and MD5 hash properties are present, the device uses only the SHA-256 hashes to verify the manifest data.

### Example Manifest

```plist
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
                    <string>display-image</string>
                    <key>url</key>
                    <string>https://data.example.com/sampleapp.png</string>
                </dict>
                <dict>
                    <key>kind</key>
                    <string>software-package</string>
                    <key>url</key>
                    <string>https://data.example.com/sampleapp.ipa</string>
                </dict>
            </array>
            <key>metadata</key>
            <dict>
                <key>bundle-identifier</key>
                <string>com.example.sampleapp</string>
                <key>bundle-version</key>
                <string>1.2</string>
                <key>developer-name</key>
                <string>Example Inc.</string>
                <key>kind</key>
                <string>software</string>
                <key>title</key>
                <string>The Sample App</string>
            </dict>
        </dict>
    </array>
</dict>
</plist>
```

## Topics

### Objects

- [ManifestURL.ItemsItem](/documentation/devicemanagement/manifesturl/itemsitem) - An array of dictionaries representing what the manifest installs.

