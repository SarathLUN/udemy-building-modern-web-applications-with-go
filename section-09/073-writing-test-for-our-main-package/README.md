# Section 9: Writing Tests

## Lesson #073: Writing tests for our main package

- To get started lesson 73, we need to create new git branch

```shell
git checkout -b 09-073-writing-tests-for-our-main-package
```

- then we need to continue our code from `071-finishing-up-our-response-to-user-and-add-alerts`

```shell
cp -rfv section-08/071-finishing-up-our-response-to-user-and-add-alerts section-09/073-writing-test-for-our-main-package
```

- then update go module path to our new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/073-writing-test-for-our-main-package
```

- then we need to find and replace some import path to make it import from correct path
    - find: `section-08/071-finishing-up-our-response-to-user-and-add-alerts`
    - replace: `section-09/073-writing-test-for-our-main-package`
    - filter: `*.go`
- now the new package for lesson 73 is ready to follow along the tutorial, let start coding.