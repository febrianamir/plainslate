let openedFiles = $state({
  files: [],
  activeId: null,
})

export function getOpenedFiles() {
  return openedFiles
}

export function getActiveFile() {
  return openedFiles.files.find((f) => {
    return f.id === openedFiles.activeId
  })
}

export function openedFilesOpen(file) {
  // Prevent duplicates
  if (!isFileExist(file.id)) {
    openedFiles.files.push(file)
  }
  openedFiles.activeId = file.id
}

export function openedFilesClose(fileId) {
  openedFiles.files = openedFiles.files.filter((f) => {
    return f.id !== fileId
  })
  // If the closed file was active, switch to another or null
  if (openedFiles.activeId === fileId) {
    openedFiles.activeId = openedFiles.files.length ? openedFiles.files[0].id : null
  }
}

export function openedFilesSelect(fileId) {
  if (isFileExist(fileId)) {
    openedFiles.activeId = fileId
  }
}

export function openedFilesUpdateFile(fileId, updatedData) {
  let file = openedFiles.files.find((f) => {
    return f.id === fileId
  })
  if (file) {
    Object.assign(file, updatedData)
  }
}

function isFileExist(fileId) {
  return openedFiles.files.find((f) => {
    return f.id === fileId
  })
}
