# DeclarationBase

Keys common to all declarations used with the Remote Management protocol.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Discussion

### Reason Codes

- `Error.ActivationFailed`: A configuration or asset cannot be activated due to an activation that failed.
- `Error.AssetCannotBeDeserialized`: The asset data does not conform to the expected data type.
- `Error.AssetCannotBeDownloaded`: The asset data cannot be downloaded.
- `Error.AssetCannotBeVerified`: The downloaded asset data cannot be verified.
- `Error.ConfigurationCannotBeApplied`: The configuration cannot be applied to the device.
- `Error.ConfigurationCannotBeDeserialized`: The configuration is not valid.
- `Error.ConfigurationFailed`: An asset cannot be activated due to a configuration that failed.
- `Error.ConfigurationIsInvalid`: The configuration is not valid for applying to the device.
- `Error.ConfigurationNotSupported`: The configuration is not supported for this platform, scope, or enrollment type.
- `Error.InvalidPayload`: A declaration is not fully loaded.
- `Error.MissingAssets`: A configuration being activated references assets that are not present.
- `Error.MissingConfigurations`: An activation being activated references configurations that are not present.
- `Error.MissingState`: A declaration is missing internal state information.
- `Error.PredicateFailed`: A predicate evaluation failed.
- `Error.UnableToEvaluatePredicate`: A predicate cannot be evaluated.
- `Error.UnableToParsePredicate`: A predicate expression cannot be parsed.
- `Error.UnableToParsePredicateWithCustomOperator`: A predicate expression with a custom operator cannot be parsed.
- `Error.Unknown`: An unrecognized NSError was generated.
- `Error.UnknownDeclarationType`: The declaration type is not known.
- `Error.UnknownPayloadKeys`: A declaration contains unknown payloads keys.
- `Info.NotReferencedByActivation`: A configuration is not referenced in any activation.
- `Info.NotReferencedByConfiguration`: An asset is not referenced in any configuration.
- `Info.Predicate`: A predicate evaluated to false.
- `Info.UnsupportedSettings`: Unsupported settings in a configuration.

## Topics

### Dictionaries

- [DeclarationBase.Payload](/documentation/devicemanagement/declarationbase/payload-data.dictionary)

