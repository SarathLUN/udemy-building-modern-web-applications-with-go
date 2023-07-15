# Section 6: Choosing a project and working with forms

## Lesson #046: Creating a landing page

- copy project from the previous lesson #045 to #046

```shell
cd section-06
cp -rfv 046-creating-pages-as-html 047-creating-a-landing-page
cd 047-creating-a-landing-page
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-06/047-creating-a-landing-page # update `go.mod` for new package
vim . # start NeoVim IDE
```

- in this lesson, we working on HTML landing page with bootstrap, the CSS framework.
- we can get bootstrap from this [URL](https://getbootstrap.com/docs/5.3/getting-started/introduction/#quick-start)
- we get `navbar` then make it dark.
- then we remove the search form from navbar.
- then we arrange the menu.
- then arrange the carousel with our prepared images, and captions.
- then we add call to action button with `btn-success`.
- then we arrange footer with 3 columns layout.

