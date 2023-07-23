# Section 7: Javascript and CSS

## Lesson #057: Adding custom alert in our Javascript module 

- in this lesson, we not yet work on Go code yet, so I only care about HTML, JS, CSS files.
- copy previous project from the previous lesson #055 to #056

```shell
git checkout -b 057-adding-custom-alert-in-our-javascript-module 

cd section-07
cp -rfv 056-implementing-a-javascript-module 057-adding-custom-alert-in-our-javascript-module
cd 057-adding-custom-alert-in-our-javascript-module   
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-07/057-adding-custom-alert-in-our-javascript-module # update `go.mod` for new package
vim . # start NeoVim IDE
```

- now we implement the [Toast](https://sweetalert2.github.io/#toast) from SweetAlert.
- and instead of hard code, we use parameters for `title`,`position`,`icon`.
- we also set default value if param are not specify.
- we do similar for the success alert.
- we also create another alert for error.
- 
