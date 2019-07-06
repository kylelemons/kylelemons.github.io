---
layout: post
title: Docker Recipes
created: 2015-03-28 09:20:50.555 -0700 PDT
categories: []
tags: []
---
This page is primarily for my benefit, but hopefully somebody else will find it useful as well!

# General Notes

## Open Sourceable Dockerfiles

I always like sharing what I do with other people.  That&#39;s part of what I love about the internet.  Thus, I try to make my Dockerfiles reusable by anyone.  I also don&#39;t want to include anything in my Dockerfiles that I would consider sensitive or overly site-specific, like paths to my data volumes.  Here are some tips on making release-ready Dockerfiles.

### Container Script

One of the easiest ways to make a Dockerfile easy to customize is to build a shell script for building and running the container.  The commands you&#39;d use are discussed [below](#startingandstopping), but here is a template for you to start:

&lt;script src=&#34;https://gist.github.com/kylelemons/faf569962ce458121dad.js&#34;&gt;&lt;/script&gt;

In your READMEs, you can easily instruct your users to configure the container with simple commands like

    echo PORT=2368 &gt;&gt; ./container.cfg

### Entrypoint Script

Let&#39;s face it, development and deployment are two very different beasts.  For development, I typically make sure that the base container image (with no volumes mounted) has everything necessary to play around.  This is also useful for writing integration testing scripts, because they can start from a clean slate each time.

# Container Management

## Starting and Stopping

Since I use data volumes copiously (see [below](#persistentdata)), I have never really found myself reusing containers.  I will usually make a [shell script](https://gist.github.com/kylelemons/faf569962ce458121dad) that collects the most common actions for my particular project.

# Persistent Data

# Dockerfiles

## Self-signed Certificates
Often you will find yourself in need of a self-signed certificate as a part of your base image.  I discovered that you can do this noninteractively with one command:

```
RUN openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -subj &#39;/CN=localhost&#39; -days 365
```

This is not something you want to serve publicly, of course, but it is a great one-liner nevertheless.