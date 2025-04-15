# Section 9: Writing Tests

## Lesson #079: Exercise - Writing tests for the Forms package

- To get started lesson 79, we need to create new git branch

```shell
git checkout -b 09-079-exercise-writing-tests-for-forms-package

```
- then we need to continue our code by copying from `078-getting-test-coverage`

```shell
cd section-09
cp -rfv 078-getting-test-coverage 079-exercise-writing-tests-for-forms-package
```

- then update go module path to our new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/079-exercise-writing-tests-for-forms-package
```

### if you use IDE such as VSCode or JetBrains, you can find and replace:
- then we need to find and replace some import path to make it import from correct path
    - find: `section-09/078-getting-test-coverage`
    - replace: `section-09/079-exercise-writing-tests-for-forms-package`
    - filter: `*.go`

### if you use NeoVIM like me, we need to use command `args` to set the search scope, then use command `argdo` to replace:
- in neovim use below command to add all `*.go` files into the search scope
```shell
:args ./section-09/079-exercise-writing-tests-for-forms-package/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```shell
:argdo %s#/078-getting-test-coverage#/079-exercise-writing-tests-for-forms-package#g |update
```

- now the new package is ready for lesson 79 is ready to follow along the tutorial, let start coding.
- let create file `forms_test.go`.
- then we can test function `Valid` and `Required`
- To test a method, the name convention is `Test{receiver_name}_{function_name}`
- As per simple test from **Mr. Sawler**, let run a test to make sure our package is working fine
```shell
=== RUN   TestForms_Valid
--- PASS: TestForms_Valid (0.00s)
=== RUN   TestForms_Required
--- PASS: TestForms_Required (0.00s)
PASS
ok      github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/079-exercise-writing-tests-for-forms-package/internal/forms      0.445s
```
