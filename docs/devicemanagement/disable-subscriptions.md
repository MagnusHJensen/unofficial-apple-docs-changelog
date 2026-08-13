# Disable Subscriptions

Declare that your device management service doesn’t support subscription management.

## Discussion

Send a POST request to declare that your device management service doesn’t support subscriptions for the organizational unit that the token represents. Apple School Manager and Apple Business Manager use this declaration to indicate to content managers that the organizational unit doesn’t support subscriptions, rather than leaving its support status unstated.

This request takes no body. The server processes it synchronously and returns the resulting `subscriptionManagement` state.

> 

### Example Request and Response

To declare that an organizational unit supports subscriptions, use [Enable Subscriptions](/documentation/devicemanagement/enable-subscriptions).

## Topics

### Response

- [SubscriptionManagementResponse](/documentation/devicemanagement/subscriptionmanagementresponse) - A confirmation response that reports your device management service’s subscription management support.
- [ErrorResponse](/documentation/devicemanagement/errorresponse) - The response that contains the error that occurs.

### Content Metadata

- [Managing subscriptions](/documentation/devicemanagement/managing-subscriptions) - Administer auto-renewable subscription seats for your organization.

