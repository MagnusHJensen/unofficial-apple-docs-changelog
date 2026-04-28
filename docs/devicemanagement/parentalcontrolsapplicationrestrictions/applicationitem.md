# ParentalControlsApplicationRestrictions.ApplicationItem

A dictionary defining an app for parental control.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### appID

- **Type:** `data`
- **Required:** Yes

The identifier of the app. Obtain this value from the Security framework using [SecCodeCopyDesignatedRequirement(_:_:_:)](/documentation/Security/SecCodeCopyDesignatedRequirement(_:_:_:)).

### bundleID

- **Type:** `string`
- **Required:** Yes

The bundle ID of the app.

### detachedSignature

- **Type:** `data`
- **Required:** No

The signature for an unsigned binary.

### disabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, this app isn’t added to the allow list.

### displayName

- **Type:** `string`
- **Required:** No

The name used for display purposes.

### subApps

- **Type:** `[ParentalControlsApplicationRestrictions.ApplicationItem]`
- **Required:** No

An array of nested helper applications.

