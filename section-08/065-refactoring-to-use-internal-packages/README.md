# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #065: Refactoring to use internal packages

- first we need to refactor our project
- replace all path of imported packages from `/064-sending-ajax-post-and-generalizing-our-custom-function/`  to `/065-refactoring-to-use-internal-packages/`
- in neovim use below command to add all **go** files into the search scope 

```vim
:args ./**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```vim
:argdo %s#/064-sending-ajax-post-and-generalizing-our-custom-function/#/065-refactoring-to-use-internal-packages/#g |update
```

- then edit go mod for our package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/065-refactoring-to-use-internal-packages
```

- now we can start our server and test to make sure it works.

```shell
go run cmd/web/*.go
```

- if there is no error and we can access from browser, means it working fine.
- in this lesson we will move our created packages into another directory called `internal`
- now let create a directory

```shell
mkdir internal
mv -fv pkg/* internal/
```

- so now the path has changed again, we need to refactor our imported path again.
- in neovim, use command `args` to put all **go** files in to search buffer.

```shell
:args section-08/065-refactoring-to-use-internal-packages/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```vim
:argdo %s#/pkg/#/internal/#g |update
```

- now we need to make sure our application is still working fine after these refactoring. Let restart and test application again.
- and Yes! it works as we expected.
- let move on to the next lesson.

