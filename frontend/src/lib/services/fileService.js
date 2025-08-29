import {
  openedFilesOpen,
  openedFilesCheckExist,
  openedFilesSelect,
  openedFilesClose,
} from '../../state/openedFile.svelte'
import { OpenOrCreateFile } from '../../../wailsjs/go/usecase/Usecase'

export async function openFile(filePath) {
  if (!filePath || filePath.trim() === '') {
    return
  }

  if (openedFilesCheckExist(filePath)) {
    return openedFilesSelect(filePath)
  }

  try {
    const req = {
      filePath: filePath,
    }
    const resp = await OpenOrCreateFile(req)
    openedFilesOpen({
      id: filePath,
      filePath: filePath,
      fileName: filePath.split('/').pop(),
      fileContent: resp,
      savedFileContent: resp,
      hasUnsavedChanges: false,
    })
  } catch (err) {
    console.error('Error open file:', err)
  }
}

export function closeFile(filePath) {
  if (!filePath || filePath.trim() === '') {
    return
  }

  if (!openedFilesCheckExist(filePath)) {
    return
  }

  openedFilesClose(filePath)
}
