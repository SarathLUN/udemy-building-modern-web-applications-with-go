# Section 9: Writing Tests

## Lesson #077: Writing tests for our Render package

- To get started lesson 77, we need to create new git branch

```shell
git checkout -b 09-077-writing-tests-for-our-Render-package
```

- then we need to continue our code by copying from `076-writing-tests-for-our-Render-package`

```shell
cd section-09
cp -rfv 076-writing-tests-for-our-Render-package 077-writing-tests-for-our-Render-package
```

- then update go module path to our new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/077-writing-tests-for-our-Render-package
```
here
### if you use IDE such as VSCode or JetBrains, you can find and replace:
- then we need to find and replace some import path to make it import from correct path
    - find: `section-09/076-writing-tests-for-our-Render-package`
    - replace: `section-09/077-writing-tests-for-our-Render-package`
    - filter: `*.go`

### if you use NeoVIM like me, we need to use command `args` to set the search scope, then use command `argdo` to replace:
- in neovim use below command to add all `*.go` files into the search scope
```shell
:args ./section-09/077-writing-tests-for-our-Render-package/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```shell
:argdo %s#/076-writing-tests-for-our-Render-package/#/077-writing-tests-for-our-Render-package/#g |update
```

- now the new package is ready for lesson 77 is ready to follow along the tutorial, let start coding.
- in this lesson we will test function `RenderTemplate` but it not return any error, which is hard to test, so we need to update this function to return error.
- this change is impact to our handlers, but we can ignore that for now.
- now we can start to test `RenderTemplate`, we need to prepare variables such as `pathToTemplates`, `tc`, `r`, `ww`.
- then in the `setup_test.go` we need to create struct for `RespondWriter`.
- now in file `renderers_test.go` we can check the error to check if our test pass or fail.

```shell
=== RUN   TestAddDefaultData
--- PASS: TestAddDefaultData (0.00s)
=== RUN   TestRenderTemplate
2025/01/29 11:18:52 cannot get template from template cache # this log because there is no cache at first run
--- PASS: TestRenderTemplate (0.01s)
PASS
ok      github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/077-writing-tests-for-our-Render-package/internal/renderers      0.541s

```