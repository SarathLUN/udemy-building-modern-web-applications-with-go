# Section 8: Converting our HTML to Go Templates, and creating handlers 

## Lesson #060: Converting our pages to Go templates

- make sure we keep version control safe, now let checkout new branch

```shell
git checkout -b 08-060-converting-our-pages-to-go-templates
```

- then we start by resuming the Go project structure from section-06/045-enabling-static-files

```shell
cp -rfv section-06/045-enabling-static-files section-08/060-converting-our-pages-to-go-templates
```

- we also need to update our go mod

```shell
cd section-08/060-converting-our-pages-to-go-templates/
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/060-converting-our-pages-to-go-templates
```

- then make sure we have all pages we needed

```shell
cd section-08/060-converting-our-pages-to-go-templates/templates/
touch contact.page.tmpl generals.page.tmpl majors.page.tmpl make-reservation.page.tmpl search-availability.page.tmpl
```

- now we can start to break the part of the page, and file `base.layout.tmpl` is still the main layout file.
- we also update the menu link and corespondent routes and handlers accordingly.
