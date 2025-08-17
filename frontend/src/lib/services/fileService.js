import {
  openedFilesOpen,
  openedFilesCheckExist,
  openedFilesSelect,
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
    const result = await OpenOrCreateFile(filepath)
    openedFilesOpen({
      id: filepath,
      filepath: filepath,
      filename: filepath.split('/').pop(),
      fileContent: result,
      savedFileContent: result,
      hasUnsavedChanges: false,
    })
  } catch (err) {
    console.error('Error open file:', err)
  }
}
