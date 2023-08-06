# Section 8: Converting our HTML to Go Templates, and creating handlers 

## Lesson #062: Creating a handler that return JSON

- first thing first, let create new git branch to secure our source code with version control

```shell
git checkout -b 08-062-creating-a-handler-that-return-json
```

- then let copy project from previous lesson

```shell
cd section-08
cp -rfv 061-creating-handlers-for-our-forms-and-adding-csrf-protection 062-creating-a-handler-that-return-json
cd 062-creating-a-handler-that-return-json
```

- start Goland from our new project folder

```shell
goland .
```

- in Goland, replace in files by `CMD + SHIFT + R` and find: "/061-creating-handlers-for-our-forms-and-adding-csrf-protection/" replace with "/062-creating-a-handler-that-return-json/"
- As I am using Goland, I also need to refactor the project name.
- 