GoAT
====

GoAT is a Go Acceptance Testing framework

**NB This is a pre-alpha work in progress!**

For extra info, see the [wiki](http://github.com/necrophonic/goat/wiki)

What?
-----

GoAT aims to provide simple behaviour driven testing for web sites. It defines a user-friendly DSL that can be easily written by non-technical product owners and tooled later by your devs (where there aren't already standard tools included).

Why?
----

You've created an awesome website, so you want to test it right? But your product owners and test experts aren't necessarily technical, and your devs aren't necessarily experienced testers. GoAT allows your non-technical people to specify human readable test specs and your devs to add magic where required.

Getting Started
===============

Prerequisites
-------------

_GoAT_ uses _Selenium Server_ for interaction with your browsers. To that end you'll need to have _java_ and the _selenium server_ jar (and have them running!)

Project layout
--------------

You don't have to follow the default project layout (pretty much everything is configurable) but for the least amount of effort you'll be wanting to do this:
```text

[root]
  + tests
    - // contains test and test suite files
  + pages
    - // contains page definition files
  + results
    - // target folder for generated output

```

Out-Of-The-Box vs. Tweak-It-Yourself
-------------------------------------

> _aka "do I have to write any Go?"_

_GoAT_ is designed to be as simple to use as possible and to that end it is possible to use without writing a single line of Go. You _do_ need to define your pages and tests but they are done via simple json/yaml files.

Simply drop the _GoAT_ binary (appropriate for your platform) into a folder, create the default folder layout around it and start creating page definitions and tests.

If you want to define your own custom instrumentation and assertions then you're going to have to get your hands dirty and use the _GoAT_ packages and API.
