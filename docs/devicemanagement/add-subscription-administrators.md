# Add Subscription Administrators

Add administrators for subscriptions.

## Discussion

Send a POST request to designate users as administrators for specific subscriptions. The request body requires an `adamIds` array of subscription identifiers and a `clientUserIds` array of user identifiers. The server processes this request synchronously.

### Example Request and Response

## Topics

### Request

- [ManageSubscriptionAdminsRequest](/documentation/devicemanagement/managesubscriptionadminsrequest)

### Response

- [ManageSubscriptionAdminsResponse](/documentation/devicemanagement/managesubscriptionadminsresponse)
- [ErrorResponse](/documentation/devicemanagement/errorresponse) - The response that contains the error that occurs.

### Content Metadata

- [Managing subscriptions](/documentation/devicemanagement/managing-subscriptions) - Administer auto-renewable subscription seats for your organization.

