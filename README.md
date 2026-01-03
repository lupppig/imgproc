Absolutely! Let’s make the README **clean, professional, and fully ready** for your `imgproc` project. I’ll format it nicely, organize sections logically, and remove any redundant wording.

---

````markdown
# imgproc

**imgproc** is a high-performance, concurrent CLI tool for image processing built in Go.  
It supports a modular, step-based pipeline with worker pools, retries, live progress metrics, and context-aware cancellation.  

---

## Features

- **Orientation Fix** – auto-correct images using EXIF metadata  
- **Fan-out Thumbnails** – generates multiple sizes (1024px, 256px, 64px)  
- **Resize** – resize images while maintaining aspect ratio  
- **Watermark Removal** – optionally remove watermarks from images  
- **EXIF Stripping** – optionally remove EXIF metadata  
- **Format Conversion** – support for JPEG, PNG, WebP, AVIF  
- **Progress Reporting** – live updates of processed, success, retries, failed images  
- **Concurrent Processing** – worker pool with inflight limiting for high throughput  
- **Context-Aware Cancellation** – safely stop processing with `Ctrl+C`  
- **Retry Mechanism** – retries failed jobs automatically  
- **Automatic Output Directory Creation** – no need to pre-create folders  

---

## Installation

Clone the repository and build the CLI:

```bash
git clone https://github.com/lupppig/imgproc.git
cd imgproc
make build
````

After building, you can run the CLI:

```bash
./bin/imgproc --help
```

---

## Usage

Basic usage:

```bash
./bin/imgproc --input ./images --output ./out --workers 8
```

Example with additional options:

```bash
./bin/imgproc \
  --input ./images \
  --output ./out \
  --workers 4 \
  --format png \
  --resize 800 \
  --watermark true \
  --strip-exif true
```

---

## CLI Flags

| Flag           | Description                                   | Default             |
| -------------- | --------------------------------------------- | ------------------- |
| `--input`      | Input image or directory (required)           | -                   |
| `--output`     | Output directory (required)                   | -                   |
| `--workers`    | Number of concurrent workers                  | Number of CPU cores |
| `--resize`     | Resize width in pixels                        | Original size       |
| `--format`     | Output format (`jpeg`, `png`, `webp`, `avif`) | Original format     |
| `--quality`    | JPEG/WebP quality (ignored for PNG)           | 85                  |
| `--watermark`  | Remove watermark if present                   | false               |
| `--strip-exif` | Remove EXIF metadata                          | false               |
| `--help`       | Show help message                             | -                   |

---

## Progress & Metrics

During processing, **imgproc** shows live progress in the terminal:

```
Processed 15/20 | Success:12 Retry:1 Failed:2
```

* **Processed** – total images processed (success + failed)
* **Success** – successfully processed images
* **Retry** – images retried due to errors
* **Failed** – images that could not be processed

---

## Architecture

* **WorkerPool** – manages concurrent image processing with inflight limiting
* **Step-Based Transformations** – modular pipeline for orientation fix, resize, watermark removal, EXIF stripping, and encoding
* **Retry & Cancellation** – jobs automatically retry on failure and listen for context cancellation

---