# MathSettings

The declaration to configure the math and calculator apps.

**Platforms:** iOS 18.0, iPadOS 18.0, macOS 15.0

## Discussion

Specify `com.apple.configuration.math.settings` as the declaration type.

### Configuration availability

### Configuration example

This configuration prevents the use of scientific and programmer modes in calculator app.

```json
{
    "Type": "com.apple.configuration.math.settings",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Calculator": {
            "ScientificMode": {
                "Enabled": false
            },
            "ProgrammerMode": {
                "Enabled": false
            }
        }
    }
}
```

## Topics

### Objects

- [MathSettingsCalculatorObject](/documentation/devicemanagement/mathsettingscalculatorobject) - The declaration to configure the calculator app.
- [MathSettingsSystemBehaviorObject](/documentation/devicemanagement/mathsettingssystembehaviorobject) - The declaration to configure math behavior at the system level.

