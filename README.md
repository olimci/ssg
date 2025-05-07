# ssg

ssg is a lightweight static site generator i made.
## features

- **markdown/html core**: write content in markdown, use html templates to control layout.
- **minimalist**: no build pipelines, no bundlers, no weird magic. just enough to work.
- **hands-off**: it tries not to fight you. bring your own css, ignore the defaults.
- **customisable**: basic templates included, but you're expected to edit or replace them.
- **cli-based**: includes a small command-line tool:
  - `init`: scaffold a new project
  - `build`: convert markdown to html
  - `dev`: run a local server with live reload
  - `serve`: preview the site

---

## installation

use `go install`:

```bash
go install github.com/olimci/ssg@latest
```

make sure `$GOPATH/bin` is in your `PATH` so you can run `ssg` globally.

---

## getting started

### 1. init a project

```bash
ssg init
```

this sets up a basic layout:

```
.
├── ssg_conf.json
└── site
    ├── content
    │   ├── index.md
    │   └── posts
    │       ├── 1.md
    │       ├── 2.md
    │       └── 3.md
    ├── static
    │   └── styles.css
    └── templates
        ├── index.tmpl
        └── post.tmpl
```

- `content/`: markdown goes here
- `static/`: css/images/etc
- `templates/`: html templates

### 2. run the dev server

```bash
ssg dev
```

runs a live dev server on port `8080` (change with `--port` if needed). watches for changes, rebuilds automatically.

### 3. build for deploy

```bash
ssg build
```

outputs to `dist/` or whatever's in `ssg_conf.json`. ready to throw on netlify, github pages, rsync, etc.

---

## License

Shizuka is open-source and available under the [MIT License](LICENSE).