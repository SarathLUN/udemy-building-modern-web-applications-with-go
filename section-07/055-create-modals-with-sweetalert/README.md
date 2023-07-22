# Section 7: Javascript and CSS

## Lesson #055: Creating modals with SweetAlert

- in this lesson, we not yet work on Go code yet, so I only care about HTML, JS, CSS files.
- copy previous project from the previous lesson #054 to #055

```shell
git checkout -b 055-create-modals-with-sweetalert 

cd section-07
cp -rfv 054-custom-alert-using-notie 055-create-modals-with-sweetalert
cd 055-create-modals-with-sweetalert  
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-07/055-create-modals-with-sweetalert # update `go.mod` for new package
vim . # start NeoVim IDE
```

- in this lesson, we use [SweetAlert](https://sweetalert2.github.io/#download)
- we add CSS and JS CDN as below:

```html

<link href="
https://cdn.jsdelivr.net/npm/sweetalert2@11.7.18/dist/sweetalert2.min.css
" rel="stylesheet">

...

<script src="
https://cdn.jsdelivr.net/npm/sweetalert2@11.7.18/dist/sweetalert2.all.min.js
"></script>

```

- we create a function to handle SweetAlert with 4 param: title, text, icon, confirmationButtonText

```js

function notifyModal(title, text, icon, confirmationButtonText){
        Swal.fire({
          title: title,
          html: text,
          icon: icon,
          confirmButtonText: confirmationButtonText,
        })
    }

```

- then, we just test to call the function by click button.

