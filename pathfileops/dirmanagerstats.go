package pathfileops

type DirTreeCopyStats struct {
  TotalDirsProcessed  uint64
  DirsCopied          uint64
  TotalFilesProcessed uint64
  FilesCopied         uint64
  FilesNotCopied      uint64
  ComputeError        error
}

type DirectoryCopyStats struct {
  DirsCopied          uint64
  TotalFilesProcessed uint64
  FilesCopied         uint64
  FilesNotCopied      uint64
  ComputeError        error
}
