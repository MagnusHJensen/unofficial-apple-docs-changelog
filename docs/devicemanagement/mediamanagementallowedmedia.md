# MediaManagementAllowedMedia

The payload that configures media management.

**Platforms:** macOS 10.7

## Properties

### logout-eject

- **Type:** `MediaManagementAllowedMedia.Logout-eject`
- **Required:** No

The media type dictionary that defines volumes to eject when the user logs out.

### mount-controls

- **Type:** `MediaManagementAllowedMedia.Mount-controls`
- **Required:** No

The media type dictionary that controls volume mounting.

### unmount-controls

- **Type:** `MediaManagementAllowedMedia.Unmount-controls`
- **Required:** No

The media type dictionary that controls volume unmounting.

## Discussion

Specify `com.apple.systemuiserver` as the payload type.

This payload is deprecated as of macOS 11.

### Profile availability

## Topics

### Objects

- [MediaManagementAllowedMedia.Logout-eject](/documentation/devicemanagement/mediamanagementallowedmedia/logout-eject-data.dictionary) - A dictionary of volumes to eject when the user logs out.
- [MediaManagementAllowedMedia.Mount-controls](/documentation/devicemanagement/mediamanagementallowedmedia/mount-controls-data.dictionary) - A dictionary of volumes to control volume mounting.
- [MediaManagementAllowedMedia.Unmount-controls](/documentation/devicemanagement/mediamanagementallowedmedia/unmount-controls-data.dictionary) - A dictionary to control volume unmounting.

