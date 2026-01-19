# Implementing the OAuth 2 authentication account-driven enrollment flow

Examine the steps between the user, client, server, and Apple services in the OAuth 2 flow.

## Overview

To implement account-driven enrollment, you need to support a series of interactions between the user’s device and your MDM server. The following diagrams illustrate the interactions, and the sections below detail each of the interaction steps.



To implement the OAuth 2 flow, steps 1–4 are identical to the simple flow explained in [Implementing the simple authentication account-driven enrollment flow](/documentation/devicemanagement/implementing-the-simple-authentication-user-enrollment-flow).

### Return the 401 response

In step 5, the server returns an `HTTP 401` response status to the client and includes a `WWW-Authenticate` response header. This response header needs to use the `Bearer` scheme and include the following parameters:

If the client’s enrollment request is invalid, the server returns a standard HTTP error response code (for example, 400 or 403) to halt the enrollment flow on the device. If the server’s response is invalid (for example, missing the `WWW-Authenticate` response header), the system cancels the enrollment.

### Send the authorization request

In steps 6–10, the client starts the OAuth 2 authorization grant flow with a public client type by constructing the authorization request URL from the parameters that return in the 401 response. It then adds a query item with the name `login_hint` and its value as the user account identifier that the user enters.



The client creates an [ASWebAuthenticationSession](/documentation/AuthenticationServices/ASWebAuthenticationSession) using the authorization request URL and a callback scheme that it sets to `apple-remotemanagement-user-login`, and then starts the session.

The authentication session performs an `HTTPS GET` request for the OAuth 2 authorization request URL, and includes the appropriate query parameters needed for the authorization code grant request. The device presents any resulting HTML data to the user in a web view.

The OAuth 2 authorization server responding to the request can prepopulate any user ID form field by extracting the relevant items from the request’s `login_hint` query item. The server can also use that query item to customize the form based on the user name or domain portions of the user account identifier.

[ASWebAuthenticationSession](/documentation/AuthenticationServices/ASWebAuthenticationSession) supports most types of browser-based single sign-on, multifactor, or federated authentication. There can be several round trips between the client and the authorization server to complete authentication.

The user has the option of canceling out of the web view at any time, which terminates the authentication flow and the enrollment.

```swift
<<<<< Request
GET /oauth2/authorization?response_type=code
    &client_id=03FDDE96-FDAB-45EF-A589-0E29C026E824
    &redirect_uri=apple-remotemanagement-user-login:/oauth2/redirection
    &state=340B948D-A84A-45A3-AC45-C93195124B00
    &login_hint=useroauth@example.com HTTP/1.1
Host: mdmserver.example.com

>>>>> Response
HTTP/1.1 200 OK
Content-Type: text/html; charset="utf-8"
Content-Length: 17643

<!DOCTYPE html>
<html>
…
</html>
```

### Return the authorization response

In step 11, the authentication session web flow completes when the server returns an HTTP 308 permanent redirect response to the client, with a `Location` header that it sets to the `redirect-url` parameter value in step 5. The URL also needs to include the OAuth 2 protocol `code` and `state` query items (note the client sets a `state` parameter in its authorization request, so the server needs to set it in the authorization response).

```other
<<<<< Request
POST /oauth2/results HTTP/1.1
Content-Type: multipart/form-data; boundary=----WebKitFormBoundarySaBDj3Bd7BeKC1s2
Content-Length: 236

------WebKitFormBoundarySaBDj3Bd7BeKC1s2
Content-Disposition: form-data; name="username"

user01
------WebKitFormBoundarySaBDj3Bd7BeKC1s2
Content-Disposition: form-data; name="password"

secret
------WebKitFormBoundarySaBDj3Bd7BeKC1s2--

>>>>> Response
HTTP/1.1 308 Permanent Redirect
Content-Length: 0
Location: apple-remotemanagement-user-login:/oauth2/redirection
    ?code=Dek1jUOEcaIPGhbDrCTm9GDV5qT1sb&state=340B948D-A84A-45A3-AC45-C93195124B00
```

### Send the token request

In steps 12–13, the client makes a token access request to fetch the OAuth 2 access and (optional) refresh tokens, using the authorization grant code from step 11, along with other required OAuth 2 parameters. The client securely stores the tokens that the server returns, to use when authorizing subsequent requests to the server.



If authentication fails, the server needs to return an appropriate HTTP error response code that terminates the enrollment on the device.

```other
<<<<< Request
POST /oauth2/token HTTP/1.1
Content-Type: application/x-www-form-urlencoded
Content-Length: 195

grant_type=authorization_code
    &code=Dek1jUOEcaIPGhbDrCTm9GDV5qT1sb
    &redirect_uri=apple-remotemanagement-user-login:/oauth2/redirection
    &client_id=03FDDE96-FDAB-45EF-A589-0E29C026E824

>>>>> Response
HTTP/1.1 200 OK
Content-Length: 245
Content-Type: application/json

{
    "token_type": "Bearer",
    "scope": "MDM",
    "access_token": "dXNlcm9hdXRo.MjZkNTkzZmMtNzk4MC00OWFkLTllZTAtZTA2NzhmNmVhNzg5",
    "expires_in": 3600,
    "refresh_token": "dXNlcm9hdXRo.NzhiNWU0NWQtMDVlZS00MTAwLThhMjEtZDA5MzI2N2RjNzUx"
}
```

### Attempt the second enrollment

In steps 14–21, the behavior is identical to the second enrollment attempt from the [Implementing the simple authentication account-driven enrollment flow](/documentation/devicemanagement/implementing-the-simple-authentication-user-enrollment-flow), with the client using the OAuth 2 access token value in the `Authorization` HTTP request header.



```other
Authorization: Bearer dXNlcm9hdXRo.MjZkNTkzZmMtNzk4MC00OWFkLTllZTAtZTA2NzhmNmVhNzg5
```

