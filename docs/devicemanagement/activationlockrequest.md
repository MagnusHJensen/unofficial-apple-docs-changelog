# ActivationLockRequest

Request enabling activation lock for a device.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### device

- **Type:** `string`
- **Required:** No

Serial number of the device (required).

### escrow_key

- **Type:** `string`
- **Required:** No

Escrow key (optional). If the escrow key is not provided, the device will be locked to the person who created the MDM server in the portal. For information about creating an escrow key see [Creating and Using Bypass Codes](/documentation/devicemanagement/creating-and-using-bypass-codes).

### lost_message

- **Type:** `string`
- **Required:** No

Lost message to be displayed on the device (optional).

