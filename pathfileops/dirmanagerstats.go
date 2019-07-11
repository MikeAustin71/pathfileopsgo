package pathfileops

type DirTreeCopyStats struct {
  TotalDirsProcessed  uint64
  DirsCopied          uint64
  DirsCreated         uint64
  TotalFilesProcessed uint64
  FilesCopied         uint64
  FilesNotCopied      uint64
  ComputeError        error
}

type DirectoryCopyStats struct {
  DirCreated          uint64
  TotalFilesProcessed uint64
  FilesCopied         uint64
  FilesNotCopied      uint64
  ComputeError        error
}

type DirectoryMoveStats struct {
  TotalSrcFilesProcessed uint64
  SourceFilesMoved       uint64
  SourceFilesRemaining   uint64
  TotalDirsProcessed     uint64
  DirCreated             uint64
  NumOfSubDirectories    uint64
  SourceDirWasDeleted    bool
  ComputeError           error
}

type DeleteDirFilesStats struct {
  TotalFilesProcessed  uint64
  FilesDeleted         uint64
  FilesRemaining       uint64
  TotalDirsProcessed   uint64
  TotalSubDirectories  uint64
  TotalDirsScanned     uint64
  DirectoriesDeleted   uint64
  DirectoriesRemaining uint64
}
