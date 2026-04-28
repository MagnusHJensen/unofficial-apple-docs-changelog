# RemoveMediaCommand.Command

The command to remove a previously installed book from a device.

**Platforms:** iOS 8.0, iPadOS 8.0, Mac Catalyst 8.0, Device Assignment Services , VPP License Management 

## Properties

### iTunesStoreID

- **Type:** `string`
- **Required:** No

The book’s iTunes Store identifier.

### MediaType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Book`

The media type, which can only be `Book`.

### PersistentID

- **Type:** `string`
- **Required:** No

The book’s persistent identifier in reverse-DNS form; for example, `com.acme.manuals.training`.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `RemoveMedia`

The request type for this command.

