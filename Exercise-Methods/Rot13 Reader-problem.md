## Exercise: rot13Reader

A common pattern is an [io.Reader](https://go.dev/pkg/io/#Reader) that wraps another `io.Reader`, modifying the stream in some way.

For example, the [gzip.NewReader](https://go.dev/pkg/compress/gzip/#NewReader) function takes an `io.Reader` (a stream of compressed data) and returns a `*gzip.Reader` that also implements `io.Reader` (a stream of the decompressed data).

Implement a `rot13Reader` that implements `io.Reader` and reads from an `io.Reader`, modifying the stream by applying the [rot13](https://en.wikipedia.org/wiki/ROT13) substitution cipher to all alphabetical characters.

The `rot13Reader` type is provided for you. Make it an `io.Reader` by implementing its `Read` method.





### Images

Package *image* defines the `Image` interface.

```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

The `Rectangle` return value of the `Bounds()` method is actually an `image.Rectangle`

The `color.Color` and `color.Model` types are also interfaces, but we can ignore them by using the predefined implementations `color.RGBA` and `color.RGBAModel`.

