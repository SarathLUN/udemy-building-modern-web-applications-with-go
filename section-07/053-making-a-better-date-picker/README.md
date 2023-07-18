# Section 7: Javascript and CSS

## Lesson #053: making a better date picker

- in this lesson, we not yet work on Go code yet, so I only care about HTML, JS, CSS files.
- copy previous project from the previous lesson #052 to #053

```shell
git checkout -b 053-making-a-better-date-picker

cd section-07
cp -rfv 052-what-is-javascript-and-why-should-I-care 053-making-a-better-date-picker
cd 053-making-a-better-date-picker 
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-07/053-making-a-better-date-picker # update `go.mod` for new package
vim . # start NeoVim IDE
```

- in this lesson, we use [Vanilla JS Datepicker](https://github.com/mymth/vanillajs-datepicker)
- we add CSS and JS CDN as below:

```html
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/vanillajs-datepicker@1.3.3/dist/css/datepicker-bs4.min.css">

...

<script src="https://cdn.jsdelivr.net/npm/vanillajs-datepicker@1.3.3/dist/js/datepicker-full.min.js"></script>

```

- and below code to activate datepicker base on ID, in which our element ID is 'reservation-dates'

```js
const elem = document.getElementById('reservation-dates');
const datepicker = new Datepicker(elem, {
  // ...options
  format: 'yyyy-mm-dd',
});
```
 - and we set format in the options.
