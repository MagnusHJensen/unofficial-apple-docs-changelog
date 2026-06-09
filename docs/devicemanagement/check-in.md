# Check-in

Authenticate devices and maintain push tokens.

## Overview

The Mobile Device Management (MDM) check-in protocol validates a deviceʼs eligibility for enrollment and informs the device management service that a deviceʼs push token has been updated.

After installing the MDM payload, the device initiates communication with the check-in service. The device validates the TLS certificate of the service, then uses the identity specified in its MDM payload as the client authentication certificate for the connection.

If the MDM payload includes a check-in URL, the check-in protocol communicates with that check-in service. If it doesn’t provide a check-in URL, the device uses the main device management service URL instead.

## Topics

### Commands

- [Authenticate](/documentation/devicemanagement/authenticate) - Authenticates a user during MDM payload installation.
- [User Authenticate](/documentation/devicemanagement/user-authenticate) - Authenticates a user with a two-step authentication protocol.
- [Check Out](/documentation/devicemanagement/check-out) - Responds to the removal of the MDM enrollment profile from a device.
- [Get Token](/documentation/devicemanagement/get-token) - Gets a token from the server.
- [Token Update](/documentation/devicemanagement/token-update) - Updates the token for a device on the server.
- [Get Bootstrap Token](/documentation/devicemanagement/get-bootstrap-token) - Gets the bootstrap token from the server.
- [Set Bootstrap Token](/documentation/devicemanagement/set-bootstrap-token) - Sends the bootstrap token to the server.
- [Return To Service](/documentation/devicemanagement/return-to-service) - Gets the return-to-service configuration from the server.

### Declarative management

- [Declarative Management](/documentation/devicemanagement/declarative-management) - Sends declarative management requests to the server.
- [Get Server Supported Declarations](/documentation/devicemanagement/declaration-items) - Get a list of the declarations available on the server.
- [Get the Device Status](/documentation/devicemanagement/status) - The request for getting the status of a device.
- [Get the Device Token](/documentation/devicemanagement/tokens) - The request for sending the device token details.

