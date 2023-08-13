# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #065: Refactoring to use internal packages

- first we need to refactor our project
- replace all path of imported packages from `064-sending-ajax-post-and-generalizing-our-custom-function`  to `065-refactoring-to-use-internal-packages`
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


