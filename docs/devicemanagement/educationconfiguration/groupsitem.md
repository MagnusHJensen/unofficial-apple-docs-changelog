# EducationConfiguration.GroupsItem

An array of dictionaries defining groups.

**Platforms:** iOS 9.3, iPadOS 9.3, Mac Catalyst 9.3, macOS 10.14, Device Assignment Services , VPP License Management 

## Properties

### BeaconID

- **Type:** `integer`
- **Required:** Yes

An unsigned 16 bit integer specifying this group’s unique beacon ID.

### ConfigurationSource

- **Type:** `string`
- **Required:** No

The source that provided this group, such as SIS, or MDM.

### Description

- **Type:** `string`
- **Required:** No

The description of the group.

### DeviceGroupIdentifiers

- **Type:** `[string]`
- **Required:** No

The identifiers that refer to entries in the `DeviceGroups` array to which the instructor can assign users from this class.

Has no effect on the configuration of the Shared iPad login screen.

### ImageURL

- **Type:** `string`
- **Required:** No

Deprecated in iOS 9.3.1 and later. The URL of an image for the group.

### LeaderIdentifiers

- **Type:** `[string]`
- **Required:** No

The user identifiers that are leaders of this group.

### MemberIdentifiers

- **Type:** `[string]`
- **Required:** Yes

The entries in the Users array that are members of the group.

### Name

- **Type:** `string`
- **Required:** Yes

The display name of the group.

