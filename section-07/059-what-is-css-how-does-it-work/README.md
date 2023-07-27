# Section 7: Javascript and CSS

## Lesson #059: What is CSS, How does it work? 

- in this lesson, we not yet work on Go code yet, so I only care about HTML, JS, CSS files.
- copy previous project from the previous lesson #058 to #059

```shell
git checkout -b 059-what-is-css-how-does-it-work

cd section-07
cp -rfv 058-using-our-javascript-module-on-the-Book-Now-button 059-what-is-css-how-does-it-work
cd 059-what-is-css-how-does-it-work   
vim . # start NeoVim IDE
```

- in this lesson, we will make external CSS file then link it from HTML files.
- in folder `static`, we create folder `css`
- in side that folder `css` we create file `style.css`
- then cut stylesheet directive into that file.
- in HTML file use tag `<link>` to embed the stylesheet by using relative path
- apply same to all other HTML files.
- order is matter for stylesheet, means the latest style will override the existing if it overlap.
- class directive is represented by `.`.
- id directive is represented by `#`.
- tag directive is represented by its tag.
- one html element can have more than 1 class same time.
- in 1 html document can have only 1 ID.
 
