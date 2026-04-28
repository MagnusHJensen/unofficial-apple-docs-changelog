# Printing

The payload that configures printers.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### AllowLocalPrinters

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, allows printers that connect directly to a user’s computer.

### DefaultPrinter

- **Type:** `Printing.DefaultPrinter`
- **Required:** No

The default printer for the user.

### FooterFontName

- **Type:** `string`
- **Required:** No

The footer font name.

### FooterFontSize

- **Type:** `string`
- **Required:** No

The footer font size.

### PrintFooter

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, prints the page footer (including the user name and date).

### PrintMACAddress

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, includes the MAC address.

### RequireAdminToAddPrinters

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, requires an administrator password to add printers.

### RequireAdminToPrintLocally

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, requires an administrator password to print locally.

### ShowOnlyManagedPrinters

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, shows only managed printers.

### UserPrinterList

- **Type:** `Printing.UserPrinterList`
- **Required:** No

The printers available to a user.

## Discussion

Specify `com.apple.mcxprinting` as the payload type.

Removing this profile from a device doesn’t automatically remove printers from the device.

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

            <key>RequireAdminToAddPrinters</key>
            <false/>
            <key>AllowLocalPrinters</key>
            <true/>
            <key>RequireAdminToPrintLocally</key>
            <true/>
            <key>ShowOnlyManagedPrinters</key>
            <true/>
            <key>PrintFooter</key>
            <true/>
            <key>PrintMACAddress</key>
            <true/>
            <key>FooterFontSize</key>
            <string>7</string>
            <key>FooterFontName</key>
            <string>Helvetica</string>
            <key>DefaultPrinter</key>
            <dict>
                <key>DeviceURI</key>
                <string>ipp://printer.example.com/</string>
                <key>DisplayName</key>
                <string>printer.example.com</string>
            </dict>
            <key>UserPrinterList</key>
            <dict>
                <key>printer_example_com</key>
                <dict>
                    <key>DeviceURI</key>
                    <string>ipp://printer.example.com/</string>
                    <key>DisplayName</key>
                    <string>printer.example.com</string>
                    <key>Location</key>
                    <string>My Office</string>
                    <key>Model</key>
                    <string>PrinterModel1</string>
                    <key>PrinterLocked</key>
                    <true/>
                    <key>PPDURL</key>
                    <string>file://localhost/System/Library/Frameworks/ApplicationServices.framework/Versions/A/Frameworks/PrintCore.framework/Resources/Generic.ppd</string>
                </dict>
            </dict>
            <key>PayloadIdentifier</key>
            <string>com.example.myprinterpayload</string>
            <key>PayloadType</key>
            <string>com.apple.mcxprinting</string>
            <key>PayloadUUID</key>
            <string>8242d870-95c0-0135-0b44-0c4de9ce4c04</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Printing</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>ab59f143-1478-419a-885e-7994fb13c9c3</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [Printing.DefaultPrinter](/documentation/devicemanagement/printing/defaultprinter-data.dictionary) - A default printer for the user.
- [Printing.UserPrinterList](/documentation/devicemanagement/printing/userprinterlist-data.dictionary) - A list of printer dictionaries.

