# Protocol Definition

The commands which can be used in the websocket communication are being tracked in the `protocol.json` located in the `protocol` folder. The format of this file is:

```javascript
{
  "Service Name": { // eg. 'user'
    "id": "ProtocolID",
    "methods": {
      "Method Name": "ProtocolID", // eg. "GetUser": "USR"
    }
  },
}
```

The *Service Name* is the key to an object which holds the *id* of this service and its *methods*. The *id* is the key to a **ProtocolID** (3 character string). *Methods* is the key to another object which stores function names and their **ProtocolIDs**.
