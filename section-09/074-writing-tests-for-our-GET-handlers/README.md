# Section 9: Writing Tests

## Lesson #074: Writing tests for our GET handlers

- To get started lesson 74, we need to create new git branch

```shell
git checkout -b 09-074-writing-tests-for-our-GET-handlers
```

- then we need to continue our code from `073-writing-test-for-our-main-package`

```shell
cd section-09
cp -rfv 073-writing-test-for-our-main-package 074-writing-tests-for-our-GET-handlers
```

- then update go module path to our new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/074-writing-tests-for-our-GET-handlers
```

- then we need to find and replace some import path to make it import from correct path
    - find: `section-08/071-finishing-up-our-response-to-user-and-add-alerts`
    - replace: `section-09/073-writing-test-for-our-main-package`
    - filter: `*.go`
- now the new package is ready for lesson 74 is ready to follow along the tutorial, let start coding.
- first, we start to set up the test file: `setup_test.go` in `internal/handlers`
- then, we work on the test file: `handler_test.go` in `internal/handlers`