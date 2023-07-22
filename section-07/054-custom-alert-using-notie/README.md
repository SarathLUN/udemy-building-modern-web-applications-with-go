# Section 7: Javascript and CSS

## Lesson #054: custom alert using Notie

- in this lesson, we not yet work on Go code yet, so I only care about HTML, JS, CSS files.
- copy previous project from the previous lesson #053 to #054

```shell
git checkout -b 054-custom-alert-using-notie

cd section-07
cp -rfv 053-making-a-better-date-picker 054-custom-alert-using-notie
cd 054-custom-alert-using-notie
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-07/054-custom-alert-using-notie # update `go.mod` for new package
vim . # start NeoVim IDE
```

- in this lesson, we use [Notie](https://github.com/jaredreich/notie)
- we add CSS and JS CDN as below:

```html
<link
  rel="stylesheet"
  type="text/css"
  href="https://unpkg.com/notie/dist/notie.min.css"
/>

...

<script src="https://unpkg.com/notie"></script>
```

- we create a function to handle alert with 2 param: message and message type

```js

function notify(msg, msgType){
        notie.alert({
                type: msgType,
                text: msg,
            })
    }

```

- then, we just test to call the function by click button.

