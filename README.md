# Resumme Builder

Build your resume with HTML/CSS and JSON Data

## Local Usage

Create a `data/resume.json` file and enter your data.

The JSON structure follows the [JSON Resume](https://jsonresume.org/) standard.

<i>You can see an example in [data/examples/example.resume.json](data/examples/example.resume.json)</i>

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

e.g example json data request in [data/examples/example.resume.json](data/examples/example.resume.json)

#### Languages

- en
- fr

e.g [data/examples/example.resume.json](data/examples/example.resume.json)

```json
"lang": "fr"
```

### Templates

- classic ([example here](data/examples/example.classic.pdf))
- basic ([example here](data/examples/example.basic.pdf))
- simple ([example here](data/examples/example.simple.pdf))

e.g [data/examples/example.resume.json](data/examples/example.resume.json)

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

### Auto Translation

You can enable automatic translation of your resume by setting the meta `autoTranslate` field to true in your meta JSON
file.

```
  "meta": {
    "template": "classic",
    "lang": "fr",
    "autoTranslate": true
  }
```

When enabled, the application will automatically translate to the language specified in the `lang` field of your JSON
file using the **DeepL API**.
To get a *DeepL authentication key*, [sign up here](https://www.deepl.com/pro#developer) for a Deepl Developer Account.
To use this feature, you'll need to set your *DeepL authentication key* in a `.env` file at the root of the project:

```dotenv
DEEPL_AUTH_KEY=<your-auth-key>
```

If the autoTranslate field is set to true and a lang field is not present in the JSON file, the application will default
to translating to English.

Note that automatic translation is an experimental feature and may not always produce accurate translations.

### TODO

- [x] Parse data to Html
- [x] Generate Pdf
- [x] Build an API
- [x] Handle Multiple Languages
- [x] Handle Auto Translation
- [ ] Build more templates

## How to Contribute

If you want to contribute you can read [Contributing](CONTRIBUTING.md)

## License

This project is under [MIT License](LICENSE)
