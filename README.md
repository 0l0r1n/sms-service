# Sms Service

Redundant implementation of sms service using fallback strategy

## Architecture

Sending Sms can be a critical part of any business. This simple sms service uses a reduntant fallback implementation in case one of the services is offline.

Implemented using the chain of responsibility design pattern.

## API

In order to send a sms, a post request is necessary containing a json body with the following fields:

```json
  {
    "originator" : "Burrito Palace",
    "recipient": "+316XXXXXXX",
    "message": "🌯🌯🌯🌯🌯Congratulations!!!! You have bought so many burritos with us that you have been awarded a burritos-for-life-time subscription🌯🌯🌯🌯"
  }
```

Please note: in case our main service provider is down, we use a fallback implementation. That fallback implementation unfortunately does not allow setting the originator value.