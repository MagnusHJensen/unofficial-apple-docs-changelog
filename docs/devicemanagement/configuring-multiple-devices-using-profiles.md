# Configuring multiple devices using profiles

Create and deploy configuration profiles to managed devices.

## Overview

Configuration profiles streamline the process of setting up a large number of devices. Custom calendar and email settings, network settings (like Wi-Fi and VPN settings), certificates, and device restrictions, are some of the properties you can configure using configuration profiles.

You have several options for deploying configuration profiles:

- Over the air using a device management service.
- Using Apple Configurator for Mac, available in the Mac App Store (iOS, iPadOS, and tvOS devices only).
- In an email message.
- On a webpage.

> 

## Pair restrictions with capabilities in managed profiles

Configure each managed profile, so it’s clear to the user what new capability they’ll receive on the device when they comply with an associated retstrictions or requirement.

For example, your IT policy may require a device to have a 6-character passcode in order to access your corporate VPN service. You can configure a managed profile to enable the VPN capability for a user once they comply with the passcode restriction. There are multiple ways to accomplish this:

- Deliver the passcode policy and VPN configuration using declarative device management. Define an activation predicate for the VPN configuration to only active if the passcode complies with the organizational policies.
- Deliver a single managed profile with both a passcode restriction payload and a VPN payload. This approach allows a grace period on iOs and iPadOS before a user must set a compliant passcode.
- Deliver a profile with a passcode restriction, poll the device until it indicates compliance, and then deliver the VPN payload.

## Define a profile

Configuration profiles are in a property list format, which any XML tool can read and write.

The configuration property list contains the properties listed in the [TopLevel](/documentation/devicemanagement/toplevel) object. These properties describe the profile and the rules for deploying it. Specific configuration values are stored in an array of payloads in the `PayloadContent` property.

Each payload’s contents contain profile-specific keys (see [Profile-specific payload keys](/documentation/devicemanagement/profile-specific-payload-keys)) and keys that are common to all payloads (see the following list of key definitions).

## Encrypt and sign a profile

Encrypting a profile protects its contents from unauthorized access. The encrypted profile can only be decrypted using a private key previously installed on a device. To encrypt a profile:

1. Remove the `PayloadContent` array and serialize it as a property list. Note that the top-level object in this property list is an array, not a dictionary.
2. CMS-encrypt the serialized property list as enveloped data.
3. Serialize the encrypted data in DER (Distinguished Encoding Rules) format.
4. Set the serialized data as the value of the `EncryptedPayloadContent` key in the profile.

Signing a profile guarantees data integrity. To sign a profile, place the XML property list in a DER-encoded, CMS Signed Data structure. When replacing a signed configuration profile, if you don’t sign the replacement using the exact same signing identity, the device rejects the replacement, unless installing the replacement using the MDM protocol.

## Example SCEP configuration profile

The listing below shows the contents of an example profile, containing a Simple Certificate Enrollment Protocol (SCEP) payload.

```xml
<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE plist PUBLIC "-//Apple Inc//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
   <dict>
      <key>PayloadVersion</key>
      <integer>1</integer>
      <key>PayloadUUID</key>
      <string>Ignored</string>
      <key>PayloadType</key>
      <string>Configuration</string>
      <key>PayloadIdentifier</key>
      <string>Ignored</string>
      <key>PayloadContent</key>
      <array>
         <dict>
            <key>PayloadContent</key>
            <dict>
               <key>URL</key> 
               <string>https://scep.example.com/scep</string> 
               <key>Name</key> 
               <string>EnrollmentCAInstance</string> 
               <key>Subject</key>
               <array> 
                  <array> 
                     <array>
                        <string>O</string>
                        <string>Example, Inc.</string>
                     </array>
                  </array>
                  <array>
                     <array>
                        <string>CN</string>
                        <string>User Device Cert</string>
                     </array>
                  </array>
               </array>
               <key>Challenge</key>
               <string>...</string>
               <key>Keysize</key>
               <integer>1024</integer>
               <key>KeyType</key>
               <string>RSA</string>
               <key>KeyUsage</key>
               <integer>5</integer>
            </dict>
            <key>PayloadDescription</key>
            <string>Provides device encryption identity</string> 
            <key>PayloadUUID</key> 
            <string>fd8a6b9e-0fed-406f-9571-8ec98722b713</string> 
            <key>PayloadType</key> 
            <string>com.apple.security.scep</string> 
            <key>PayloadDisplayName</key> 
            <string>Encryption Identity</string> 
            <key>PayloadVersion</key> 
            <integer>1</integer> 
            <key>PayloadOrganization</key> 
            <string>Example, Inc.</string> 
            <key>PayloadIdentifier</key> 
            <string>com.example.profileservice.scep</string> 
         </dict>
      </array>
   </dict>
</plist>

```

