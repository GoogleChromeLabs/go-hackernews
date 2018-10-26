<p align="center">
  <a href="https://gohackernews.com/">
    <img alt="Go Hacker News" title="Go Hacker News" src="https://i.imgur.com/nbp4kUU.png" width="250">
  </a>
</p>

<p align="center">
  A Hacker News client written in Go
</p>

<p align="center">
  <img alt="PRs Welcome" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg">
</p>

## What is this?

A Hacker News client (yes, another one) written in [Go](https://golang.org/) using [GopherJS](https://github.com/gopherjs/gopherjs).

## What is this built with?

* [GopherJS](https://github.com/gopherjs/gopherjs) to compile Go to JavaScript
* [myitcv.io/react](https://github.com/myitcv/x/tree/master/react) for React bindings
  * [JSX](https://godoc.org/myitcv.io/react/jsx) is supported, but this app does not have any :)
* [Humble/Router](https://github.com/go-humble/router) for routing

### Additional 

* Service Worker added with [Workbox](https://developers.google.com/web/tools/workbox/)
* Hosting on [Firebase](https://www.google.ca/search?q=firebase+hosting&oq=firebase+hosting&aqs=chrome.0.0j69i60l2j69i61j0l2.1327j0j4&sourceid=chrome&ie=UTF-8)

## Setup

1. Fork/clone the repo
2. Install packages:

```bash
go get -u github.com/gopherjs/gopherjs	
go get -u myitcv.io/react myitcv.io/react/cmd/reactGen	
go get -u honnef.co/go/js/xhr github.com/go-humble/router	
```

3. Add GopherJS and ReactGen to PATH:

```bash
export PATH="$(dirname $(go list -f '{{.Target}}' myitcv.io/react/cmd/reactGen)):$PATH"
```

4. Create generated files for each component:

```bash
go generate
```

5. Build the application:

```bash
gopherjs build --output build/script.min.js --minify
```

This will save create `script.min.js` in the `build/` folder. You can use any local testing server in `build/` to boot up the application (for example: `python -m SimpleHTTPServer` if you have Python installed).

## Can I contribute?

Of course you can! Please take a look at the [contributing](./CONTRIBUTING.md) documentation for more info.

## License

Apache 2.0

This is not an official Google product.
