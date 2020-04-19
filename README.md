md2pdf
======

Converts a MarkDown file to a PDF file, with an option to add an external stylesheet.

## Usage

```bash
$ go get github.com/AlbinoGeek/md2pdf
$ md2pdf -css your-stylesheet.css source.md target.pdf
```

## Requirements

[wkhtmltopdf/wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)

### Instructions for Fedora 30/31

```bash
$ sudo dnf install wkhtmltopdf
```

## See Also

- [russross/blackfriday](https://github.com/russross/blackfriday)
- [wkhtmltopdf/wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)
