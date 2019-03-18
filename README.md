# Pusher as a microservice
An OMG service for Pusher, it allows to send message to the devices who have subscribed for the channel on event.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.org/heaptracetechnology/microservice-pusher.svg?branch=master)](https://travis-ci.org/heaptracetechnology/microservice-pusher)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-pusher/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-pusher)

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Send Message
```sh
$ omg run send -a appId=<APP_ID> -a cluster=<CLUSTER> -a title=<TITLE> -a message=<MESSAGE> -a channel=<CHANNEL> -a event=<EVENT> -e SECRET=<SECRET> -e KEY=<KEY>
```
## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-pusher
```
### RUN
```
docker run -p 3000:3000 microservice-pusher
```
