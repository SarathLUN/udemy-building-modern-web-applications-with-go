# Section 6: Choosing a project and working with forms

## Lesson #051: Creating the reservation page

- copy project from the previous lesson #050 to #051

```shell
cd section-06
cp -rfv 050-improving-our-form 051-creating-the-reservation-page
cd 051-creating-the-reservation-page 
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-06/051-creating-the-reservation-page # update `go.mod` for new package
vim . # start NeoVim IDE
```

- we add form in page `make-reservation.html`.
- we make 4 input for `First Name`, `Last Name`, `Email`, `Phone Number`.
