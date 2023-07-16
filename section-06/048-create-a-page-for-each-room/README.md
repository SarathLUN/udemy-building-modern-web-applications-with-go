# Section 6: Choosing a project and working with forms

## Lesson #048: Creating a page for each room

- copy project from the previous lesson #047 to #048

```shell
cd section-06
cp -rfv 047-creating-a-landing-page 048-create-a-page-for-each-room
cd 048-create-a-page-for-each-room
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-06/048-create-a-page-for-each-room # update `go.mod` for new package
vim . # start NeoVim IDE
```

- first we make our carousel fade.
- then we make menu to link to respective page.
- create `generals.html` then load the room image and change button.
- create `majors.html` then load the room image and change button.
- we make image center with the `mx-auto d-block` and create new CSS class with `max-width=50%;`
- create page `contact.html`, `about.html` and `reservation.html`
