[![Build](https://github.com/dhcgn/epub-to-md/actions/workflows/build.yml/badge.svg)](https://github.com/dhcgn/epub-to-md/actions/workflows/build.yml)

# epub-to-md

Tool to extract the text from epub files to markdown.

## Usage

### ebpub as md to console

```ps1
epub-to-md.exe -file ../path/to/file.epub
```

### ebpub as md to file

```ps1
epub-to-md.exe -file ../path/to/file.epub -out ../path/to/output.md
```

## Thanks

I really only put together existing great open source projects and have a sufficient solution for my use case. So a special thank you to the developers of these packages who have done the real work.

- [github.com/JohannesKaufmann/html-to-markdown](github.com/JohannesKaufmann/html-to-markdown)
- [github.com/taylorskalyo/goreader/epub](github.com/taylorskalyo/goreader/epub)