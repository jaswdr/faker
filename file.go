package faker

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var (
	extensions = []string{"ods", "xls", "xlsx", "csv", "ics", "vcf", "3dm", "3ds", "max", "bmp", "dds", "gif", "jpg", "jpeg", "png", "psd", "xcf", "tga", "thm", "tif", "tiff", "yuv", "ai", "eps", "ps", "svg", "dwg", "dxf", "gpx", "kml", "kmz", "webp", "3g2", "3gp", "aaf", "asf", "avchd", "avi", "drc", "flv", "m2v", "m4p", "m4v", "mkv", "mng", "mov", "mp2", "mp4", "mpe", "mpeg", "mpg", "mpv", "mxf", "nsv", "ogg", "ogv", "ogm", "qt", "rm", "rmvb", "roq", "srt", "svi", "vob", "webm", "wmv", "yuv", "aac", "aiff", "ape", "au", "flac", "gsm", "it", "m3u", "m4a", "mid", "mod", "mp3", "mpa", "pls", "ra", "s3m", "sid", "wav", "wma", "xm", "7z", "a", "apk", "ar", "bz2", "cab", "cpio", "deb", "dmg", "egg", "gz", "iso", "jar", "lha", "mar", "pea", "rar", "rpm", "s7z", "shar", "tar", "tbz2", "tgz", "tlz", "war", "whl", "xpi", "zip", "zipx", "xz", "pak", "exe", "msi", "bin", "command", "sh", "bat", "crx", "c", "cc", "class", "clj", "cpp", "cs", "cxx", "el", "go", "h", "java", "lua", "m", "m4", "php", "pl", "po", "py", "rb", "rs", "sh", "swift", "vb", "vcxproj", "xcodeproj", "xml", "diff", "patch", "html", "js", "html", "htm", "css", "js", "jsx", "less", "scss", "wasm", "php", "eot", "otf", "ttf", "woff", "woff2", "ppt", "odp", "doc", "docx", "ebook", "log", "md", "msg", "odt", "org", "pages", "pdf", "rtf", "rst", "tex", "txt", "wpd", "wps", "mobi", "epub", "azw1", "azw3", "azw4", "azw6", "azw", "cbr", "cbz"}
)

// File is a faker struct for File
type File struct {
	Faker *Faker
}

// Extension returns a fake Extension file
func (f File) Extension() string {
	return f.Faker.RandomStringElement(extensions)
}

// FilenameWithExtension returns a fake file name with extension
func (f File) FilenameWithExtension() string {
	extension := f.Faker.RandomStringElement(extensions)
	text := f.Faker.Lorem().Word()

	return fmt.Sprintf("%s.%s", text, extension)
}

// AbsoluteUnixPath returns a fake absolute unix path to a file
func (f File) AbsoluteUnixPath(nbLen int) string {
	return filepath.Join(f.Directory(nbLen), f.FilenameWithExtension())
}

// AbsoluteWin32Path returns a fake absolute win32 path to a file
func (f File) AbsoluteWin32Path(nbLen int) string {
	return filepath.Join(f.DriveLetter(), f.Directory(nbLen), f.FilenameWithExtension())
}

// Directory returns a fake directory (the directory path style OS dependent)
func (f File) Directory(nbLen int) string {
	if runtime.GOOS == "windows" {
		return f.DriveLetter() + filepath.Join(f.Faker.Lorem().Words(nbLen)...)
	} else {
		return "/" + filepath.Join(f.Faker.Lorem().Words(nbLen)...)
	}
}

// DriveLetter returns a fake Win32 drive letter
func (f File) DriveLetter() string {
	return f.Faker.RandomLetter() + ":\\"
}
