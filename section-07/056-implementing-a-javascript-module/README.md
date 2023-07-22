# Section 7: Javascript and CSS

## Lesson #056: Implementing a Javascript module

- in this lesson, we not yet work on Go code yet, so I only care about HTML, JS, CSS files.
- copy previous project from the previous lesson #055 to #056

```shell
git checkout -b 056-implementing-a-javascript-module 

cd section-07
cp -rfv 055-create-modals-with-sweetalert 056-implementing-a-javascript-module
cd 056-implementing-a-javascript-module   
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-07/056-implementing-a-javascript-module # update `go.mod` for new package
vim . # start NeoVim IDE
```
- in this lesson, we create a function called `Prompt()`.
- in side this function, we create another function called `toast`, then return it.
- we also create 1 more function called `success`, and add into return object.
- then we instantiate the function at the top.
- after that, we can access to its members when button clicked.
- so this is the principle to of using module, which make our code is re-usable, so we are DRY.
