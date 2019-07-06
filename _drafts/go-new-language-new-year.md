---
layout: post
title: "Go: A New Language for a New Year"
created: 2012-01-06 19:30:08.948 -0700 PDT
categories: Blog
tags: 
  - Go
  - Programming
  - Open Source
long: true
---

**As 2011 makes its way out and 2012 takes its place**, it's time for a bit of reflection and a bit of looking forward. I haven't been writing software professionally for particularly long, but I have written software in a number of languages (everything from Pascal to Python), and I think we can all agree that none of the usual suspects are particularly ideal. If you're like me, you hunker down with your favorite set of language features and write your code in as much isolation as possible, so that you can work around or ignore whatever problems your language and/or environment cause you. You probably have some tools lying around to help you do various things for which the language or your editor/environment aren't well-equipped. You probably don't even realize all of the things that bother you about the language after awhile, because it's the norm: every programmer has to deal with them, it just comes with the territory. Please note: I will be comparing Go to other languages extensively in this blog post; do not take it to be an indictment of them or even, really, as reasons to not use them. I'm simply giving my opinion on their differences and why I personally find that Go is a more suitable language for my development. I have used all of the languages that I discuss and will continue to use them when their particular strengths are required.

<!-- snip -->

Let's start with [Python](http://python.org/).  I loved my time writing Python, because I felt like I could be really productive.  I created a library that my team and I could use to replace huge pieces of boilerplate with a few lines of very terse but legible code.  I used just about every new, shiny feature of Python on which I could get my grubby little hands: I used context objects, generators, fancy inheritance and especially I abused the data model in ways that now make me cringe.  My code was a pleasure to read and even more of a pleasure to write.  It got the job done with a really small amount of code, and my team and managers loved me for how productive I was and how productive I was making my team.  What I knew somewhere in the back of my mind but didn't tell anyone was how horribly painful it was to debug the code written this way.  My code reviewers were able to read the code and admit that it looked right, see that the test cases were comprehensive enough and verify that I was covering all of the requirements, but I doubt that they had any idea what was really going on behind-the-scenes or how brittle some of the libraries could be.  Since I had written one of the libraries myself and knew most of the nuances of the Python features employed, I was able to debug problems with the library (or my use of it) with relative ease.  Unfortunately, it often meant tracing through many layers and scrutinizing pieces of code with a fine-toothed comb to find all of the places where Python was secretly calling (or not calling) out to other pieces of code.  Often, it was simply a matter of hiding yet another __special__ function to an object to handle a new use case (for those of you familiar with C++ and the [Rule of 3](http://en.wikipedia.org/wiki/Rule_of_three_(C++_programming)), there needs to be a Python Rule of 17 or something...) that wasn't already covered.  The rest of the team pretty much just used my example code as a template and never ran into these issues, but if they're still using this particular library I really hope that it's not causing them more problems than it solved.  I'm not even going to get into the various issues that flimsy duck-typing and non-static typing can cause, often a long time after the offending code is written...  What remains clear to me is that Python (and langauges like it) are not suitable for large-scale development without very strict adherence to a style guide which precludes almost all of what I described above, and to me that would just take all of the fun out of it.

That, however, is Python, and is the price you expect to pay for such a dynamic language.  Even lower-level languages like C and Pascal have their issues.  Don't get me wrong, I absolutely love C and feel like I am a _deus intra machinam_ whenever I write it; I still think it's the first programming language that a student should learn, and that they shouldn't be able to move to a higher- or lower-level language until they've mastered it.  The problem with these languages, especially C, is that you have to take such care when writing and reviewing any significant amount of code that verifying that the code is correct, robust, and has no memory leaks is the next closest thing to impossible.  The boilerplate and the bookkeeping that have to be done before you even get to the meat of your application or algorithm are tedious and error-prone in the worst way.  There are certainly situations in which such lower-level languages are the only language that _can_ solve your problem, whether it be for performance, real-time/deterministic execution or bare metal proximity reasons, but they just aren't the most suitable languages for large-scale development.


I've also done a bit of graphical (read: drag-and-drop, flow chart) programming, which is interesting to say the least.  I won't go into much detail here, except to say that from my limited exposure, it seems best suited as a special-purpose tool for small- to medium-scale logic at best, and should probably not be used when there are more than a few developers on a particular piece of "code."  It's fun at first, but the limitations show up fast, and after awhile it seems more tedious, repetitive, and limiting than anything else.

Enter [Go](http://golang.org).  It is a new programming language by pretty much any definition, despite having been in development at Google since 2007 and as an open source project since 2009.  It's a compiled, statically typed, garbage collected language with some cool primitives for concurrency.  Despite being compiled and statically typed, it has a very dynamic feel.  The syntax has some concise shortcuts to minimize repetition and variable typing.  The presence of closures, a garbage collector, and a fairly expansive standard library also contribute to its dynamic feel.  These things, coupled with the clarity and explicit nature of Go source code, all help make Go one of the best general purpose languages that I have ever used.  I won't be providing code directly (this blog post is long enough already), but many of the links I provide will be to the documentation or to the relevant portions of the [Go Tour](http://tour.golang.org/), where you can play with the features yourself.

We look forward to the first long-term stable release of the Go standard library and compilers sometime early this year: [Go version 1](http://blog.golang.org/2011/10/preview-of-go-version-1.html).  I've been convinced for awhile now that Go is a production-ready language, but in my mind this piece is really the last thing left barring wider adoption.  There are some exciting language changes leading up to this Go 1 release, and I am really excited to see what gets done after it is finished.  To date, the fast pace at which Go has evolved (both the language and the standard libraries) has posed a challenge for third-party library maintainers, who often have to maintain both a "release" branch of their code and a "weekly" branch due to the rate at which the weekly branch progresses and the relatively high percentage of Go users who are doing their development at (or beyond) the "weekly" branch to take advantage of the new, shiny features or to get the latest bug fixes.  With Go 1 this should change, and the number of libraries available to use with a single command (more on that later) will be able to grow even faster.


## Why Go?

Compared to most languages in use today, Go was designed with many programmers, concurrent systems, and large-scale problems in mind.  Some languages, like C, were simply not designed at a time where multicore and multiprocessor systems were in broad deployment.  Some languages, like Python, weren't originally intended to scale well onto multiple cores and have been slow to adopt strategies for utilizing such resources.  Many languages to this day don't have networking libraries which work predictably and have a clean, understandable API.  Programs written in many languages, especially those with operator overloading, require in-depth knowledge of both the architecture of the program itself and the libraries in use in order to be able to understand the code well.  Go attempts to address all of these problems and more; in my opinion, it does so brilliantly.

Two of the first things you will hear most people talk about when you hear about Go are channels and goroutines.  They're certainly two of the "shiniest" features of the language, and are the reasons why I tried out Go back in late 2009.  They're far from the biggest reasons why I think you should try Go, but they're as good a place as any to start.

### Channels and Goroutines

[These two features](http://tour.golang.org/#61) are inspired by Hoare's Communicating Sequential Processes ([CSP](http://en.wikipedia.org/wiki/Communicating_sequential_processes)).  A goroutine (similar to a [coroutine](http://en.wikipedia.org/wiki/Coroutine)) is a function that runs independently of the calling function.  They're similar to threads in that a scheduler manages which ones are running at any given time, but they are much lighter-weight than their operating system counterparts.  In fact, a Go application with many thousands of goroutines can run comfortably on a single operating system thread.

Channels are a data type in Go which allows for simple, clear communication between goroutines.  They are first-class values that have a data type associated with them.  Any value of the associated type can be sent over the channel and then received from the channel, usually by another goroutine.  They can be passed into functions and goroutines, stored in objects, even sent through other channels.

In most programming languages, when you have multiple threads that need to share data or communicate, you arrange to do so via the use of mutual exclusion locks ([mutexes](http://en.wikipedia.org/wiki/Mutex)).  At first glance, it seems simple: before you access the data you lock it, and when you're done you unlock it.  It gets [more](http://en.wikipedia.org/wiki/Dining_philosophers_problem) [complicated](http://en.wikipedia.org/wiki/Sleeping_barber_problem) [very](http://en.wikipedia.org/wiki/Producers-consumers_problem) [quickly](http://en.wikipedia.org/wiki/Readers-writers_problem), of course.  In Go, a mutex (while available) is not the preferred solution to most problems.  As stated in [Effective Go](http://golang.org/doc/effective_go.html#sharing), the slogan has become:
_Do not communicate by sharing memory; instead, share memory by communicating._Thus instead of using a mutex to protect a shared, global object, you either pass the object around (changing ownership as you do) or you maintain a single owner and all other goroutines access the object by sending requests to that goroutine.

Using these two tools, it turns out to be quite easy to implement solutions to a problem that make logical sense and that very closely reflect the way we might solve the problem in our head.  Small, focused pieces of code are concerned with specific aspects of the solution, and are either given their input via the usual way (function parameters) or continually from channels (they also have the same alternatives available for their output).  Ownership of data is clear, communication is explicit, the code is highly legible.  Such designs are also often easy to unit test; simulating the input and verifying the output via channels is straightforward and very common.

### Closures and Deferred Functions

One of the many features of Go that makes it have a very dynamic feel are [closures](http://tour.golang.org/#37).  Like lambdas in python or blocks in ruby, a function can be declared as a local variable or even passed as a literal to a function.  All local variables in scope at the closure declaration are also in scope within the closure.  They're powerful tools for functional-style APIs and can be a really fun alternative to (or improvement upon) callbacks.  They are also commonly executed directly as a goroutine.

In addition to the "go" directive that runs a function in its own goroutine, Go also has a ["defer" directive](http://golang.org/doc/effective_go.html#defer).  When a function (or closure) call is preceded by the "defer,", the call is evaluated but not executed until just before the function returns.  Multiple deferred calls will be executed in reverse order.  This turns out to be a very powerful tool for using any object or API which requires cleanup: as soon as you create it or lock it, you can defer the cleanup or unlock.  This makes it very easy to spot when you forgot to close a file descriptor or (should you find yourself using one) unlocking a mutex.

These two features provide the ability to confine relevant code all in the same place, which in turn makes reading, debugging, and maintaining the code easier.  This is a common theme, and I think was at or near the top of the Go designers' priority list when designing the language.  The concepts are simple, orthogonal, and flexible without sacrificing readability, debuggability, or maintainability.

### Interfaces

[Interfaces](http://tour.golang.org/#53) in Go are a welcome twist on a familiar concept.  Typically, in a language like Java, interfaces must be explicitly declared by any implementing class. This, of course, means that adding a new interface requires modifying (and having access to modify) all classes that you plan to use in values of that interface type.  The paradigm in Go is reversed: instead of defining an interface and declaring all of the objects that satisfy it, the interface is defined and any object which could satisfy the interface can implicitly be used as a value of the interface.  This also provides a useful tool, the empty interface, which acts as a container for any possible value (similar to Java's Object, from which all classes are derived) because all values satisfy it.

In contrast to interfaces in other languages (for instance, in Java), Go interfaces are often very small, and contain one or just a few methods.  Because of the presence of such small interfaces in the standard library, they have naturally given rise to some common, idiomatic function names and signatures.  Possibly the [best](http://golang.org/pkg/io/#Reader) [examples](http://golang.org/pkg/io/#Writer) of these from the standard library are the `io.Reader` and `io.Writer` interfaces.  Because of the prevalence of objects which implement these two interfaces, stringing together producers and consumers of a data stream is often reminiscent of a unix pipeline.  It is simple, for instance, to open a file, unzip it, decrypt its contents, and stream it directly to an HTTP client because all of these interfaces satisfy or consume the above interfaces.  Even Go's version of the old favorite `printf` can write to any object which satisfies the `io.Writer` interface, which gives rise to the very concise [web-based version of "Hello, World!" in Go](http://tour.golang.org/#56).

### Simplicity and Explicitness

Some of the best features of Go are, unlike those listed above, not something you'll find listed in the specification.  They probably stem from the [guiding philosophy](http://golang.org/doc/go_faq.html#principles) behind its design more than anything else.  I think that the simplicity, orthogonality, and explicitness of the features of the Go language make it a prime candidate for a teaching language, especially at the middle- or high-school level.

First, the features of Go are simple and minimal.  The [entire language specification](http://golang.org/doc/go_spec.html) fits in 45 printed pages.  By contrast, the C++ specification clocks in at 750 pages and Javascript has nearly 250.  It is short enough that an average programmer can actually read it, understand it, and internalize it in its entirety.  This turns out to be quite an asset in a large project setting with multiple developers, each of whom will have a slightly different coding style and may utilize a different different set of features.  Being able to understand all of the language features in use, their nuances, and their side effects turns out to be easy in Go when it could be a nearly impossible task in many others.

Second, the features of Go are designed to be orthogonal to one another.  When you understand two concepts independently, you understand them together.  I have already given one example of this: deferred function calls.  If you understand defer (e.g. that it doesn't make the call immediately, that it evaluates arguments at the defer site, and that it is executed after any subsequent defer calls as the function returns) and you understand closures (that any local variables are in scope, etc), you already understand what happens if you defer a call to a function literal.  Compare this to my Python anecdotes earlier, and you'll find that this is not as ubiquitous a quality as you might expect.  There are also a fair number of C++ and Java features that interact non-orthogonally and require some more in-depth understanding (hence the 750-page C++ specification).

### Ease of Development

This one is more difficult to quantify.  I have found, qualitatively, that writing Go is easier and more fun than pretty much any other language I have used to date.  Dynamic languages like Python and Ruby are fun to learn and are fun languages with which to experiment, but I have found that the amount of time spent debugging them (especially when it's somebody else's code) takes almost all of the fun out of large project and multi-developer situations.  By contrast, languages like C and C++ (and to a somewhat lesser extent, Java) are more painful to write up front but have well-established tools that make analyzing and debugging them a much more pleasant experience.  Until I tried Go, I accepted this as the nature of the beast.  You just had to to choose.  With Go, however, I think the designers managed to strike a balance.

On the development side, I have found that it is easy to translate designs into code.  As I mentioned previously, channels and goroutines make it easy to implement your design in pieces that all connect together and communicate.  The lack of verbose syntaxes for declarations and the implicit nature of interfaces reduce the amount of arguably unnecessary up-front work.  The simplicity of the type system lets you (or forces you to) focus on the implementation instead of fixating on design patterns or type hierarchies.  The lack of certain features in the language like operator overloading, destructors, exceptions, and the like all force code to be explicit.  An oft-quoted mailing list post by Andrew Gerrand states that "[in] Go, the code does exactly what it says on the page."  Error handling, function calls, math, map/slice access and inter-goroutine communication all look exactly like what they are.  When examining a function in isolation, it is a relatively small feat to understand the side-effects of a given statement, which makes the debugging process much more pleasant than in some other languages in which the simplest statement like "a+b" could potentially have unlimited side effects.

Outside of the code itself, there are many benefits to the Go programming language.  The standard library is very complete, cohesive, and easy to understand.  The third party package ecosystem is nurtured by the Go team, and the standard distribution includes an application called "[goinstall](http://golang.org/cmd/goinstall/)" (soon to be bundled into a ["go" meta-command](http://tip.goneat.org/cmd/go/)) by means of which a package hosted in any web-accessible Git, Mercurial, or Subversion repository can be installed (and any such dependencies, ad infinitum) with a single command.  By default, such installations are anonymized and aggregated on a package dashboard, which shows the most- and most-recently installed packages as well as their build status as of the latest release.  The standard distribution also includes a formatting tool called "gofmt" which knows how to format a piece of source code according to "Go style," which reduces or eliminates the need for unproductive debating within a project about things like brace placement or whitespace around operators.  There is also a tool called "gofix" which contains modules to automatically fix up any mechanical changes that have been made to the standard library or language between releases (for instance, changing function names, package imports or method signatures) where possible.  The last of the tools that I will mention is "godoc," which is similar to pydoc or javadoc.  It uses the comments in the source to provide easily accessible documentation both on the command-line and from a web browser.  All of these tools add up to make even the experience of development that doesn't involve code production as easy as possible.

## Summary

Go is an industrial-strength programming language for solving industrial-sized problems.  There are many features that I haven't even touched upon, like the fact that Go is now the third runtime on Google App Engine, which could probably fill another equally long blog post.  If you read this far, I hope you at least take the time to [take the Go tour](http://tour.golang.org) or [peruse the language specification][http://golang.org/doc/go_spec.html].  Keep an eye out for future projects or tools where you might be able to try out a new language.  You may not like it, it may not be accepted in your place of work, or it may simply not be adequate for your needs, but I think that it has the potential to make significant waves in the software industry and I encourage everyone [even if you've never programmed before in your life] to give it a shot.

## Comments (archived)

*Kyle Lemons*

The kinds of things you have to debug in Go are very different.  In Python, often you have to figure out why the function you called on an object doesn't exist sometimes and figure out what call to/from the code in question is generating that value.  In Go, this isn't an issue because of the strict typing.  In Python, you're forced to debug things that a language like Go can tell you up front are errors, which (to me) is a fundamental difference in my experience as a programmer, and can really shorten my cycles.  In Go, the code is also highly explicit, and something like "a=a+b" can only mean one thing.  In Python, that could cause up to (if I recall) four different functions which could have potentially unlimited bugs or side-effects.

GDB support is also in the works, as I said.  In Go, you also have a built-in memory and CPU profiler.

_Posted: January 11, 2012 10:21 PM_

**sakesun**
I see. Thanks. Looking forward to give Go a try.

_Posted: January 12, 2012 4:35 AM_

**A. Non**
Thanks for posting your review. Did you consider erlang? If so, why did you pick Go over Erlang?

_Posted: January 7, 2012 12:58 AM_

**sakesun**
Then what make debugging with Go any different from python print, logging, and pdb ?

_Posted: January 8, 2012 4:33 AM_

**Fei.Yan**
Awesome analysis, cool pieces

_Posted: January 8, 2012 9:31 AM_

**Steven Blenkinsop**
" and that it is executed before any subsequent defer calls as the function returns"

Should that not read "*after* any subsequent..."?

_Posted: January 9, 2012 4:59 AM_

**sakesun**
May I ask, how do you debug Go code ?  gdb ?

_Posted: January 7, 2012 5:03 PM_

**Kyle Lemons**
I have only used gdb with Go in limited contexts.  GDB does not yet fully understand the makeup and operation of a Go binary, though there is a patch pending that adds some of that functionality and there is a script included with the Go distribution that provides some other features when debugging interactively.  I use logging extensively, the Go equivalent of printf, and when necessary I dump the stack of all goroutines (via SIGQUIT).

_Posted: January 7, 2012 9:45 PM_

**Kyle Lemons**
I've never written Erlang, no.  I originally tried Go not long after it was open sourced because of channels and goroutines; I didn't have a chance to write it professionally until about a year ago, and I've never been in a setting in which I could write Erlang.  Go is an easy language to program if you're familiar with other languages, which I believe to be less true of Erlang.  The concurrency model it uses is asynchronous and fault tolerant, which is a very good model for a concurrent language, but it's also very different from Go: channels communicate within a process instead of between them, channels can be synchronous, and channel buffers have a fixed size (mailboxes, I believe, are unlimited).

_Posted: January 7, 2012 9:42 PM_

**Kyle Lemons**
Hmm.  Probably :)

_Posted: March 4, 2012 1:38 AM_

**Kyle Lemons**
I use Go for command-line utilities and client/server applications.  There isn't a coherent idea about GUIs in Go yet (unless you count web interfaces), so that's pretty much what I stick to.  The various things I tinker with are on https://github.com/kylelemons/ if you want to see the sorts of things I use it for in my spare time.

_Posted: March 4, 2012 1:40 AM_

**Anton Kovalyov**
I'm curious what kind of programs do you use Go for? I love the language but having troubles finding a suitable project. I just don't write servers and stuff; but, on the other hand, maybe I should?

_Posted: January 9, 2012 5:06 AM_
