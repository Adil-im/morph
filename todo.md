### Notes

Detecting input format automatically
Don't trust the file extension alone — check the actual file content via magic bytes (the first few bytes of a file are a signature). Go makes this easy:

```
func detectFormat(data []byte) string {
    switch {
    case bytes.HasPrefix(data, []byte{0xFF, 0xD8, 0xFF}):
        return "jpeg"
    case bytes.HasPrefix(data, []byte{0x89, 0x50, 0x4E, 0x47}):
        return "png"
    case bytes.HasPrefix(data, []byte("GIF87a")) || bytes.HasPrefix(data, []byte("GIF89a")):
        return "gif"
    case bytes.HasPrefix(data, []byte("RIFF")) && bytes.Contains(data[8:12], []byte("WEBP")):
        return "webp"
    case bytes.HasPrefix(data, []byte{0x42, 0x4D}):
        return "bmp"
    default:
        return ""
    }
}
```

Even simpler: Go's stdlib already does this for you for registered formats. image.DecodeConfig(reader) returns the format name as a string once you've imported the relevant packages (_ "image/png", _ "image/jpeg", etc. for their side-effect registration). So you often don't need to hand-roll magic byte detection at all — just read a few bytes, call image.DecodeConfig, and it tells you what it found.

Some formats (like plain text-based ones, or corrupted files) won't have clean magic bytes. In that case, fall back to the file extension, but still warn the user: "could not verify format from content, assuming .png from filename". This avoids silent misconversions.
