---
layout: post
title: "Go Testing Patterns and Anti-patterns"

categories: Blog
tags:
  - Go
  - Testing
---

* i-get-replaced-with-a-toc
{:toc}

# Background

Writing tests is more art than science. On top of that, what feels right when initially writing your
test might turn out to be wrong in the long term.

I've been writing Go for over a decade now, and I often find myself talking about testing and how I
do it with new and veteran Gophers alike. In this blog post, I will lay out a few things that I
consider to be good "patterns"
and common "anti patterns" along with some justification, both in hopes that others find it useful,
but also so that I can reference this post rather than having to repeat myself too often.

The advice below applies to what I would term "mature" code-bases. If you are working on your own,
if you are part of a startup, or if you are building a prototype that is going to be thrown away,
this may well not apply to you. In order to help you understand what is going into this advice and
to help you understand my implied assumptions, here are some things that I hold to be true:

* **The test author is not the only audience for a test.**

  The audience for a test (and its failures) is not just the author of the test. Testing is
  certainly a tool for authors to ensure their code is correct while writing it, but once the code
  is complete they are no longer the sole audience for the test.

  The next audience is the code reviewer. For a reviewer, the test gives insight into how the code
  they are reviewing will be used. It also can answer questions about the code, including what
  corner cases were considered in writing it.

  Once the code is submitted, a final audience (and arguably the most important audience) comes into
  play:
  whomever sees the test fail. This could be the author themselves in six months when they've
  forgotten all about the code, it could be whomever has taken over the codebase and is working on a
  refactor, it could be the author of a library trying to diagnose a breakage, or any one of a
  number of other individuals who might end up inadvertently triggering or simply observing the test
  failing.

* **Test failure messages should be actionable.**

  An ideal test failure should not even require reading the test code in order to understand the
  failure.

  This might not be possible in every case, but striving to achieve this will result in improved fix
  times. Explaining not only what the incorrect values were, but also how they were derived and what
  the test was attempting to do in a particular test case _in the failure message itself_ saves the
  reader from having to divine the author's intent without the benefit of local knowledge.

* **Tests are read more often than they are written.**

  This means that it is, on balance, worthwhile if it is slower to write a test if that means that
  it is easier to read, easier to review, easier to maintain, and if its failures are more
  actionable.

# Patterns

First, I'll discuss some patterns that I find particularly helpful when writing tests, 
regardless of the domain.

## Table Driven Tests

Table driven tests are the Go version of "parameterized tests."  The goal of table-driven 
testing is to increase the "signal-to-noise" ratio of a set of a test suite by collecting the
important inputs and outputs together separately from the implementation details of the test.

I think the best way to explain table driven testing is with a few illustrative examples, and 
then I'll go over some tips for writing them.  Every test will be a little different, and you 
don't have to use the same names that I choose below, but you will want them to be visually 
similar to these examples to ensure they're familiar to other programmers.

Here is the table driven test pattern that I use the most frequently for positive tests:

```go
tests := []struct{
  name  string
  input InputTypeHere
  want  OutputTypeHere
}{
  {
    name: "...",
    // ...
  },
  // more test cases
}

for _, test := range tests {
  t.Run(test.name, func(t *testing.T) {
    result, err := FunctionUnderTest(test.input)
    if err != nil {
      t.Fatalf("FunctionUnderTest(%#v) failed: %v", test.input, err)
    }
    if got, want := result, test.want; got != want {
      t.Errorf("FunctionUnderTest(%#v) = %v, want %v", test.input, got, want)
    }
  })
}
```

Here is a good pattern for negative test cases:

```go
tests := []struct{
  name    string
  input   InputTypeHere
  wantErr error
}{
  {
    name:    "example",
    wantErr: errExample,
  },
  // more test cases
}

for _, test := range tests {
  t.Run(test.name, func(t *testing.T) {
    result, err := FunctionUnderTest(test.input)
    if err == nil {
      t.Fatalf("FunctionUnderTest(%#v) succeeded unexpectedly with result: %#v", test.input, result)
    }
    if got, want := err, test.wantErr; !errors.Is(got, want) {
      t.Fatalf("FunctionUnderTest(%#v) failed with error %q, want errors.Is %q", test.input, got, want)
    }
  })
}
```

If you have just a few negative test cases, you can combine them with your positive cases, but
beware the more logical paths there are through the subtest code, the more cognitive load it 
requires to understand the test.

Often the inputs won't be as simple as the above, in which case important subsets of the inputs
or summaries of them can be printed instead. The goal is to communicate what is special about the
input to the reader, not to overwhelm them with too much information.

If the inputs are simpler, however, sometimes even subtest names are overkill.  This is often 
the case for numeric or string-based helper function tests.  For these, both the test table and 
the test cases themselves can often be radically simplified.  For example:

```go
tests := []struct{
  input string
  want  int
}{
  {"Go", "1"},
  {"Hello World", 2},
  {"Hello  World", 2},
  {"Hello, World!", 2},
  {"", "0"},
  {"  ", "0"},
}

for _, test := range tests {
  if got, want := CountWords(test.input), test.want; got != want {
    t.Errorf("CountWords(%q) = %d, want %d", test.input, got, want)
  }
}
```

One good way to tell that you've gone past what can be handled in this "simplified" form is whether
you need to use `continue` or want to use `t.Fatalf` -- if you need either of these, then it's 
probably better to use a subtest so that you can more simply ensure isolation between test cases.

### Tips for table driven tests

Every programmer will have a slightly different way of approaching these.  Some patterns encourage
creating an `args` struct, though I typically find this overkill as a default.  Play around and 
find what works best for you, but here are some things that I think most good implementations of 
this pattern will share, and ways you can ensure your table tests are as readable and 
maintainable as possible:

* **Omit zero value fields** when they aren't important to understand the test case.

  For example, if you are going with the `wantErr bool` approach, it's ideal to not mention 
  errors at all in positive test cases, and only include that field for the failing ones:

  ```go
  {
    name:  "simple",
    input: some.valid.input,
  },
  {
    name:    "invalid",
    input:   some.invalid.input,
    wantErr: true,
  },
  ```
  
* **Pick hierarchical identifier names** for your subtests.

  You can use hierarchical names to organize your subtests, and the `go test` harness will let you
  filter them based on the components.  For example, you can `--test.run=TestYourTest/good/.*`
  to run all positive tests, or you could `--test.run=Test.*/.*/new_user.*` to run all 
  `new_user`-prefixed subtests across all tests whether they're positive or not.

  Using names that are valid identifiers prevents the test harness from having to sanitize them.
  For example, `valid input` will appear in the command-line output as `valid_input`, which you
  won't be able to use as a search to locate the test case if you don't have line numbers to follow.

  There is no set format for what these should look like, but I've found that it's often useful to
  pick a small number of "top level" groups, and use those across the various test cases in a 
  package for consistency.  For example, the following groups are likely to have filterable names:

  * `good/simple`, `good/with_redudant_names`, `good/legacy_format`, `bad/invalid_prefix`, `bad/too_many_separators`
  * `positive/single_factor`, `postive/multiple_factors`, `negative/invalid`

* **Keep the values inside the test case** rather than making common, shared variables, and

* **Use helper functions** to abbreviate test cases.

  It's pretty common that you'll want to have a full object in your test cases so that you have 
  the flexibility to tweak every field, but most of the test cases won't need to write 
  everything out and having the extra syntax can get in the way of seeing which fields are 
  important.  Having local helper functions can drastically increase the signal-to-noise ratio and
  call attention to the values that are actually important for the reader to see.

  <details>
  <summary>
  Expand for a (very contrived) example
  </summary>

  ```go
  func TestAccessControlServer_IsMemberOfGroup(t *testing.T) {
    // Test case helpers
    group := func(name string, members ...string) map[string][]string {
      return map[string][]string{
        name: members,
      }
    }
    isInGroup := func(user, group string) *IsMemberOfGroupRequest {
      return &IsMemberOfGroupRequest{
        CheckType: IsMemberOfGroupRequest_DirectMembersOnly.Enum(),
        User:      proto.String(user),
        Group:     proto.String(group),
      }
    }
    resp := func(result bool) *IsMemberOfGroupResponse {
      return &IsMemberOfGroupResponse{
        UserPresent: proto.Bool(result),
      }
    }

    tests := []struct{
      name   string
      groups map[string][]string
      req    *IsMemberOfGroupRequest
      resp   *IsMemberOfGroupResponse
    }{
      {
        name:   "is_member",
        groups: group("party_members", "garrus", "kaidan", "tali"),
        req:    isInGroup("garrus", "party_members"),
        resp:   resp(true),
      },
      {
        name:   "not_a_member",
        groups: group("party_members", "garrus", "kaidan", "tali"),
        req:    isInGroup("anderson", "party_members"),
        resp:   resp(false),
      },
      {
        name:   "unknown_user",
        groups: group("party_members", "garrus", "kaidan", "tali"),
        req:    isInGroup("saren", "party_members"),
        resp:   &IsMemberOfGroupResponse{
          UserPresent: proto.Bool(false),
          Note:        proto.String("unknown_user"),
        },
      },
    }
  }
  ```
  </details>

  Notice that there are a few things we did _not_ do:

  * Factor the "common" group into a local (or global) variable
  * Define the set of "fixture" groups outside the table itself

  Doing either of the above increases the "cognitive overhead" of the test table. Understanding a
  test case would require remembering important details from elsewhere in the program. By selecting
  self-describing helper names and evocative test data (e.g. "party_members" not "foo")
  it allows the reader to see exactly what the inputs and outputs are without having to sift through
  extra syntax, common but irrelevant fields, etc.

  Test helpers like this let you see only the important parts of the test case, and it makes it very
  easy to copy/paste a test case, make a small change, and have confidence that your changes are
  both correct and also don't affect other test cases the way they might if you had to update a
  shared value.

## Setup Helpers

## Validation helpers

# Anti-patterns

## Assertions

## Mock Generators