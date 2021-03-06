package bitcask

var DefaultConfig = &Config{
	MaxFileSize:  2 << 10,
	MaxKeySize:   2 << 5,
	MaxValueSize: 2 << 6,
	Sync:         false,
}

type Config struct {
	MaxFileSize  int64  // 每个文件最大值
	MaxKeySize   uint32 // key最大值
	MaxValueSize uint64 // value最大值
	Sync         bool   // 是否强制落盘
}
