---
layout: post
title: "Test driving the new Generics"

categories: Blog
tags:
  - Generics
  - Go
---

This is a collection of my thoughts as I examine the [latest Go generics proposal], announced in [this blog post].

I don't come across many times where I think that generics in Go would solve a problem,
but there are a few times where I've wanted something like it on multiple occasions:

* Channels that represent mutexes.
* Metrics with arbitrary value and dimension types.

[latest Go generics proposal]: https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md
[this blog post]: https://blog.golang.org/generics-next-step
[protocol buffer]: https://blog.golang.org/protobuf-apiv2

## Chantex

Chantex, which is my head-canon name for channels that are taking the place of mutexes,
are a thing that I have started doing a lot after spending some time mulling over Bryan C. Mills'
awesome [Rethinking Classical Concurrency Patterns] talk.

I won't dive too much in detail, but here is what it can look like:

```
type Server struct {
    state chan *state // chantex
}

func (s *Server) Login(ctx context.Context, u *User) error {
    var state *state
    select {
    case state = <-s.state:
        defer func() { s.state <- state }()
    case <-ctx.Done():
        return ctx.Err()
    }
    return state.addUser(u)
}

type state struct{
    activeUsers map[userID]*User
}

func (s *state) addUser(u *User) error {
    if _, ok := s.activeUsers[u.id]; ok {
        return fmt.Errorf("user %q is already logged in", u.name)
    }
    s.activeUsers[u.id] = u
    return nil
}
```

There are a number of benefits here over the standard mutex pattern, but here are my favorite three:
* You can attach methods to the state type that are very difficult to call without "holding" the mutex.
* You can wait for the mutex with a context, cancellation channel, timeout, deadline, etc.
* It is obvious what data is covered by the mutex.

So, I won't try to convince you to start using this pattern if you're not already inclined to do so,
but I wanted to try to see if it could be improved with generics.  [Here's what I came up with][chplay]:

```
type Chantex(type T) chan T

func New(type T)(initialValue *T) Chantex(T) {
	ch := make(Chantex(T), 1)
	ch <- initialValue
	return ch
}

func (c Chantex(T)) Lock() *T {
	return <-c
}

func (c Chantex(T)) Unlock(t *T) {
	c <- t
}

func (c Chantex(T)) TryLock(ctx context.Context) (t *T, err error) {
	select {
	case t = <-c:
		return t, nil
	case <-ctx.Done():
		return t, ctx.Err()
	}
}
```

Using this with the example above, it turns into this:

```
type Server struct {
	state Chantex(state)
}

func (s *Server) Login(ctx context.Context, u *User) error {
	state, err := s.state.TryLock(ctx)
	if err != nil {
		return fmt.Errorf("state lock: %s", err)
	}
	defer s.state.Unlock(state)
	return state.addUser(u)
}

type state struct {
	activeUsers map[userID]*User
}

func (s *state) addUser(u *User) error {
	if _, ok := s.activeUsers[u.id]; ok {
		return fmt.Errorf("user %q is already logged in", u.name)
	}
	s.activeUsers[u.id] = u
	return nil
}
```

It's only two lines shorter, so you'd have to use it over a dozen times before it makes up for its line count delta.
Aside from that, though, it is going to look a lot more familiar to readers who are familiar with the mutex,
and it doesn't have the somewhat-unusual-looking anonymous-defer-in-a-select bit.
Luckily, however, with this approach you don't actually lose the flexibility to use it in a select if you need something more custom:

```
func (s *Server) ActiveUsers(ctx context.Context) ([]*User, error) {
	var state *state
	select {
	case state = <-s.state:
		defer func() { s.state <- state }()
	case <-time.After(1*time.Second):
		panic("state: possible deadlock")
	case <-ctx.Done():
		return nil, fmt.Errorf("state lock: %w", ctx.Err())
	}

	var users []*User
	for _, user := range state.activeUsers {
		users = append(users, user)
	}
	return users, nil
}
```

So, you can actually use this with effectively same code as from before the `Chantex(T)` refactor if it makes sense in context.

Overall I am pretty happy with how this one came out.  Check out the [full code onthe playground][chplay] if you're interested.
I think it would be even more useful for some of the other types discussed in the [Rethinking Classical Concurrency Patterns] talk,
in particular the `Future` and `Pool` types.

The last thing that took me a bit to figure out: how to make a `Locker` method.
This required me to swap over to requring a pointer type explicitly for `Chantex`, when I had originally just made it `type Chantex(type T) chan T`, but this lines up with how I normally use it:

```
type Chantex(type T) chan *T

func (c Chantex(T)) Locker(ptr *T) sync.Locker {
	return locker(T){ptr, c}
}

func (l locker(P)) Lock() {
	*l.ptr = *l.mu.Lock()
}

func (l locker(P)) Unlock() {
	l.mu.Unlock(l.ptr)
}
```

This seems like a worthwhile change, since it potentially avoids any confusion about copying values if the underlying implementation is not well-understood,
and it had effectively no change on the caller side of the API.

[chplay]: https://go2goplay.golang.org/p/zmGE9enFd63
[Rethinking Classical Concurrency Patterns]: https://about.sourcegraph.com/go/gophercon-2018-rethinking-classical-concurrency-patterns

## Metrics

To track data at scale, we have a metrics pipeline that collects data from running servers.  It is similar to [prometheus].
Each metric can be one of a fixed number of types (basically numbers, floats, and strings) and can have a variable number (fixed on a per-metric basis) of dimensions,
which can also be of a (slightly smaller) set of fixed types (strings and ints basically).

_Yes, I realize that this example is actually in the generics proposal, but I wanted to play around with other approaches too._

Code could look something like this:

```go
package mine

import (
    "log"

    "example.com/monitoring/metrics"
    "example.com/monitoring/fields"
) 

var (
    serversOnline = metric.NewStoredIntCounter("servers_online", field.Int("zone"))
)

func init() {
    zone, err := currentZone()
    if err != nil {
        log.Panicf("Failed to determine local zone: %s", err)
    }

    scrape := time.Tick(30*time.Second)
    go func() {
        for {
            <-scrape

            servers, err := countServers()
            if err != nil {
                log.Errorf("Failed to count servers: %s", err)
                continue
            }

            serversOnline.Set(servers, zone)
        }
    }()
}
```

There are a number of downsides of this approach:

* The `Set` method is defined as `(m *intMetric) Set(value int, dimensions ...interface{})`
    * The variadic interface does not provide compile-time type safety
    * The interface boxing is costly when used in tight loops
* This requires creating specific types for each allowable value
    * Each metric type (`int`, `float64`, `string`, etc) needs its own type
    * Each metric "style" ("counter", "gauge", "histogram", etc) needs its own constructor
    * Each field type (`int`, `string`, etc) also requires its own constructor and type
* The implementations require using reflection and/or type switches even though the API surface is required to enumerate possible value types

In other words, this approach is getting the worst of all worlds: it is not performant, not type-safe, and it requires lots of copy/pasted code.

I would like to think that generics could solve this problem, and while I think the current proposal does help,
it still doesn't leave us in what might be the optimal place.  I don't think it precludes it in the future, however, but let's get to it.

### Enumerated Dimension Counts

So far, this is the only approach that will actually work under the generics proposal.
As you will see below though, it is not entirely satisfying.

The "generic" version of the (relevant pieces of) the code could look like this under the current proposal:

```go
serversOnline := metric.NewStoredCounter(int)("servers_online", field.New(string)("zone"))
serversOnline.Set(servers, zone)
```

Unfortuantely, this doesn't get us type-safety: the `Set` method must still be defined as `(m *Metric(T)) Set(value T, dimensions ...interface{})`.  So, instead, we could change the API to look more like this:

<https://go2goplay.golang.org/p/UPsvxoDw9m6>
```go
package metric

type Sample1(type V, D1) struct {
	Timestamp  time.Time
	Value      V
	Dimension1 D1
}

type Metric1(type V, D1) struct {
	Name   string
	Values []Sample1(V, D1)
}

func NewMetric1(type V, D1)(name string) *Metric1(V, D1) {
	return &Metric1(V, D1){Name: name}
}
func (m *Metric1(V, D1)) Set(value V, dim1 D1) {
	m.Values = append(m.Values, Sample1(V, D1){time.Now(), value, dim1})
}
```

With this approach, we do manage to get type-safety for our `Set` function. Note the major downside here though:
instead of having to define one top-level type for each value and dimension type (generics gives us that),
we have to define a new top-level type (two, actually) for each _number of arguments_... you would also need this:

<https://go2goplay.golang.org/p/zpMDuqd9s-F>
```go
package metric

type Sample2(type V, D1, D2) struct {
    Timestamp  time.Time
    Value      V
    Dimension1 D1
    Dimension2 D2
}

type Metric2(type V, D1, D2) struct {
    Name   string
    Values []Sample2(V, D1, D2)
}

func NewMetric2(type V, D1, D2)(name string) *Metric2(V, D1, D2) {
    return &Metric2(V, D1, D2){Name: name}
}
func (m *Metric2(V, D1, D2)) Set(value V, dim1 D1, dim2 D2) {
    m.Values = append(m.Values, Sample2(V, D1, D2){time.Now(), value, dim1, dim2})
}
```

... and so on and so forth.

Hand-crafting this would likely be onerous, so it would likely require code generation for the generic
implementations at each number of dimensions up to some maximum number,
which seems like it almost defeats the purpose, because you could then just generate the code for the metric directly.

There are a few kinds of API that I could envision that might work with some other kind of generic:

### "Lisp" pattern
```
m := field.Add(int)("age", field.Add(string)("name", metric.New(float64)("height")))
m.Set(age, name, height)
```

This seems like it should be somewhat possible at first glance,
particularly if you expand the Set method to be something like

```
m.With(age).With(name).Set(height)
```

However, this requires the With method to include the proper type for the sub-function
in its type parameters, which is not practical recursively.

### "Variadic" pattern

```
m := metric.New(float64, string, int)("height", "name", "age")
m.Set(height, name, age)
```

This pattern is simimlar to the `Set` problem above: there is no way to write the Set method,
which would require some form of variadic type parameters:

```
// Set with [] type parameters defines the recursive component of the type.
//
// This is only what a variadic type parameter *could* look like.
func (m *Metric(T,[D0,DN...])) Set(value T, dim0 D0, dims ...DN) *Sample(T, DN...) {
    return m.Set(value, dims...).addDimension(m.d0.name, dim0)
}
// Set without the [] type parameters defines the "base case" for the recursion.
func (m *Metric(T)) Set(value T) *Sample(T) {
    return &Sample(T){Value: value}
}
```

### "Builder" pattern
```
m := metric.New(float64)("height").WithField(string)("name").WithField(int)("age")
m.Create().With(name).With(age).Set(height)
```

The `Set` is not possible to write for the same reasons as above, and the `WithField`
method is not possible to write because methods may only include the type parameters
of the type itself.

[prometheus]: https://prometheus.io/docs/concepts/metric_types/

## Conclusions

Over the course of experimenting with this, I have a few overall impressions:

* It's pretty inconvenient to predeclare variables of type T just to get access to the zero value.
    A `zero` builtin or allowing `nil` to be used for the zero value of type parameters would be useful.
* It gets very confusing to deal with both pointer- and non-pointer versions of a type parameter.
    In particular, if you want to let the user choose whether they're dealing with a pointer or a value
    type but intend it to be thought of as a reference by using the `(type *T)` 
* Thinking and reasoning about self-referential, recursive, and/or mutually recursive types will melt your brain.
    This is going to be very important for deciding when to use them and when not to use them.
    I suspect that for the time being, we will want to stick to reflection when the readers will not
    already have a good mental model of the type relationships.
* Type parameters as proposed should be used when the type parameters can be inferred unless
    the types being specified are simple and it is easy to understand what the types mean in context.
    In particular, if the user has to specify multiple different versions of a type in a single
    instantiation (think `x := F(T1, []T1, T2(T1))(y, z)` or something)), it is probably too complex.
* Having interface types that can't be used anywhere that interfaces are allowed feels strange.
    I am used to Go features being generally orthogonal:
    any feature can be used with any other, and they interact in ways that intuitively make sense.

Also, I have to say, I really appreciate that the team got this up and running on the playground,
it enables a ton of fun experimentation.
And double kudos for fixing [the bug I found](https://github.com/golang/go/issues/39653) so fast!
