# Section 7: Javascript and CSS

## Lesson #058: Using our javascript module on the "Book Now" button 

- in this lesson, we not yet work on Go code yet, so I only care about HTML, JS, CSS files.
- copy previous project from the previous lesson #057 to #058

```shell
git checkout -b 058-using-our-javascript-module-on-the-Book-Now-button

cd section-07
cp -rfv 057-adding-custom-alert-in-our-javascript-module 058-using-our-javascript-module-on-the-Book-Now-button
cd 058-using-our-javascript-module-on-the-Book-Now-button   
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-07/058-using-our-javascript-module-on-the-Book-Now-button # update `go.mod` for new package
vim . # start NeoVim IDE
```

- in this lesson, we will use button [Click me] to simulate the modal form pop-up.
- we create new function in our javascript module, called `custom`.
- we will use code from [Multiple Inputs](https://sweetalert2.github.io/#input-types)
- since this function is the asynchronous function, we need to put keyword `async` in front of keyword `function`
- we will pass the HTML as parameter from `msg`, same scenario for `title`
- we also want to show the Cancel button.
- then we add the datepicker on the modal form by using sweetalert lifecycle hook called `willOpen`.
- then we need to override z-index of the class `datepicker` to make the datepicker pop-up ontop of the modal form.
- then we disable inputs in our modal form so it not auto focus and immediately activate the datepicker which is not good for UX. 
- then we need to enable inputs by using sweetalert lifecycle hook called `didOpen`.
