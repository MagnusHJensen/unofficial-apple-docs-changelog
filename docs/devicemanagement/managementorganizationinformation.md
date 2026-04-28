# ManagementOrganizationInformation

The declaration to configure the managing organization’s contact information.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Email

- **Type:** `string`
- **Required:** No

The email address of the contact person for the organization.

### Name

- **Type:** `string`
- **Required:** Yes

The name of the organization.

### Proof

- **Type:** `ManagementOrganizationInformationProofObject`
- **Required:** No

The additional properties that verify the identity and authenticity of the organization.

### URL

- **Type:** `string`
- **Required:** No

The website of the organization to contact for support.

## Discussion

Specify `com.apple.management.organization-info` as the declaration type.

### Management declaration example

```json
{
    "Type": "com.apple.management.organization-info",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Name": "Example, Inc.",
        "Email": "admin@example.com",
        "URL": "https://site.example.com/support"
    }
}
```

## Topics

### Objects

- [ManagementOrganizationInformationProofObject](/documentation/devicemanagement/managementorganizationinformationproofobject) - The additional properties that verify the identity and authenticity of the organization.

