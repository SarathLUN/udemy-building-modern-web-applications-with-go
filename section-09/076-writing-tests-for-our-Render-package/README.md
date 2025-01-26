# Section 9: Writing Tests

## Lesson #076: Writing tests for our Render package

- To get started lesson 76, we need to create new git branch

```shell
git checkout -b 09-076-writing-tests-for-our-Render-package
```

- then we need to continue our code by copying from `075-writing-tests-for-our-POST-handlers`

```shell
cd section-09
cp -rfv 075-writing-tests-for-our-POST-handlers 076-writing-tests-for-our-Render-package
```

- then update go module path to our new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/076-writing-tests-for-our-Render-package
```

### if you use IDE such as VSCode or JetBrains, you can find and replace:
- then we need to find and replace some import path to make it import from correct path
    - find: `section-09/075-writing-tests-for-our-POST-handlers`
    - replace: `section-09/076-writing-tests-for-our-Render-package`
    - filter: `*.go`

### if you use NeoVIM like me, we need to use command `args` to set the search scope, then use command `argdo` to replace:
- in neovim use below command to add all `*.go` files into the search scope
```shell
:args ./section-09/076-writing-tests-for-our-Render-package/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```shell
:argdo %s#/075-writing-tests-for-our-POST-handlers/#/076-writing-tests-for-our-Render-package/#g |update
```

- now the new package is ready for lesson 76 is ready to follow along the tutorial, let start coding.
- first we create our test file in side render package: `setup_test.go` and `renderers_test.go`.
- in this lesson, Mr. Sawler has demonstrated about test fail on the HTTP request without session 

```shell
=== RUN   TestAddDefaultData
--- FAIL: TestAddDefaultData (0.00s)
panic: scs: no session data in context [recovered]
        panic: scs: no session data in context
...
```

- so now we need to setup session for our test, in our file `setup_test.go`.
- then reference `app` with the `testApp`.
- then create function `getSession` in file `renderers_test.go`
- after that we can test `AddDefaultData`.

```shell
=== RUN   TestAddDefaultData
--- PASS: TestAddDefaultData (0.00s)
PASS
ok      github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/076-writing-tests-for-our-Render-package/internal/renderers      0.590s
```
