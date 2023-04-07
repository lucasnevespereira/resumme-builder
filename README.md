# Resumme Builder

Build your resume with HTML/CSS and JSON Data

## Local Usage


Create a `data/resume.json` file and enter your data.

<i>You can see an example in [data/examples/example.resume.json](data/examples/example.resume.json)</i>

Run `make output`

Open `output/resume.pdf` or `output/resume.html`

## Api Usage

```
make run
```

request to generate pdf

```
[POST] http://localhost:9000/pdf
```

e.g example json data request in [data/examples/example.resume.json](data/examples/example.resume.json)

#### Languages

- en
- fr

e.g [data/examples/example.resume.json](data/examples/example.resume.json)

```json
"lang": "fr"
```

### Templates

- classic

e.g [data/examples/example.resume.json](data/examples/example.resume.json)

```json
  "template": "classic"
```

<hr />

### Picture

Concerning the picture, if you want one in your resume you will need to pass in a link to the photo json field

```json
 "photo": "https://i.imgur.com/tHA5l7T.jpg"
```

For that you can upload your picture to a service like [imgur](https://imgur.com/) and copy the direct link.

### TODO

- [x] Parse data to Html
- [x] Generate Pdf
- [x] Build an API
- [x] Handle Multiple Languages
- [ ] Build more templates
