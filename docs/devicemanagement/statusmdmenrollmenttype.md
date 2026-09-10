# StatusMDMEnrollmentType

The status item that reports the device’s management enrollment type.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0, watchOS 27.0

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

