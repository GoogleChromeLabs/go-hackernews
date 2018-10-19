# Contributing

Contributions are ALWAYS welcome. If you would like to modify, add or fix anything with the site itself (including adding/removing packages), please put up an issue first so we can chat about it!

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement. You (or your employer) retain the copyright to your contribution;
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Where can I contribute?

**The major priority is minimizing the final bundle size and improving the general performance of the site.** Any suggestions here will be much appreciated!

### Improvements to consider

1. Reimplement fetch requests instead of XHR. The default `net/http` package is _extremely_ large unfortunately. [`fasthttp`](https://github.com/valyala/fasthttp) might be worth a shot.
2. Find a lighter alternative to `encoding/json` for JSON decoding. [`jsonparser`](https://github.com/buger/jsonparser) may be a suitable option.
3. Hosting on a different platform and using [Brotli](https://github.com/google/brotli)
4. Find a way to split bundle to separate route-specific chunks.
    - Split chunks by package and then dynamic import where possible? [JSGO](https://github.com/dave/jsgo) might be useful here.
5. Cache dynamic results using our service worker for offline support (need to implement fetch first!)