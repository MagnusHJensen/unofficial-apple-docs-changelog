# Trigger Enhanced Log Collection

Trigger enhanced log collection on the device.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta)

## Discussion

When the organization’s IT support reports a problem to AppleCare, AppleCare may request that the device runs enhanced log collection. AppleCare provides a token to IT support, and they send the `TriggerEnhancedLogCollection` command to the device, with the token, to initiate the enhanced log collection procedure.

When the device processes the command it starts the enhanced log collection operation. There are two modes:

- `interactive`: the device shows the user a notification that allows the user to initiate the enhanced log collection. The device prompts the user to consent to both log collection and log upload at the appropriate times. The user can decline to proceed with enhanced log collection or upload. This mode is required for macOS devices. This mode isn’t available on tvOS devices or Shared iPad.
- `non-interactive`: the device shows a notification that enhanced log collection is in progress. The device collects and uploads the logs in the background without any user intervention. This mode is always available for tvOS devices and Shared iPad. This mode is only available for iOS devices when all these conditions are met:

AppleCare, in conjunction with IT support, determines the mode and the token encodes the mode.

You can cancel an active enhanced log collection session by sending the [Cancel Enhanced Log Collection](/documentation/devicemanagement/cancel-enhanced-log-collection-command) to the device.

### Declarative status

The enhanced log collection process reports declarative status to the device management service using the following status items:

- `enhanced-logging.status`: Reports the device’s enhanced log collection session state. See [StatusEnhancedLogging](/documentation/devicemanagement/statusenhancedlogging).
- `enhanced-logging.applecare-token`: Reports the device’s enhanced log collection session AppleCare token. See [StatusEnhancedLoggingAppleCareToken](/documentation/devicemanagement/statusenhancedloggingapplecaretoken).
- `enhanced-logging.timestamp`: Reports the device’s enhanced log collection session timestamp. See [StatusEnhancedLoggingTimestamp](/documentation/devicemanagement/statusenhancedloggingtimestamp).

### Tokens for testing

You can use a set of test tokens to verify correct operation of your device management service product. The set of tokens are:

- `test-token-normal`: This token activates an interactive enhanced log collection session that results in the device reporting the `finished` declarative state if all parts of the flow succeed. Any failure, cancellation, or decline, result in the device reporting the corresponding state.
- `test-token-normal-headless`: This token activates a non-interactive enhanced log collection session that results in the device reporting the `finished` declarative state if all parts of the flow succeed. Any failure, cancellation, or decline, result in the device reporting the corresponding state.
- `test-token-failed`: This token activates an interactive enhanced log collection session that results in the device reporting the `failed` declarative state if all parts of the flow succeed. Any failure, cancellation, or decline, result in the device reporting the corresponding state.

### Command availability

## Topics

### Commands and responses

- [TriggerEnhancedLogCollectionCommand](/documentation/devicemanagement/triggerenhancedlogcollectioncommand) - The command to trigger enhanced log collection on the device.
- [TriggerEnhancedLogCollectionResponse](/documentation/devicemanagement/triggerenhancedlogcollectionresponse) - A response from the device after it processes the command to trigger enhanced log collection on the device.

