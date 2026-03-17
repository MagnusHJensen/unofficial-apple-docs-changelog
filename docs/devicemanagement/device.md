# Device

A device’s properties and their values.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### asset_tag

- **Type:** `string`
- **Required:** No

The device’s asset tag.

### color

- **Type:** `string`
- **Required:** No

The color of the device.

### description

- **Type:** `string`
- **Required:** No

A description of the device.

### device_assigned_by

- **Type:** `string`
- **Required:** No

The email of the person who assigned the device.

### device_assigned_date

- **Type:** `string`
- **Required:** No

A time stamp in ISO 8601 format that indicates when the device was assigned to the MDM server.

### device_family

- **Type:** `string`
- **Required:** No

The device’s Apple product family: `iPad`, `iPhone`, `iPod`, `Mac`, `AppleTV`, or `Vision`. This key is valid in X-Server-Protocol-Version 2 and later.

### mdm_migration_deadline

- **Type:** `string`
- **Required:** No

A time stamp in ISO 8601 format that indicates the MDM migration deadline. This key is valid with X-Server-Protocol-Version 8 and later.

### model

- **Type:** `string`
- **Required:** No

The model name.

### op_date

- **Type:** `string`
- **Required:** No

A time stamp in ISO 8601 format that indicates when the device was added, updated, or deleted. If the value of `op_type` is added, this is the same as `device_assigned_date`. This field is only applicable with the [Sync the List of Devices](/documentation/devicemanagement/sync-devices) command.

### op_type

- **Type:** `string`
- **Required:** No

Indicates whether the device was added (assigned to the MDM server), modified, or deleted. Contains one of the following strings: `added`, `modified`, or `deleted.This` field is only applicable with the `sync the list of devices` command.

### os

- **Type:** `string`
- **Required:** No

The device’s operating system: `iOS`, `iPadOS`, `OSX`, `tvOS`, or `visionOS`.

This key is valid in X-Server-Protocol-Version 2 and later.

With X-Server-Protocol-Version 7 and later, iPad product os will return `iPadOS`.

### profile_assign_time

- **Type:** `string`
- **Required:** No

A time stamp in ISO 8601 format that indicates when a profile was assigned to the device.

### profile_push_time

- **Type:** `string`
- **Required:** No

A time stamp in ISO 8601 format that indicates when a profile was pushed to the device.

### profile_status

- **Type:** `string`
- **Required:** No

The status of profile installation—either `empty`, `assigned`, `pushed`, or `removed`.

### profile_uuid

- **Type:** `string`
- **Required:** No

The unique ID of the assigned profile.

### serial_number

- **Type:** `string`
- **Required:** No

The device’s serial number.

