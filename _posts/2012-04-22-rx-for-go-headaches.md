---
layout: post
title: "Rx: My prescription for your Go dependency headaches"
created: 2015-03-28 01:35:28.76 -0700 PDT
categories: Writeup
tags: 
    - Go
    - Programming
long: true
---
There has been a lot of discussion on the [Go Nuts](https://groups.google.com/forum/#!forum/golang-nuts) mailing list about how to manage versioning in the nascent [Go](http://golang.org) package ecosystem.  We muttered about it at first, and then muttered about it some more when `goinstall` came about, and there has been a pretty significant uptick in discussion since the `go` tool began to take shape and the [Go 1](http://golang.org/doc/go1.html) release date approached.  In talking with other Gophers at the [GoSF](http://www.meetup.com/golangsf/) meet-up recently, there doesn't seem to be anyone who really has a good solution.

TL;DR: [kylelemons.net/go/rx](http://kylelemons.net/go/rx)

![Relevant XKCD (online package tracking)](https://imgs.xkcd.com/comics/online_package_tracking.png)

_Oddly appropriate, don't you think?_

<!-- snip -->

# The Problem

Before I get too far, let me first summarize the problem.

When developing a Go application, you will most likely find yourself depending on another package, often written by another author.  The ease of utilizing such third-party packages with the `go` tool makes this an even likelier scenario, and it is, in fact, encouraged.  Inevitably, however, the author of some package on which you depend will make a change to his package; this could be anything from an innocuous bug fix to a large-scale API reorganization, and you are suddenly left with two choices: stick with the version you have (often by cloning it locally) or bite the bullet and update.  This is complicated by the fact that you may both directly and indirectly depend on the same package, which means that both your project and your intermediate dependency need to agree on which of the above choices to take, and in a relatively timely manner.

There have been many proposals and complaints, both on- and offline, with respect to this problem.  It's not a problem that's unique to Go, either; tools like Apache's [Maven](http://maven.apache.org/), Ruby's [Bundler](http://gembundler.com/), etc all attempt to solve this problem to a greater or lesser degree.  It is such a prevalent theme in development that a term, [DLL Hell](http://en.wikipedia.org/wiki/DLL_Hell) (and the more technically correct term [dependency hell](http://en.wikipedia.org/wiki/Dependency_hell)), has come into common use to describe it.

# Strategies

The most obvious thing to do is to be paranoid about package maintainers, and thus copy your dependencies into your project.  If this strategy is sufficient, I highly recommend checking out [goven](https://github.com/kr/goven), which will streamline this process (it even rewrites the imports!) for you.  I take a different tack because I am lazy and don't want to have to maintain other people's code.  I also don't think this strategy simplifies the process of pulling in new changes from upstream, because you still have to update them one at a time until/unless something breaks.

The next obvious thing is to specify somewhere what version you want to check out, in the source code, so that go get knows about it and can do the right thing.  This essentially boils down to something like `import "path/package/version"` (though various proposals suggest using `@rev` or similar).  This is certainly a solution, and I suspect we will see tools emerge that will download source and update it to the proper revisions as a `go get` alternative.  I didn't choose this solution because this requires rewriting import paths when you update code and it makes it difficult to ensure that there is only a single version of a library built into the same binary, which can cause problems (if there are more, the `init()` calls will run twice, for one thing).  It also doesn't help with pulling in changes: you still are taking a chance that you'll break something (sometimes without realizing it) whenever you pull from upstream.

Another reasonable strategy is to version-control the entire (or at least the dependencies within) GOPATH(s).  This has the advantage that multiple developers always check out the correct versions, and branches and merges work nicely.  A very simple tool along these lines is being developed as [gogo](http://go.pkgdoc.org/github.com/davecheney/gogo), which allows you to version control your dependencies and share them between developers.  As long as your version control system doesn't mind having other version control systems' (or its own) metadata stored inside it, this will work.  The downside of this is that you are storing a lot of redundant data in your vcs, and it _still_ doesn't address the issue of how to figure out when and if you can update what packages.

# Enter `rx`

So, since my ancient pre-goinstall build tool has been obsoleted, I figured I'd try my hand at distilling a reasonable, achievable set of goals out of the sea of requirements and suggestions and turn them into a tool for people to use.  If you didn't guess this from the previous section, the biggest problem that I think I can solve is helping you figure out what dependencies you can update without breaking your world.  This can probably work in addition to at least a few of the strategies listed above for a more complete versioning solution, depending on your particular needs.  Here are my informal design ideas/goals/requirements/notes:

- It shouldn't try to "solve" dependency hell.  Making people's lives easier is enough for now.
- It should leverage the existing `go` tool and GOPATH conventions as much as possible.
- It should be easy to see the versions of packages, and to change the active one.
- It should be intelligent about updating and notice when an update breaks something else.
- It should be able to save a "known good" set of versions for easy rollback and sharing.
- It should be fun to use, and should not get in the way of the developer.

In that vein, I have started work on [rx](http://kylelemons.net/go/rx), my prescription for your Go dependency version headaches.  It's starting to approach a few of the the requirements above already.  To whet your appetite, here are a few examples of what it can do:

- `rx list` will show you inter-repository dependencies
- `rx tags` will show you the what tags are available in a repository
- `rx prescribe` will update a repository and test its transitive dependents

Each command also has plenty of fun options to play with; `rx tags` has, for instance, options to only show tags that are up- or downgrades. The structure of the program is strongly reminiscent of the design of the `go` tool (and, in fact, uses it for a lot of backend logic), and so should be familiar for most Gophers and fit nicely into your existing workflows.

Installation is, of course, rather simple:
`go get -u kylelemons.net/go/rx`

Here's a brief example of using `rx`:

    $ rx --rescan list | grep rpc
    /gopath/src/github.com/kylelemons/go-rpcgen: codec webrpc main main echoservice main main offload wire webrpc
    $ rx tags go-rpcgen | egrep v\|HEAD
    193746c88dfebdc5462382b93c1038a29496d9af v2.0.0
    a6938fa6ec0fb6a63fefab2c462d3cd1102cc477 v1.2.0
    bf28cdf3e683dd0919800f6916141c17aa93c36d HEAD
    bf28cdf3e683dd0919800f6916141c17aa93c36d v1.1.0
    f73c5c8ea85bdfbdc69e6aa24dd90b43c7265c67 v1.0.0
    $ rx pre go-rpcgen v2.0.0
    ok      github.com/kylelemons/go-rpcgen/codec   0.051s
    ok      github.com/kylelemons/go-rpcgen/examples/echo   0.139s
    ok      github.com/kylelemons/go-rpcgen/examples/remote 0.019s
    ok      github.com/kylelemons/blightbot/bot     0.029s
    ok      github.com/kylelemons/go-paxos/paxos    0.053s
    $ rx tags go-rpcgen | egrep v\|HEAD
    193746c88dfebdc5462382b93c1038a29496d9af HEAD
    193746c88dfebdc5462382b93c1038a29496d9af v2.0.0
    a6938fa6ec0fb6a63fefab2c462d3cd1102cc477 v1.2.0
    bf28cdf3e683dd0919800f6916141c17aa93c36d v1.1.0
    f73c5c8ea85bdfbdc69e6aa24dd90b43c7265c67 v1.0.0
  

There's not a whole lot here, but you can see that the `list` command (in its short form) found the repository and listed the (short) names of the packages that exist under it.  The `--rescan` option told it to actually scan my repositories, instead of using the cached dependency graph.  The `tags` command then showed me the interesting tags in the repository (it's git, so HEAD also shows where it was currently), and then the `prescribe` command updated it to the latest tag.  Notice that the repository's tests were run, as well as tests for packages that depended on packages in that repository (transitively).  They were also built and installed (except binaries, by default), though this isn't displayed unless you use the `-v` option.

# Expected Use Cases

To help elucidate the problem I'm trying to solve, here are a few use cases that I'd like to support.

## Hobbyist Developer

As a single developer, you've probably got a single GOPATH into which all of your dependencies are installed alongside your own projects.  You freely import between them, and everything generally works.  You don't run `go get` very often to pull down remote packages, unless you find a bug that has been fixed or you find a new feature in a newer library.

- The `rx fetch` command will let you fetch the latest changesets without actually applying them.
- The `rx tags --up` command will show you what tags you can upgrade to.
- The `rx prescribe` command will allow you to update to a new tag.
- The `rx prescribe` command automatically builds and tests depenants transitively.
- The `rx prescribe` command will roll back the update if it turns out to have broken something.

## Small Team

As a small team working on a Go project, your concerns are much different from that of a single developer.  You want your team members to easily stay in sync with one another, and you will only rarely pull changes in from upstream once you have your project working with a particular dependency.

- The `--rxdir` flag and RX_DIR environment variable let you version or share an rx configuration.
- The `rx cabinet --save` command saves the versions of all repositories.
- The `rx cabinet --load` command reverts/upgrades repositories to their saved state.
- The `rx cabinet --export` command saves a relocatable cabinet that can be sshared.
- The `rx pin` command lets you configure what repositories are considered for upgrade.
- The `rx auto` command will try to upgrade packages automatically, keeping seamless upgrades.

The common theme among these commands is maintaining a cohesive group of dependency versions.  When you update a dependency (which we've seen that `rx prescribe` can do automatically), you can save that as a "known good" configuration that you can share, save, and (if things go south) restore later.  For packages that are known to misbehave or for the package you're editing, the `rx pin` command allows you to specify manually what behavior they should have (never upgrade, always tip, never change, etc).  To help with exploring what updates might apply seamlessly, the `rx auto` command will do the heavy lifting of figuring out which repositories depend on each other and will successively try updates.

## Large Project

On a large project, you care about most of the same things as a small team, but there is also a good chance that you are working on multiple versions of your software simultaneously.  There is also a good chance that any given developer may have multiple projects on his workstation which are independently versioned.

- The `rx cabinet --exclude` command (and friends) configure exactly what cabinets track.
- The `rx cabinet --diff` command shows differences in dependencies between cabinets.
- The advanced `rx prescribe` optiosn can manage package upgrades `auto` can't handle.

The theme here is that the same commands that worked in a small and medium environment continue to work, but that their concepts can be extended (and modified slightly) to accomodate the needs of a larger development team.  The larger the team is, the more chances are that there will be multiple branches in play, and `rx` will need to understand this.

# The Catch

There are still problems with this approach.  As long as you start with a working project, you should generally be able to keep it working.  You may not be able to ever update a package if one of its dependents never comes into line, though, which leads me to the biggest problem with this approach: it doesn't make it easy to simply install a remote repository that has external dependencies.  It's intended primarily to support development and releasing of e.g. a binary, where your local development environment doesn't matter to the end user.  I'd like for there to be a nice way to import a package's cabinet file when you're importing it (so that your version of rx learns about what versions do and don't work with various dependency versions), but I haven't fully mapped this out.

Another problem which remains currently unsolved is the requirement to manually update when a dependency's API changes.  It would be nice to have some way for the author of a package to provide a way for dependent packages to fix themselves automatically; a tool like `gofix`.  If this convention were widespread enough, it could vastly simplify the process of updating packages.  This is something else about which I am thinking, and I hope that there are good libraries for easily making `gofix`-like tools in the future as well as a convention for including them in your projects.

# Coming Soon

There is a lot of work to do, but I think it's at the point where the best feedback is feedback from real users who have a real need for a tool like this.  The next priorities on my list are:

- Save and restore global repository state
- Intelligently run "upgrade" experiments to find what new tags can be seamlessly integrated
- Support branches and branch switching
- Clean up and document more of the code

Your feedback, constructive criticism, and pull requests are all greatly appreciated!

P.S. I'm slowly cleaning up my [many side-projects](http://github.com/kylelemons) and making sure they work with Go 1.  I'll be listing them on [kylelemons.net/go](http://kylelemons.net/go) as I do, so feel free to e-mail me or find me on IRC if you have a favorite package that you want updated.

# Comments

##  Graham Anderson
My consistent beating of the drum on this one is let the domain that really matters solve the problem. Outside of development, the OS level is where this problem should be solved and the problem set is already well addressed.

In development might be another kettle of fish but at the same time this can be mitigated by developing against canonical versions of other packages, managed by OS packaging tools. 

This model works *really* well hacking Go on Linux, I don't know what OS packaging tools are available for the BSD's outside of FreeBSD ports.

_Posted: April 22, 2012 7:41 AM_

##  Kyle Lemons
That pushes the problem of getting "working" combinations of software to a different person (the package maintainers), and doesn't alleviate the problem of how to upgrade one without possibly breaking all of a package's dependencies.  My approach is not designed to solve the dependency mess.  There are times where an update can be made seamlessly, and a program can do that for you.  For all the times where it requires human intervention, I'll leave that to a human or another tool.

_Posted: May 6, 2012 12:36 AM_
