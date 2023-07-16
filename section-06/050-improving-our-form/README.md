# Section 6: Choosing a project and working with forms

## Lesson #050: Improving our form

- copy project from the previous lesson #049 to #050

```shell
cd section-06
cp -rfv 049-adding-a-form-to-search-availability 050-improving-our-form
cd 050-improving-our-form 
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-06/050-improving-our-form # update `go.mod` for new package
vim . # start NeoVim IDE
```

- we edit form in page `reservation.html`.
- we want to put start date and end date in same row.
- by using bootstrap class `form-row` and wrap each form group into class `col`.
- then we make form into the center with class `offset-md-3`
