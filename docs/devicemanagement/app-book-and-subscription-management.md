# App, Book, and Subscription Management

Manage apps, books, and subscriptions for your students and employees.

## Topics

### Getting started

- [Getting started with the management API](/documentation/devicemanagement/getting-started-with-the-management-api) - Configure your device management service to handle content and user assignments.
- [Apps and books metadata for organizations](/documentation/devicemanagement/apps-and-books-metadata-for-organizations) - Get metadata for apps and books your organization owns.

### Managing content

- [Managing assets](/documentation/devicemanagement/managing-assets) - Assign and revoke app and book licenses across your organization.
- [Managing subscriptions](/documentation/devicemanagement/managing-subscriptions) - Administer auto-renewable subscription seats for your organization.
- [Managing users](/documentation/devicemanagement/managing-users) - Register and manage users for your organization’s managed location.
- [Setting up and assigning content](/documentation/devicemanagement/setting-up-and-assigning-content) - Distribute purchased licenses to managed users through your device management service.

### Common tasks

- [Using paginated endpoints](/documentation/devicemanagement/using-paginated-endpoints) - Traverse large result sets with page-index and cursor-based pagination.
- [Subscribing to notifications](/documentation/devicemanagement/subscribing-to-notifications) - Monitor events for assets, assignments, and users in your organization.
- [Handling error responses](/documentation/devicemanagement/handling-error-responses) - Investigate and resolve service request errors.

### Configuration management

- [Client Config](/documentation/devicemanagement/client-config-4szk1) - Store client-specific information on the server.
- [Service Config](/documentation/devicemanagement/service-config) - Provides the full list of web service URLs, notification types, request limits, and possible error codes.

### Asset management

- [Get Assets](/documentation/devicemanagement/get-assets-4ski1) - Get the set of assets that your organization manages.
- [Associate Assets](/documentation/devicemanagement/associate-assets) - Associate assets with client user IDs and serial numbers.
- [Disassociate Assets](/documentation/devicemanagement/disassociate-assets) - Disassociate assets from client user IDs and serial numbers.
- [Revoke Assets](/documentation/devicemanagement/revoke-assets) - Revoke assets from client user IDs and serial numbers.
- [Get Assignments](/documentation/devicemanagement/get-assignments-9wv1e) - Get the set of current assignments for users or devices.

### Subscription management

- [Get Subscriptions](/documentation/devicemanagement/get-subscriptions) - Get the subscriptions that your organization manages.
- [Get Subscription Assignments](/documentation/devicemanagement/get-subscription-assignments) - Get the subscription assignments for users in your organization.
- [Associate Subscriptions](/documentation/devicemanagement/associate-subscriptions) - Associate subscriptions with client user IDs.
- [Disassociate Subscriptions](/documentation/devicemanagement/disassociate-subscriptions) - Disassociate subscriptions from client user IDs.
- [Get Subscription Administrators](/documentation/devicemanagement/get-subscription-administrators) - Get the administrators for subscriptions that your organization manages.
- [Add Subscription Administrators](/documentation/devicemanagement/add-subscription-administrators) - Add administrators for subscriptions.
- [Remove Subscription Administrators](/documentation/devicemanagement/remove-subscription-administrators) - Remove administrators from subscriptions.

### User management

- [Get Users](/documentation/devicemanagement/get-users-4mwln) - Get information about a set of users.
- [Create Users](/documentation/devicemanagement/create-users) - Create users to assign apps, books, and subscriptions to.
- [Update Users](/documentation/devicemanagement/update-users) - Update details for existing users.
- [Retire Users](/documentation/devicemanagement/retire-users) - Retire users by client user IDs.

### Event management

- [Event Status](/documentation/devicemanagement/events-status) - Retrieve the status of an asynchronous event.

### Objects and data types

- [Asset](/documentation/devicemanagement/asset) - A product in the store.
- [ResponseAsset](/documentation/devicemanagement/responseasset) - The asset that the organization owns.
- [UnlimitedResponseAsset](/documentation/devicemanagement/unlimitedresponseasset) - An asset with an unlimited license that the organization owns.
- [Assignment](/documentation/devicemanagement/assignment) - The asset assignment for a user or device.
- [RequestUser](/documentation/devicemanagement/requestuser) - The requested user in the organization.
- [ResponseUser](/documentation/devicemanagement/responseuser) - The user in the organization.
- [ResponseSubscription](/documentation/devicemanagement/responsesubscription) - A subscription with its assignment counts.
- [ResponseSubscriptionAssignment](/documentation/devicemanagement/responsesubscriptionassignment) - An assignment of a subscription to a user.
- [SubscriptionCounts](/documentation/devicemanagement/subscriptioncounts) - The subscription assignment counts broken down by assigned and available.
- [SubscriptionCountsBreakdown](/documentation/devicemanagement/subscriptioncountsbreakdown) - The breakdown of subscription counts by renewing and expiring status.
- [ManageSubscriptionsRequest](/documentation/devicemanagement/managesubscriptionsrequest) - The request for subscription management.
- [ManageSubscriptionAdminsRequest](/documentation/devicemanagement/managesubscriptionadminsrequest)
- [ManageSubscriptionAdminsResponse](/documentation/devicemanagement/managesubscriptionadminsresponse)
- [ResponseSubscriptionAdmin](/documentation/devicemanagement/responsesubscriptionadmin)
- [MdmInfo](/documentation/devicemanagement/mdminfo) - Information about the MDM client.
- [EventResponse](/documentation/devicemanagement/eventresponse) - The response that contains the event identifier.
- [ErrorResponse](/documentation/devicemanagement/errorresponse) - The response that contains the error that occurs.
- [StatusResponse](/documentation/devicemanagement/statusresponse)

### Legacy API

- [App and book management (Legacy)](/documentation/devicemanagement/app-and-book-management-legacy) - Manage apps and books for your students and employees.
- [Upgrading to the new management API](/documentation/devicemanagement/upgrading-to-the-new-management-api) - Migrate from API version 1 to version 2 for improved performance.

