---
layout: post
title: Static IPs with Google Container Engine
created: 2015-04-03 23:16:04.43 -0700 PDT
categories: Writeup
tags:
  - Google Cloud
---
The big trick with using a static IP is to specify in your service config

```
portalIPs:
  - put.your.ip.here
```

and, in particular, you must _not_ have

```
createExternalLoadBalancer: true
```

because that will attempt to create one for you.
