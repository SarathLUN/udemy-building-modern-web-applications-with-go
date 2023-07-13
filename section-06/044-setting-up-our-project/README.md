# Section 6: Choosing a project and working with forms

## Lesson #044: Setting up our project

- copy project from the previous lesson #042

```shell
cd section-06
cp -rfv ../section-05/042-experimenting-with-session 044-setting-up-our-project
cd 044-setting-up-our-project
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-06/044-setting-up-our-project # update `go.mod` for new package
goland . # start Goland IDE
```

- then in Goland we can use **Replace in files** (Cmd + Shift + R) to replace "section-05/042-experimenting-with-session" by "section-06/044-setting-up-our-project"

This is just my own setup, you can prefer your own setup also.
Just need to make sure your project is ready to continue to the next lesson.