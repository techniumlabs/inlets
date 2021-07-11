## Contributing

### Guidelines

Here are a few guidelines for contributing:

* If you would like to contribute to the codebase then please raise an issue to propose the change or feature
* Do not work on an issue / PR until it gets a `design/approved` label from a maintainer
* Do not mix feature changes or fixes with refactoring - it makes the code harder to review and means there is more for the maintainers (with limited time) to test

* If you have found a bug please raise an issue and fill out the whole template.
* Don't raise PRs for typos, these aren't necessary - just raise an Issue
* If the documentation can be improved / translated etc please raise an issue to discuss. 

* Please always provide a summary of what you changed, how you did it and how it can be tested.

### Compliance

All commits need to be signed-off in accordance with the Developer Certificate of Origin (DCO) as per below.

The [Derek bot](https://github.com/alexellis/derek) will enforce this policy, if you need help please get in touch.

#### License

This project is licensed under the MIT License.

### Reporting a suspected vulnerability / security issue

If you would like to report a suspected vulnerability / security issue, please email alex@openfaas.com. Bear in mind that this is a community project, and it may take a few days to get back to you. If you have a working code sample in a private GitHub repo, please feel free to give access to that also.

#### Sign your work

The sign-off is a simple line at the end of the explanation for a patch. Your
signature certifies that you wrote the patch or otherwise have the right to pass
it on as an open-source patch. The rules are pretty simple: if you can certify
the below (from [developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
1 Letterman Drive
Suite D4700
San Francisco, CA, 94129

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.

Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

Then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe.smith@email.com>

Use your real name (sorry, no pseudonyms or anonymous contributions.)

If you set your `user.name` and `user.email` git configs, you can sign your
commit automatically with `git commit -s`.

* Please sign your commits with `git commit -s` so that commits are traceable.

### Getting started with a contribution

For development you will need [Golang 1.13 or newer](https://golang.org/dl/)

You can get the code like this:

```bash
go get -u github.com/inlets/inlets
cd $GOPATH/src/github.com/inlets/inlets
```

Contributions are welcome. All commits must be signed-off with `git commit -s` to accept the [Developer Certificate of Origin](https://developercertificate.org).
