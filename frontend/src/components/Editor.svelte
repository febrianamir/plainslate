<script>
  import { OpenOrCreateFile, SaveFile } from '../../wailsjs/go/usecase/Usecase.js'
  import { openedFile } from '../stores/global.js'
  import { onMount, onDestroy } from 'svelte'
  import { get } from 'svelte/store'

  let {
    filename = $bindable('Untitled'),
    fileContent = $bindable(''),
    savedFileContent = $bindable(''),
    hasUnsavedChanges = $bindable(false),
    placeholder = 'Write your markdown here...',
    rows = 35,
    cols = 100,
  } = $props()

  let unsubOpenedFile
  let errorMessage = ''
  let successMessage = ''

  function handleInput(e) {
    fileContent = e.target.value
    hasUnsavedChanges = fileContent !== savedFileContent
    dispatchEvent(new CustomEvent('input', { detail: fileContent }))
  }

  async function saveFile() {
    // Only save changed files
    if (hasUnsavedChanges) {
      errorMessage = ''
      successMessage = ''
      try {
        let saveFilePath = get(openedFile)
        await SaveFile(saveFilePath, fileContent)
        successMessage = 'File saved!'
        savedFileContent = fileContent
        hasUnsavedChanges = false
      } catch (err) {
        errorMessage = err.message || 'Failed to save file.'
      }
    }
  }

  async function onKeyDown(e) {
    if ((e.ctrlKey || e.metaKey) && e.key === 's') {
      // Ctrl+S or Command+S for MacOS
      e.preventDefault()
      await saveFile()
    }
  }

  onMount(() => {
    unsubOpenedFile = openedFile.subscribe(async (path) => {
      if (path && path.trim() !== '') {
        // Open file
        try {
          const result = await OpenOrCreateFile(path)
          fileContent = result
          savedFileContent = result
          hasUnsavedChanges = false
        } catch (err) {
          console.error('Error fetching file:', err)
          fileContent = ''
        }

        // Set tab filename
        filename = path.split('/').pop()
      }
    })
    window.addEventListener('keydown', onKeyDown)
  })

  onDestroy(() => {
    unsubOpenedFile?.()
    window.removeEventListener('keydown', onKeyDown)
  })
</script>

<div class="editor">
  <div class="editor-tab">
    <div class="editor-tab-item">
      <div class="editor-tab-text">
        {filename}
      </div>
      <div class="editor-tab-indicator" class:active={hasUnsavedChanges}></div>
    </div>
  </div>
  <textarea
    class="editor-textarea"
    bind:value={fileContent}
    rows={rows}
    cols={cols}
    placeholder={placeholder}
    oninput={handleInput}
  ></textarea>
</div>

<style>
  .editor {
    display: flex;
    flex-direction: column;
    height: 100vh;
    overflow: hidden;
    flex: 1;
  }

  .editor-tab {
    background-color: #282828;
    border-bottom: 2px solid #1e1e1e;
  }

  .editor-tab-item {
    display: inline-flex;
    align-items: center;
    gap: 0.4rem;
    position: relative;
    cursor: pointer;
    font-size: 0.85rem;
    padding: 0.5rem 0.4rem 0.5rem 0.7rem;
    color: #9d9d9d;
  }

  .editor-tab-indicator {
    margin-top: 1px;
    width: 5px;
    height: 5px;
    border-radius: 5px;
  }

  .editor-tab-indicator.active {
    background-color: #c9e6c1;
  }

  .editor-tab-item::after {
    content: '';
    position: absolute;
    bottom: -1px; /* Align with .tab's black border */
    width: 100%;
    left: 0;
    bottom: 0;
    height: 2px;
    background-color: #6da96f;
    border-radius: 1px;
  }

  .editor-textarea {
    all: unset; /* Resets most browser default styles */
    flex: 1;
    overflow: auto;
    resize: none;
    font-size: 0.85rem;
    font-family: 'JetBrains Mono', monospace;
    white-space: pre-wrap; /* Preserve newlines */
    overflow-wrap: break-word;
    padding: 2rem;
    background: #282828;
    color: #cccccc;
    outline: none;
    box-shadow: none;
  }
</style>
