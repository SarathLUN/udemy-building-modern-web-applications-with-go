# Section 6: Choosing a project and working with forms

## Lesson #049: Adding a form to search for availability

- copy project from the previous lesson #048 to #049

```shell
cd section-06
cp -rfv 048-create-a-page-for-each-room 049-adding-a-form-to-search-availability 
cd 049-adding-a-form-to-search-availability 
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-06/049-adding-a-form-to-search-availability # update `go.mod` for new package
vim . # start NeoVim IDE
```

- we add form in page `reservation.html`.
- form has 2 input date at the moment, `start_date` and `end_date`, the form method is `get`.
