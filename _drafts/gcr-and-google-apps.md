---
layout: post
title: Google Container Registry with Google Apps
created: 2015-03-31 19:02:20.338 -0700 PDT
categories: Writeup
tags: 
  - Google Cloud
  - Docker
long: true
---
The [Container Registry](https://cloud.google.com/tools/container-registry/) allows you to easily push your docker images to Cloud Storage.

Nominally, the registry entry for an image will be `gcr.io/projectname/imagename`, where `projectname` is the name of the project on the Developer Console and `containername` is whatever id you want.  At this point, however, both of these only reliably seem to support `A-Za-z_-`.
    
**TL;DR:** If you&#39;re using my [container script](/docker-recipes#containerscript):

    echo REMOTE=gcr.io/projectname/imagename &gt;&gt; container.cfg
    
Or, for Google Apps:
    
    echo REMOTE=b.gcr.io/bucketname/imagename &gt;&gt; container.cfg
    
Then, you can simply

    ./container.sh push

<!-- snip -->

## Bucket Setup

Since I have a Google Apps domain, my Developer Console projects are all `kylelemons.net:$project` which can&#39;t be used as the `projectname` in the gcr.io registry.  I&#39;ve found two ways around this problem

### Solution 1: Separate Project

My first solution was to use a non-Google-Apps project for the Cloud Storage.  This turned out to be somewhat more complicated than I had anticipated.  I did end up getting it working, so I want to try to document it here.

1. If you haven&#39;t already, spin up a GKE Cluster.
    * I believe this will ensure that the right robot accounts are created.
1. Create a public project (`yourproject` below)
1. Create a storage bucket named `artifacts.yourproject.appspot.com`
1. Configure the ACL for the bucket
    * Log into the developer console for your apps project (`example.com:project` below)
    * Open the Permissions tab and copy the Compute Engine Service Account (it will be something like `123...789@project.gserviceaccount.com`)
    	* If you have multiple service accounts listed, you can run `curl http://metadata/computeMetadata/v1beta1/instance/service-accounts/default/email` from a running instance to find the service account to use
    * Go to the public storage bucket in Storage &gt; Cloud Storage &gt; Storage Browser
    * Edit the bucket permissions to add a &#34;User&#34; with the service account as a &#34;Reader&#34;
    * Do the same for the default object permissions
1. Use `gcloud docker push` to push the image.
	* Check the permissions on the `repositories/library/imagename/tag_latest` file within the bucket to ensure that the permissions applied correctly.
    
If you are doing this after the fact, you can use the following command to update the already-created objects (note the `:R` after the service account):

```
gsutil -m acl ch -r -u 123...789@project.gserviceaccount.com:R gs://artifacts.yourproject.appspot.com
```

## Solution 2: Bucket registry

~~It turns out that you can use a special `_b_` prefix to specify a bucket name instead of a project name!~~ You can use `b.gcr.io` to [push to an existing][gcrb] Google Cloud Storage bucket!

With this solution, it&#39;s as simple as

    docker tag your/docker-image b.gcr.io/bucketname/imagename 
    gcloud docker push b.gcr.io/bucketname/imagename

[gcrb]: https://cloud.google.com/tools/container-registry/#using_an_existing_google_cloud_storage_bucket
