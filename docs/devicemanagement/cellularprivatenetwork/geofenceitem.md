# CellularPrivateNetwork.GeofenceItem

A geofence for a private network.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, Device Assignment Services , VPP License Management 

## Properties

### GeofenceId

- **Type:** `string`
- **Required:** Yes

A geofence identifier that’s unique within a list of geofences.

### Latitude

- **Type:** `number`
- **Required:** Yes

The latitude of the geofence.

### Longitude

- **Type:** `number`
- **Required:** Yes

The longitude of the geofence.

### Radius

- **Type:** `number`
- **Required:** Yes

Specifies the radius of the geofence in meters. Set this value slightly greater than the private cellular network coverage area.

