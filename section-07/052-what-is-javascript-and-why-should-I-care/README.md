# Section 7: Javascript and CSS

## Lesson #052: What is Javascript, and Why should I care?

- in this lesson, we not yet work on Go code yet, so I only care about HTML, JS, CSS files.
- copy previous project from the previous lesson #051 to #052

```shell
git checkout -b 052-what-is-javascript-and-why-should-I-care

cd section-07
cp -rfv section-06/051-creating-the-reservation-page ./052-what-is-javascript-and-why-should-I-care
cd 052-what-is-javascript-and-why-should-I-care
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-07/052-what-is-javascript-and-why-should-I-care # update `go.mod` for new package
vim . # start NeoVim IDE
```

- Javascript execute in the `script` tag in HTML file.
- use function `console.log()` to echo in the browser console.
- use function `alert()` to popup on the browser window.
- use function `document.getElementById()` to select an element that existed within the current HTML page.
- we create button to toggle the CSS class `redText` on paragraph.

```js
document.getElementById('colorButton').addEventListener('click', function () {
  let myEl = document.getElementById('myParagraph')

  if (myEl.classList.contains('redText')) {
    myEl.classList.remove('redText')
  } else {
    myEl.classList.add('redText')
  }
})
```

- we also test Bootstrap validation [here](https://getbootstrap.com/docs/4.6/components/forms/#validation)
