package classpath
import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)
type Entry interface {
  readClass(className string) ([]byte, Entry, error)
  String() string
}

func newEntry(path string) Entry {
  if strings.Contains(path, pathListSeparator) {
    return newCompositeEntry(path)
  }
  if strings.Contains(path, "*") {
    return newWildcardEntry(path)
  }
  if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") {
    return newZipEntry(path)
  }
  return newDirEntry(path)
}


