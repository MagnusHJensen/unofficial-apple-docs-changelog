# Enable Subscriptions

Declare that your device management service supports subscription management.

## Discussion

Send a POST request to declare that your device management service supports subscription management. Until you enable subscription management for a token, content managers can’t purchase subscriptions into the organizational unit that the token represents.

This request takes no body. The server processes it synchronously and returns the resulting `subscriptionManagement` state.

> 

### Example Request and Response

To declare that an organizational unit doesn’t support subscriptions, use [Disable Subscriptions](/documentation/devicemanagement/disable-subscriptions).

## Topics

### Response

- [SubscriptionManagementResponse](/documentation/devicemanagement/subscriptionmanagementresponse) - A confirmation response that reports your device management service’s subscription management support.
- [ErrorResponse](/documentation/devicemanagement/errorresponse) - The response that contains the error that occurs.

### Content Metadata

- [Managing subscriptions](/documentation/devicemanagement/managing-subscriptions) - Administer auto-renewable subscription seats for your organization.

