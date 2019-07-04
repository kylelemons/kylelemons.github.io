---
layout: post
title: Static IPs with Google Container Engine
categories: []
tags: []
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