# Resumme Builder

Build your resume with HTML/CSS and JSON Data

## Table of Contents

- [Introduction](#introduction)
- [Local Usage](#local-usage)
- [API Usage](#api-usage)
- [Templates](#templates)
- [Image](#image)
- [Languages](#languages)
- [Date Formats](#date-formats)
- [Todo List](#todo-list)
- [How to Contribute](#how-to-contribute)
- [License](#license)

## Introduction

Resumme Builder is a tool that allows you to generate a resume using HTML/CSS templates and JSON data.
It follows the [JSON Resume](https://jsonresume.org/) standard for structuring resume data.

## Local Usage

Create a JSON file with your resume data. You can specify the file path using the `file` flag.

<i>You can see a json file example in [examples/example.resume.json](examples/example.resume.json)</i>

Run the following command to generate the resume in PDF and HTML formats:

```shell
make local file="data/resume.json"
```

Alternatively, you can use the following command:

```shell
go run *.go local -f="data/resume.json"
```

The generated resume files (`resume.pdf` and `resume.html`) will be saved in the `output` directory.

## API Usage

To use Resumme Builder as an API, follow these steps:

Start the server by running the following command:

```
make server
```

You can also use Docker to run the server:

```
make docker-build
make docker-run
```

Send a POST request to `http://localhost:9000/pdf` with the JSON resume data in the request body.
You can use the example JSON data provided in [examples/example.resume.json](examples/example.resume.json).

The server will generate the resume in PDF format and return it as a response.

e.g example json data request in [examples/example.resume.json](examples/example.resume.json)

## Templates

Resumme Builder provides the following templates for generating resumes:

- Classic: [Example](examples/example.classic.pdf)
- Basic: [Example](examples/example.basic.pdf)
- Simple: [Example](examples/example.simple.pdf)
- Oldman: [Example](examples/example.oldman.pdf)
- Stackoverflow: [Example](examples/example.stackoverflow.pdf)
- Figacy: [Example](examples/example.figacy.pdf)

To use a specific template, specify the template name in the JSON resume data:

```json
  "template": "classic"
```

## Image

If you want to include an image in your resume, provide the image URL in the JSON resume data:

```json
 "image": "https://i.imgur.com/tHA5l7T.jpg"
```

Upload your image to a service like [imgur](https://imgur.com/) and copy the direct link.

## Languages

Resumme Builder supports multiple languages for your resume, allowing you to create your resume in a language that suits
your needs. The default language is English (en), but you can choose to use other supported languages as well.

Currently, the following languages are supported:

- English (en_US)
- French (fr_FR)

This will automatically translate labels such as "Education," "Experiences," and other sections based on the chosen
language.

To set the language for your resume, include the following field in the JSON resume data:
e.g [examples/example.resume.json](examples/example.resume.json)

```json
"lang": "fr_FR"
```

## Date Formats

Resumme Builder supports the following date formats:

- `2006-01-02` (e.g., "2024-07-09")
- `2006-01` (e.g., "2024-07")
- `January 2 2006` (e.g., "July 9 2024")
- `January 2006` (e.g., "July 2024")
- `2006` (e.g., "2024")

Example of date fields in JSON resume data:

```json
{
  "education": [
    {
      "institution": "University of Example",
      "area": "Computer Science",
      "studyType": "Bachelor",
      "startDate": "2015-09-01",
      "endDate": "2019-06-30"
    }
  ],
  "work": [
    {
      "company": "Example Corp",
      "position": "Software Engineer",
      "startDate": "2020-01",
      "endDate": "2024-07"
    }
  ]
}
```

## Todo List

- [x] Parse data to HTML
- [x] Generate PDF
- [x] Build an API
- [x] Handle multiple languages (i18n)
- [ ] Expand template options
- [ ] Implement automatic translation support
- [ ] Add unit tests and improve test coverage

Feel free to contribute additional templates and features to enhance the Resumme Builder project!

## How to Contribute

If you want to contribute you can read [Contributing](CONTRIBUTING.md)

## License

This project is under [MIT License](LICENSE)
