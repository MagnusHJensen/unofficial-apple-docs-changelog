# Dock

The payload that configures the Dock.

**Platforms:** macOS 10.7

## Properties

### AllowDockFixupOverride

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, use the file in `/Library/Preferences/com.apple.dockfixup.plist` when a new user or migrated user logs in. This option has no effect for existing users. Available in macOS 10.12 and later. Only available on the device channel.

### autohide

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Automatically hide and show the Dock.”

### autohide-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, locks “Automatically hide.”

### contents-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, disables changes to the Dock.

### dblclickbehavior

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `minimize`, `maximize`, `none`

The behavior when the window’s title bar is double-clicked.

### dblclickbehavior-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, locks “Double-click a window’s title bar.”

### largesize

- **Type:** `integer`
- **Required:** No

The size of the largest magnification.

### launchanim

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Animate opening applications.”

### launchanim-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, locks “Animate opening applications.”

### magnification

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables magnification.

### magnify-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, locks magnification.

### magsize-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, locks the magnification slider.

### MCXDockSpecialFolders

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AddDockMCXMyApplicationsFolder`, `AddDockMCXDocumentsFolder`, `AddDockMCXSharedFolder`, `AddDockMCXOriginalNetworkHomeFolder`

One or more special folders that may be created at user login time and placed in the Dock.

The “My Applications” item is only used for Simple Finder environments. The “Original Network Home” item is only used for mobile account users.

### mineffect

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `genie`, `scale`

The minimize effect.

### mineffect-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, locks “Minimize windows using.”

### minimize-to-application

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Minimize windows into application icon.”

### minintoapp-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, disables the “Minimize windows into application icon” checkbox.

### orientation

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `bottom`, `left`, `right`

The orientation of the Dock.

### persistent-apps

- **Type:** `[Dock.StaticItem]`
- **Required:** No

An array of items located on the Applications side of the Dock that can be removed from the Dock.

### persistent-others

- **Type:** `[Dock.StaticItem]`
- **Required:** No

An array of items located on the Documents side of the Dock that can be removed from the Dock.

### position-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, locks the position.

### show-process-indicators

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If true, shows the process indicator.

### show-recents

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Show recent items.”

### showindicators-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, locks “Show indicators.”

### showrecents-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, disables “Show recent applications” checkbox.

### size-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, locks the size slider.

### static-apps

- **Type:** `[Dock.StaticItem]`
- **Required:** No

An array of items located on the Applications side of the Dock and cannot be removed from that location.

### static-only

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, uses the `static-apps` and `static-others` dictionaries for the Dock and ignores any items in the `persistent-apps` and `persistent-others` dictionaries. If `false`, the contents are merged with the static items listed first.

### static-others

- **Type:** `[Dock.StaticItem]`
- **Required:** No

An array of items located on the Documents side of the Dock and cannot be removed from that location.

### tilesize

- **Type:** `integer`
- **Required:** No

The tile size. Values must be in the range from 16 to 128.

### windowtabbing

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `manual`, `always`, `fullscreen`

Set the “Prefer tabs when opening documents” to the provided value.

### windowtabbing-immutable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, disables “Prefer tabs when opening documents” checkbox.

## Discussion

Specify `com.apple.dock` as the payload type.

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
            <key>tilesize</key>
            <integer>40</integer>
            <key>size-immutable</key>
            <true/>
            <key>magnification</key>
            <true/>
            <key>magnify-immutable</key>
            <true/>
            <key>largesize</key>
            <integer>100</integer>
            <key>magsize-immutable</key>
            <true/>
            <key>orientation</key>
            <string>left</string>
            <key>position-immutable</key>
            <true/>
            <key>mineffect</key>
            <string>genie</string>
            <key>mineffect-immutable</key>
            <true/>
            <key>windowtabbing</key>
            <string>manual</string>
            <key>windowtabbing-immutable</key>
            <true/>
            <key>dblclickbehavior</key>
            <string>maximize</string>
            <key>dblclickbehavior-immutable</key>
            <true/>
            <key>minimize-to-application</key>
            <true/>
            <key>minintoapp-immutable</key>
            <true/>
            <key>launchanim</key>
            <true/>
            <key>launchanim-immutable</key>
            <true/>
            <key>autohide</key>
            <false/>
            <key>autohide-immutable</key>
            <true/>
            <key>show-process-indicators</key>
            <false/>
            <key>showindicators-immutable</key>
            <true/>
            <key>show-recents</key>
            <false/>
            <key>showrecents-immutable</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.mydockpayload</string>
            <key>PayloadType</key>
            <string>com.apple.dock</string>
            <key>PayloadUUID</key>
            <string>8d443602-52f2-48ff-aaa9-35b16c7c54c9</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Dock</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>c139d3e0-5468-43b0-90bf-5e05b2e8cd6f</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [Dock.StaticItem](/documentation/devicemanagement/dock/staticitem) - Items that are located on the Documents side of the Dock and cannot be removed from that location.

