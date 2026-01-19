# Dealing with Inactive MDM Devices and Invalid Push Tokens

Handle when devices become unmanageable due to inactivity or invalid push tokens.

## Overview

When a device becomes inactive, your server can take action, such as limiting the deviceʼs access to your organizationʼs resources until the device resumes responding to push notifications.

### Handle When a Device Becomes Inactive

To receive a notification when a device becomes inactive, set the `CheckOutWhenRemoved` key to `true` in the MDM payload. This causes the device to contact your MDM server when it becomes unmanaged.

The following actions cause a device to become inactive:

- Powering off the device.
- Disconnecting from the network.
- Erasing the device.
- Removing the MDM profile.

> 

Because a managed device makes only a single attempt to deliver this message, employ a timeout to detect devices that fail to check out due to network conditions. Your server then sends a push notification periodically to ensure that managed devices are still listening to it.

If the device fails to respond to push notifications after a specified time, you can consider the device inactive. The time to wait before considering that a device is inactive can vary according to your IT policy. A good time period to use ranges from several days to one week.

You don’t need to send push notifications on a daily basis to make sure a device is responding. Appleʼs push notification servers cache your last push notification and deliver it to the device when it reconnects to the network.

