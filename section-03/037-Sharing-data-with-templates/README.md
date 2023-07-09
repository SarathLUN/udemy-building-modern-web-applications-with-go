# Section 3: Building a Basic Web Application

## video 037-Sharing-data-with-templates

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/037-Sharing-data-with-templates
```

- when we execute the template, we would be able to pass the data there, which is currently we put `nil` there.
- so now we want to pass data through the 3rd parameter of function `RenderTempalte()`, but what type of data? it can be `string, number, bool, object` we can just create all of these params.
- so we create a new type to whole the template data in `handlers.go` called `TemplateData`.
- when we try to run our program, we got error of cycle import not allowed.
- this is the example to show that Go is strickly on this.
- now we can create another package called `models` and put `TemplateData` in there. This package will never import from any package. Then make sure we change our code to `TemplateDate` from `models`, not `handlers` anymore.
- we're also introducing middleware concept to add default template data by function `AddDefaultData()` in package `renderers`. In the future if we want to add any default data which will able to all the templates, we can use this function.
