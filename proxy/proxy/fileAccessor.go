package proxy

type FileAccessor interface {
	ReadFile(filePath string)
}
