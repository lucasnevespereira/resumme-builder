# resume-builder
Build your resume with HTML/CSS and JSON Data

## Local Usage

Enter your data into the `data/resume.json` file.

Run `make output`

Open `output/resume.pdf` or `output/resume.html`

## Api Usage

`make run`

Example request to generate pdf

```
[POST] http://localhost:9000/pdf
```
body data example in `data/resume.json`

#### Templates
- ``"template": "basic"``

<hr />

### TODO

- [x] Parse data to Html
- [x] Generate Pdf
- [x] Build an API
- [ ] Build more templates
