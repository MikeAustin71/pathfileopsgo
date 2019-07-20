package pathfileops

type DirectoryStatsDto struct {
  numOfFiles    uint64
  numOfSubDirs  uint64
  numOfBytes    uint64
  isInitialized bool
}

func (dirStats *DirectoryStatsDto) IsInitialized() bool {
  return dirStats.isInitialized
}

func (dirStats *DirectoryStatsDto) NumOfFiles() uint64 {
  return dirStats.numOfFiles
}

func (dirStats *DirectoryStatsDto) NumOfSubDirs() uint64 {
  return dirStats.numOfSubDirs
}

func (dirStats *DirectoryStatsDto) NumOfBytes() uint64 {
  return dirStats.numOfBytes
}

type DirTreeCopyStats struct {
  TotalDirsScanned    uint64
  DirsCopied          uint64
  DirsCreated         uint64
  TotalFilesProcessed uint64
  FilesCopied         uint64
  FileBytesCopied     uint64
  FilesNotCopied      uint64
  FileBytesNotCopied  uint64
  ComputeError        error
}

type DirectoryCopyStats struct {
  DirCreated          uint64
  TotalFilesProcessed uint64
  FilesCopied         uint64
  FileBytesCopied     uint64
  FilesNotCopied      uint64
  FileBytesNotCopied  uint64
  ComputeError        error
}

type DirectoryMoveStats struct {
  TotalSrcFilesProcessed   uint64
  SourceFilesMoved         uint64
  SourceFileBytesMoved     uint64
  SourceFilesRemaining     uint64
  SourceFileBytesRemaining uint64
  TotalDirsProcessed       uint64
  DirsCreated              uint64
  NumOfSubDirectories      uint64
  SourceDirWasDeleted      bool
  ComputeError             error
}

type DeleteDirFilesStats struct {
  TotalFilesProcessed        uint64
  FilesDeleted               uint64
  FilesDeletedBytes          uint64
  FilesRemaining             uint64
  FilesRemainingBytes        uint64
  TotalSubDirectories        uint64
  TotalDirsScanned           uint64
  NumOfDirsWhereFilesDeleted uint64
  DirectoriesDeleted         uint64
}
