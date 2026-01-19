# Implementing the Enrollment SSO flow

Examine the steps between the user, client, and server in the Enrollment SSO flow.

## Overview

To support Enrollment SSO, an MDM server delivers details about an Enrollment SSO app the system needs to implement the single sign-on behavior for MDM enrollment. Enterprise apps can then use single sign-on to access enterprise websites and other resources. This app needs to contain an extensible SSO extension that the MDM server configures to support the single sign-on behavior. For more information about how to create an app with an enterprise SSO extension, see [Enterprise single sign-on (SSO)](/documentation/AuthenticationServices/enterprise-single-sign-on-sso).

> 

The diagram below illustrates the interactions the [Implementing the simple authentication account-driven enrollment flow](/documentation/devicemanagement/implementing-the-simple-authentication-user-enrollment-flow) and [Implementing the OAuth 2 authentication account-driven enrollment flow](/documentation/devicemanagement/implementing-the-oauth2-authentication-user-enrollment-flow) enrollment flows use when they start.



Enrollment SSO modifies the [Implementing the simple authentication account-driven enrollment flow](/documentation/devicemanagement/implementing-the-simple-authentication-user-enrollment-flow) and [Implementing the OAuth 2 authentication account-driven enrollment flow](/documentation/devicemanagement/implementing-the-oauth2-authentication-user-enrollment-flow) enrollment flows as follows. In step 5 of both flows, an `X-Apple-MDM-ESSO` response header returns. When the client receives that header in the 401 response from the server, the system inserts the following additional steps into the authentication flows before performing the actions in step 5.

### Extract the X-Apple-MDM-ESSO header

In (modified) step 5a, the server returns an HTTP 401 response status to the client and includes an `X-Apple-MDM-ESSO` response header. This response header value needs to be an `https` URL that points to a JSON document with the Enrollment SSO details necessary for downloading and configuring an app with an extensible SSO extension.

### Send the Enrollment SSO document request

In (modified) step 5b, the client sends an HTTPS GET request for the Enrollment SSO URL that the system retrieves in the previous step.

### Return the Enrollment SSO document response

In (modified) step 5c, the server needs to return a JSON document with a JSON object that has the following keys:

The `iTunesStoreID` key value needs to refer to an App Store-hosted app that includes an extensible SSO extension. The system uses this app for single sign-on during the enrollment flow.

The `AssociatedDomains` key is an array of strings that specifies a set of domain names associated with the Enrollment SSO app.

The `AssociatedDomainsEnableDirectDownloads` key is set to `true` to allow claimed site association verification to occur at the domain directly, instead of at Apple’s servers. Use this key only for domains that are inaccessible on the public internet.

Enrollment SSO-capable devices support configuration using MDM profiles (with the `ConfigurationProfile` key), or using declarative management declarations (with the `Declarations` key). At least one of these keys needs to be present in the Enrollment SSO document.

> 

Both keys can be present, so you can use the same Enrollment SSO document with sets of devices where some support declarative management and others don’t. If the device receiving the Enrollment SSO document supports declarative management, the  system uses the declarations instead of the MDM profile. Declarative management-capable devices use the MDM profile if the `Declarations` key isn’t present.

The `ConfigurationProfile` key value is a Base64-encoded profile containing a `com.apple.extensiblesso` payload for configuring the extensible SSO extension in the downloaded app prior to enrollment. This profile needs to contain at least one `com.apple.extensiblesso` payload, and may contain certificate payloads. Other payload types in the profile generate an error.

The `Declarations` key value is an array of Base64-encoded declarative management declarations that need to meet the following requirements:

- The declarations need to contain a single `com.apple.activation.simple` activation declaration that references all the configuration declarations in the `Declarations` key.
- The declarations need to contain one `com.apple.configuration.app.managed` configuration declaration for an App Store-hosted app that includes an extensible SSO extension. The configuration needs to have an `AppStoreID` key with a value that’s the same as the `iTunesStoreID` key in the Enrollment SSO document.
- The declarations need to contain any asset declarations that the app configuration references.
- The declarations need to contain one or more `com.apple.configuration.legacy` configuration declarations. The profile that these configurations reference needs to conform to the same rules as the profile for the `ConfigurationProfile` key in the Enrollment SSO document.
- The system ignores the `AssociatedDomains` and `AssociatedDomainsEnableDirectDownloads` keys in the Enrollment SSO document when using declarative management. Instead, you can set those values with the `Attributes` key in the declarative management app configuration.

When using declarations, the system automatically enables the declarative management feature of the MDM protocol after a successful MDM enrollment, without the server having to first send the MDM `Declarations` command. Because this automatic enablement immediately triggers a declarative management synchronization operation, it’s critical that the MDM server prestages the declarations it sends in the Enrollment SSO’s `Declarations` key, so that those same declarations return during the synchronization. Failure to do so results in the system immediately removing those declarations from the device, with the result that single sign-on in MDM mode no longer works.

### Install the Enrollment SSO app

In (modified) step 5d, the client pauses the enrollment flow on the device to request the user’s permission to download the Enrollment SSO app. If the user grants permission, the client first creates a data-separated volume to store MDM-related managed data. The client then downloads and installs the Enrollment SSO app as a managed app, and stores its data (including any keychain items) on the data-separated volume. The client also installs the configuration profile or declarations from the Enrollment SSO JSON document.

If the user subsequently cancels the overall enrollment flow, the system removes the data-separated volume, along with the Enrollment SSO app and associated configuration profile.

After this step is complete, the client continues with processing the actions in step 5 of the overall enrollment flow.

## Test the Enrollment SSO app

Developers can test their Enrollment SSO app with the account-driven enrollment flow prior to submitting the app for App Store deployment by using a ** variant of the Enrollment SSO JSON document described above.

To use test mode, the developer needs to enable the Enrollment SSO option in the Developer section of the Settings app on their device.

Additionally, the server needs to return an `X-Apple-MDM-ESSO-Developer` response header in (modified) step 5a of the flows described above. This header contains a URL that needs to return a JSON document with a JSON object that has the following keys:

The `AppIDs` key is an array of strings containing App IDs that specify the apps available for use by the Enrollment SSO test mode. This key needs to contain at least one value. App IDs are the concatenation of the app’s Team ID and bundle ID: `team-id.bundle-id`, for example `ABCD1234.com.example.app`.

The other keys in this JSON document are the same as those described for step 5c above, with one exception for the declarations. One or more `com.apple.configuration.app.managed` configuration declarations needs to be present, corresponding to each of the apps that the `AppIDs` key specifies in the Enrollment SSO document. These app configurations need to have a `BundleID` key with values that match the bundle IDs in the `AppIDs` key.

In test mode, the Enrollment SSO flow proceeds as described above. However, the developer receives a prompt to manually install the Enrollment SSO app when the system needs it. The client ensures the app is properly managed and associated with the data-separated volume. The flow then continues as usual using the developer-installed app for single sign-on.

