# ManageVppLicensesByAdamIdResponse

The response from managing licenses.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### adamIdStr

- **Type:** `string`
- **Required:** No

The unique identifier for a product in the iTunes Store.

### associations

- **Type:** `[VppAssociation]`
- **Required:** No

An array of dictionaries representing successful and failed associations. If an association succeeds, its dictionary contains the license and either a client-user ID or the serial number of the device associated with the license. If an association fails, its dictionary contains the error message and number, and either the client-user ID or the serial number of the device that couldn’t be associated with the license.

### clientContext

- **Type:** `string`
- **Required:** No

The value currently associated with the provided sToken. This field is only included in the response when a value has been set via the [Client Configuration](/documentation/devicemanagement/client-configuration) endpoint.

See [Protecting Your VPP Account](/documentation/devicemanagement/protecting-your-vpp-account) for more information.

### disassociations

- **Type:** `[VppAssociation]`
- **Required:** No

An array of dictionaries representing successful and failed disassociations. If a disassociation succeeds, its dictionary contains a license and either a client-user ID or the serial number of the device disassociated. The license may not be the license ID that was just disassociated; rather the license ID is any currently available license at the time of disassociation. If the disassociation fails, its dictionary contains the error message and number, and either the client-user ID or the serial number of the device that couldn’t be disassociated from the license.

### errorMessage

- **Type:** `string`
- **Required:** No

The human-readable explanation of the error.

### errorNumber

- **Type:** `int32`
- **Required:** No

The numeric code of the error.

### expirationMillis

- **Type:** `int64`
- **Required:** No

The UNIX epoch timestamp, in milliseconds, when the account’s sToken or password expires (whichever is earlier).

### isIrrevocable

- **Type:** `boolean`
- **Required:** No

If `true`, licenses for the specified product can’t be revoked and reassigned.

### location

- **Type:** `VppLocation`
- **Required:** No

The location associated with the provided sToken. This field is only returned when a location token is used with an Apple School Manager account.

### pricingParam

- **Type:** `string`
- **Required:** No

The quality of a product in the iTunes Store. Possible values are:

STDQ: Standard quality

PLUS: High quality

### productTypeId

- **Type:** `int32`
- **Required:** No

The type of a product. Possible values are:

7 = macOS software

8 = iOS or macOS app from the App Store

10 = Book

### productTypeName

- **Type:** `string`
- **Required:** No

The name of the product type.

### status

- **Type:** `int32`
- **Required:** No

The status code for the response. Possible values are:

`0` = Success. All requested associations and disassociations have succeeded.

-1 = Failure. All requested associations and disassociations have failed. An errorNumber, errorCode, and errorMessage is returned for each failed association and disassociation. See [VppAssociation](/documentation/devicemanagement/vppassociation) and  [Interpreting Error Codes](/documentation/devicemanagement/interpreting-error-codes) for more information.

`-3` = The licenses can’t be changed as requested. Some of the requested associations and disassociations have succeeded and others have failed. An errorNumber, errorCode, and errorMessage is returned for each failed association and disassociation. See [VppAssociation](/documentation/devicemanagement/vppassociation) and [Interpreting Error Codes](/documentation/devicemanagement/interpreting-error-codes) for more information.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates. In this way, you can avoid double-counting records when duplicate tokens are uploaded by different content managers.

