---
layout: post
title: Spring Cleaning
categories: []
tags: []
---
On the off-chance that you&#39;re someone who actually reads this blog, you&#39;ll notice that things have changed.  I&#39;ve moved over to using [Ghost](https://ghost.org/) for my blog (despite the fact that it&#39;s a [Node.js](https://nodejs.org/) application).  It&#39;s minimalist, beautiful, responsive, and uses [Markdown](https://daringfireball.net/projects/markdown/)!

# What&#39;s new?

## Docker Containers
I&#39;ve been playing around with [Docker](https://www.docker.com/) a lot lately, and have been thoroughly enjoying it.  I&#39;ve got [nginx](http://nginx.org/), [Minecraft](https://minecraft.net/), [Ghost](https://ghost.org/) (of course), and some personal side projects running in them.  Running third-party apps inside Docker helps with keeping things separate and makes it easy to do the testing on my laptop before pushing to my VM.  I really feel like they shine for development though, where you can build awesome integration tests that exercise everything from basic setup onward. Like magic.

## Reverse Proxy

I&#39;m using [nginx](http://nginx.org/) as my reverse proxy these days... hand-rolling my own was fun, but required more effort in maintenance than I really wanted to expend on into the future.

## Old Blog Posts

I&#39;m working on migrating some of my more interesting blog posts over.  If there is one you are looking for but can&#39;t find, let me know and I&#39;ll pull it back in.