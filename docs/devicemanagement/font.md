# Font

The payload that configures fonts.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.9, visionOS 2.0

## Properties

### Font

- **Type:** `data`
- **Required:** Yes

The contents of the font file.

### Name

- **Type:** `string`
- **Required:** No
- **Default:** ``

The user-visible name for the font. The device replaces this field with the actual name of the font after installation. Each payload must contain exactly one font file in trueType (.ttf) or OpenType (.otf) format. The device doesn’t support collection formats (.ttc or .otc).

The device identifies fonts by their embedded PostScript names. The device considers two fonts with the same PostScript name to be the same font, even if their contents differ. The device doesn’t support installing two different fonts with the same PostScript name, and the resulting behavior is undefined.

## Discussion

Specify `com.apple.font` as the payload type.

In iPadOS 18 and later, the font profile is available on the user channel for Shared iPad.

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
                <key>Font</key>
                <data>YmFzZTY0ZW5jb2RlZGRhdGEK</data>
                <key>Name</key>
                <string>Font.ttf</string>
                <key>PayloadIdentifier</key>
                <string>com.example.myfontpayload</string>
                <key>PayloadType</key>
                <string>com.apple.font</string>
                <key>PayloadUUID</key>
                <string>99659c06-bbbf-45aa-9674-06b378ec2cd5</string>
                <key>PayloadVersion</key>
                <integer>1</integer>
            </dict>
        </array>
        <key>PayloadDisplayName</key>
        <string>Fonts</string>
        <key>PayloadIdentifier</key>
        <string>com.example.myprofile</string>
        <key>PayloadType</key>
        <string>Configuration</string>
        <key>PayloadUUID</key>
        <string>9800ea91-30c4-4212-8033-d21cad4fe99a</string>
        <key>PayloadVersion</key>
        <integer>1</integer>
    </dict>
</plist>
```

