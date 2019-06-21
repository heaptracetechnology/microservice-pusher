# _Pusher_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.org/heaptracetechnology/microservice-pusher.svg?branch=master)](https://travis-ci.org/heaptracetechnology/microservice-pusher)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-pusher/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-pusher)

An OMG service for Pusher, it allows to send message to the devices who have subscribed for the channel on event.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Send Message
```coffee
>>> pusher send appId:'appId' cluster:'cluster' title:'title' message:'messageText' channel:'channelName' event:'eventName'
{"success":"true/false","message":"success/failure message","statusCode":"HTTPstatusCode"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Send Message
```shell
$ omg run send -a appId=<APP_ID> -a cluster=<CLUSTER> -a title=<TITLE> -a message=<MESSAGE> -a channel=<CHANNEL> -a event=<EVENT> -e SECRET=<SECRET> -e KEY=<KEY>
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/pusher/blob/master/LICENSE).

