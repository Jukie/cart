# cart
Fetcher of build artifacts from Circle CI.

## Install
`go get github.com/Jukie/cart`

> One step closer to continuous delivery

### Get an artifact from the latest green build of your current project's `master`

``` console
$ cart path/to/artifact
```

Authentication uses `$CIRCLE_TOKEN` in your shell's environment or the `-token` flag on the command line.

### Get an artifact from a specific branch

``` console
$ cart -branch feature1 path/to/artifact
```

### Get an artifact from a specific build number

``` console
$ cart -build 42 path/to/artifact
```

### Get an artifact from a specific user/repo

``` console
$ cart -repo Jukie/cart path/to/artifact
```

### All together now

``` console
$ cart -repo Jukie/cart -branch feature1 -o output/elsewhere path/to/artifact
```
