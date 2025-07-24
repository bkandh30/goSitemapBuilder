# Go Sitemap Builder

A CLI tool written in Go that crawls websites and generates XML sitemaps.

The tool uses Breadth-First Search (BFS) to traverse web pages and extracts all internal links to create a comprehensive sitemap.

## Installation

```bash
git clone https://github.com/bkandh30/goSitemapBuilder.git
cd goSitemapBuilder
go mod tidy
```

## Usage

### Command-Line Options

| Flag     | Description                 | Default Value             | Example                      |
| :------- | :-------------------------- | :------------------------ | :--------------------------- |
| `-url`   | Target website URL to crawl | `https://gophercises.com` | `-url="https://example.com"` |
| `-depth` | Maximum crawling depth      | `5`                       | `-depth=5`                   |

### Basic Usage

```bash
go run *.go -url="https://example.com"
```

### Advanced Usage

```bash
go run *.go -url="https://example.com" -depth=3
```

## Output

The tool outputs a standard XML sitemap to stdout.

```bash
go run *.go -url="https://example.com" > map.xml
```
