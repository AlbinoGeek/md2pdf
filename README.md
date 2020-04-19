md2pdf
======

Converts a MarkDown file to a PDF file, with an option to add an external stylesheet.

## Usage

```bash
MD2PDFSRC="github.com/AlbinoGeek/md2pdf" \
go install "$MD2PDFSRC" && cd "$GOPATH/src/$MD2PDFSRC" && \
"$GOPATH/bin/md2pdf" -css markdown.css README.{md,pdf}
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
