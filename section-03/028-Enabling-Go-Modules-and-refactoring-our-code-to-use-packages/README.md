# Section 3: Building a Basic Web Application

## video 28:

- continue from previous session. 
- because I already initialized the go mod, so I will edit to reflex the actual new package.

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/028-Enabling-Go-Modules-and-refactoring-our-code-to-use-packages
```

- we create package `cmd/web` to be our entry point.
- move our `main.go` file into this directory.

```shell
mkdir -p cmd/web
mv main.go cmd/web/
```

- next we create we organize our package into directory `pkg`.

```shell
mkdir -p pkg/handlers pkg/render
mv handlers.go pkg/handlers/
mv render.go pkg/render/
```

- then we make function `renderTemplate()` to be exportable by just capitalize the function name to `RenderTemplate()`.
- in the handler function we need to import package `render` and access to function `RenderTemplate()`.
- in file `main.go` we import package `handlers` and access to function `Home()` and `About()`.
- then we start our server by below command:

```shell
go run cmd/web/*.go # this will load imported packages and compile.
```