# Section 9: Writing Tests

## Lesson #075: Writing tests for our POST handlers

- To get started lesson 75, we need to create new git branch

```shell
git checkout -b 09-075-writing-tests-for-our-POST-handlers
```

- then we need to continue our code by copying from `074-writing-tests-for-our-GET-handlers`

```shell
cd section-09
cp -rfv 074-writing-tests-for-our-GET-handlers 075-writing-tests-for-our-POST-handlers
```

- then update go module path to our new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/075-writing-tests-for-our-POST-handlers
```

### if you use IDE such as VSCode or JetBrains, you can find and replace:
- then we need to find and replace some import path to make it import from correct path
    - find: `section-09/074-writing-tests-for-our-GET-handlers`
    - replace: `section-09/075-writing-tests-for-our-POST-handlers`
    - filter: `*.go`

### if you use NeoVIM like me, we need to use command `args` to set the search scope, then use command `argdo` to replace:
- in neovim use below command to add all `*.go` files into the search scope
```vim
:args ./section-09/075-writing-tests-for-our-POST-handlers/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```vim
:argdo %s#/074-writing-tests-for-our-GET-handlers/#/075-writing-tests-for-our-POST-handlers/#g |update
```

- now the new package is ready for lesson 75 is ready to follow along the tutorial, let start coding.
- so now we continue to write test for our POST handler in file `handler_test.go`
- we are having 3 routes: `search-availability`, `search-availability-json`, `make-reservation`
