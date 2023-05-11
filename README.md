# Resumme Builder

Build your resume with HTML/CSS and JSON Data

## Local Usage

Create a `data/resume.json` file and enter your data.

The JSON structure follows the [JSON Resume](https://jsonresume.org/) standard.

<i>You can see an example in [examples/example.resume.json](examples/example.resume.json)</i>

Run `make local`

Open `output/resume.pdf` or `output/resume.html`

## Api Usage

```
make server
```

request to generate pdf

```
[POST] http://localhost:9000/pdf
```

e.g example json data request in [examples/example.resume.json](examples/example.resume.json)

#### Languages

- en
- fr

e.g [examples/example.resume.json](examples/example.resume.json)

```json
"lang": "fr"
```

### Templates

- classic ([example here](examples/example.classic.pdf))
- basic ([example here](examples/example.basic.pdf))
- simple ([example here](examples/example.simple.pdf))

e.g [examples/example.resume.json](examples/example.resume.json)

```json
  "template": "classic"
```

<hr />

### Image

Concerning the image, if you want one in your resume you will need to pass in a link to the image json field

```json
 "image": "https://i.imgur.com/tHA5l7T.jpg"
```

For that you can upload your image to a service like [imgur](https://imgur.com/) and copy the direct link.

### TODO

- [x] Parse data to Html
- [x] Generate Pdf
- [x] Build an API
- [x] Handle Multiple Languages
- [ ] Build more templates
- [ ] Handle Auto Translation


## How to Contribute

If you want to contribute you can read [Contributing](CONTRIBUTING.md)


## License

This project is under [MIT License](LICENSE)
