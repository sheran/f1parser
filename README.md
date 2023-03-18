# F1 Parser

This package will parse articles from various sites related to Formula 1 news.
It will output the Title and Body text of the article after fetching the URL and
parsing the HTML body according to a TOML file telling it what to do.

Look at the `TestBareToml()` function in the testing file. This shows you how to
load up the TOML files, called filters, and then parse a url accordingly.

## Filters

A filter is a TOML file that describes what elements to select and what elements
to discard in an HTML file. When ther `parser.go` file runs, it first builds a
set of `filter.go` files based on the TOML files in the filters director.

### Let's look at the `filter.go` file in more depth:

A filter is basically an interface that implements the following methods
(`parser.go`):

```
Init(string) error
Run() (*models.Post, error)
Match(string) bool
Snippet(string) bool
GetHost() string
```

The implementation of the filter is in the `filter.go` file. A filter is created
from a TOML file by decoding it using the third party package `BurntSushi/toml`.
During the decoding process, the struct fields are filled in. At minimum you
need a `host` entry in the TOML file. But obviously you will want to add more.
Therefore the following fields exist for use:

`host` - Tells the filter and the overall implementation which host the filter
is registered to.

`path` - In case the website you want to parse articles from has categories
separated by a path, you can specify it here. This way the filter knows to
ignore everything else. A use case of this is in a motorsports site that has
Formula 1, Formula E, NASCAR, Indy news articles, then we can limit the parsing
to just `/formula1/` news.

`title` - This is the CSS selector from which to parse the title of the article.
Look through the `filters/` directory for examples here

`body` - This is the CSS selector to get the body text. In our case we are of
the opinion that any `<p>` elements contained within this selector are text
elements that we wish to extract, so thats what we do.

`skip_classes` - In case we encounter a `<p>` element with a specific class that
has text that we do not want, we specify it here as a string array. These
classes will then duly be skipped

`skip_elements` - In case a `<p>` element has a group of elements that also have
text that we want to avoid, we can specify those elements here as a string
array.

`skip_children` - This is a more nuclear option where if a `<p>` element has any
children, that entire element will be skipped. This is a boolean value

`skip_text` - This will tell the filter to skip any text that we don't want. For
instance if a `<p>` element is used to also specify advertising or non-relevant
text, we identify the text itself and specify it in a string array so that that
element is skipped.
