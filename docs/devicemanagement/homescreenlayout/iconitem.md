# HomeScreenLayout.IconItem

An array of dictionaries that conform to the icon dictionary format.

**Platforms:** iOS 9.3, iPadOS 9.3, tvOS 11.0

## Properties

### BundleID

- **Type:** `string`
- **Required:** No

The bundle identifier of the app. This setting is required if the type is `Application`.

### DisplayName

- **Type:** `string`
- **Required:** No

The human-readable string shown to the user. This setting is valid only if the type is `Folder`.

### Pages

- **Type:** `[[HomeScreenLayout.IconItem]]`
- **Required:** No

An array of arrays of dictionaries, each conforming to the icon dictionary format. This setting is valid only if the type is `Folder`.

### Type

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Application`, `Folder`, `WebClip`

The type of the Dock item.

### URL

- **Type:** `string`
- **Required:** No

The URL of the existing web clip for this item. This setting is required if `type` is `WebClip`. If more than one web clip exists with the same URL, the behavior is undefined.

Specifying a web clip in this payload doesn’t create the web clip. Use the [WebClip](/documentation/devicemanagement/webclip) payload to create a web clip.

