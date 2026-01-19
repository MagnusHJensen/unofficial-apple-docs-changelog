# Generating developer tokens

Create a JSON Web Token to authorize your requests to the Apps and Books for Organizations API.

## Overview

The header of every Apps and Books for Organizations API request requires authorization in the form of a developer token. A developer token is a signed token that authenticates you as a trusted developer and member of the Apple Developer Program.

### Construct your developer token

The Apps and Books for Organizations API supports the JSON Web Token (JWT) specification, so you can pass statements and metadata called **. For more information, see the [JWT specification](https://tools.ietf.org/html/rfc7519) and the available libraries for generating signed JWTs.

Use your developer account to [create a Services identifier and obtain a key ID](https://developer.apple.com/help/account/manage-service-configurations/apps-and-books-for-organizations) and to [locate your Team ID](https://developer.apple.com/help/account/manage-your-team/locate-your-team-id).

Construct a developer token as a JSON object whose header contains:

> 

In the ** payload of the token, include:

A decoded developer token has the following format:

```other
{
     "alg": "ES256",
     "kid": "ABC123DEFG"
}
{
     "iss": "DEF123GHIJ",
     "iat": 1437179036,
     "exp": 1493298100
}
```

After you create the token, sign it with your private key using the ES256 algorithm.

> 

### Authorize requests

If you manage request authorization directly, in all requests, pass the `Authorization: Bearer` header set to the developer token.

```other
curl -v -H 'Authorization: Bearer [developer token]' "https://api.ent.apple.com/v1/test"
```

For more information about requests, responses, and error handling, see [Handling requests and responses](/documentation/devicemanagement/handling-requests-and-responses).

### Limit request rate

The Apps and Books for Organizations API limits the number of requests your app can make using a developer token within a specific period of time. If you exceed this limit, you’ll temporarily receive `429 Too Many Requests` error responses for requests that use the token. This error resolves itself shortly after the request rate has reduced.

