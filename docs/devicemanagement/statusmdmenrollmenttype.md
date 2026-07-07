# StatusMDMEnrollmentType

The status item that reports the device’s management enrollment type.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta), watchOS 27.0 (Beta)

## Properties

### mdm.enrollment-type

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `none`, `supervised`, `device`, `user`

The device management enrollment type that indicates how the device is enrolled, which has the following possible values:

- `none`: Device isn’t enrolled
- `supervised`: Device is supervised
- `device`: Device enrollment
- `user`: User enrollment

## Discussion

### Status item availability

### Status item example

```json
{
    "mdm": {
        "enrollment-type": "supervised"
    }
}
```

