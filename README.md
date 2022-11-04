# resume-builder

## Local Usage

Enter your data into the `data/resume.json` file.

Run `make output`

Open `output/resume.pdf` or `output/resume.html`

## Api Usage

Run `make run`

Example POST request to generate pdf

```
[POST] http://localhost:9000/pdf
```
<small>body data example in `data/resume.json` </small>

### TODO

- [x] Parse data to Html
- [x] Generate Pdf
- [ ] Build an API
- [ ] Build more templates
