import {
  openedFilesOpen,
  openedFilesCheckExist,
  openedFilesSelect,
  openedFilesClose,
} from '../../state/openedFile.svelte'
import { OpenOrCreateFile } from '../../../wailsjs/go/usecase/Usecase'

export async function openFile(filepath) {
  if (!filepath || filepath.trim() === '') {
    return
  }

  if (openedFilesCheckExist(filepath)) {
    return openedFilesSelect(filepath)
  }

  try {
    const req = {
      filePath: filepath,
    }
    const resp = await OpenOrCreateFile(req)
    openedFilesOpen({
      id: filepath,
      filepath: filepath,
      filename: filepath.split('/').pop(),
      fileContent: resp,
      savedFileContent: resp,
      hasUnsavedChanges: false,
    })
  } catch (err) {
    console.error('Error open file:', err)
  }
}

export function closeFile(filepath) {
  if (!filepath || filepath.trim() === '') {
    return
  }

  if (!openedFilesCheckExist(filepath)) {
    return
  }

  openedFilesClose(filepath)
}
