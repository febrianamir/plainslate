<script>
  import { debounce } from '../../lib/utils.js'
  import { OpenOrCreateFile, SaveFile } from '../../../wailsjs/go/usecase/Usecase.js'
  import {
    getOpenedFiles,
    getActiveFile,
    openedFilesSelect,
  } from '../../state/openedFile.svelte.js'
  import { onMount, onDestroy } from 'svelte'
  import { X } from 'lucide-svelte'
  import { closeFile } from '../../lib/services/fileService.js'

  let openedFiles = getOpenedFiles()
  let rows = 35
  let cols = 100
  let placeholder = 'Write your markdown here...'

  let activeFile = $state({})
  $effect(() => {
    if (openedFiles.activeId) {
      activeFile = getActiveFile()
    } else {
      activeFile = {}
    }
  })

  const debouncedMarkUnsaved = debounce(() => {
    activeFile.hasUnsavedChanges = activeFile.fileContent !== activeFile.savedFileContent
  }, 200)

  function handleInput(e) {
    debouncedMarkUnsaved()
  }

  async function saveFile() {
    // Only save changed files
    if (activeFile.hasUnsavedChanges) {
      try {
        let saveFilePath = activeFile.filepath
        await SaveFile(saveFilePath, activeFile.fileContent)
        activeFile.savedFileContent = activeFile.fileContent
        activeFile.hasUnsavedChanges = false
      } catch (err) {
        console.log('Failed to save file:', err)
      }
    }
  }

  function switchActiveFile(fileId) {
    openedFilesSelect(fileId)
  }

  async function onKeyDown(e) {
    if ((e.ctrlKey || e.metaKey) && e.key === 's') {
      // Ctrl+S or Command+S for MacOS
      e.preventDefault()
      await saveFile()
    }
  }

  onMount(() => {
    window.addEventListener('keydown', onKeyDown)
  })

  onDestroy(() => {
    window.removeEventListener('keydown', onKeyDown)
  })
</script>

<div class="editor">
  <div class="editor-tab">
    {#each openedFiles.files as file}
      <div
        role="button"
        tabindex="0"
        class:active={file.id === openedFiles.activeId}
        class="editor-tab-item"
        onclick={(e) => switchActiveFile(file.id)}
        onkeydown={(e) => {
          handleEnter(e, () => switchActiveFile(file.id))
        }}
      >
        <div class="editor-tab-text">
          {file.filename}
        </div>
        <div class="editor-tab-actions">
          <div class="editor-tab-indicator" class:active={file.hasUnsavedChanges}></div>
          <div
            role="button"
            tabindex="0"
            class="editor-tab-close"
            onclick={(e) => {
              e.stopPropagation()
              closeFile(file.id)
            }}
            onkeydown={(e) => {
              handleEnter(e, () => {
                e.stopPropagation()
                closeFile(file.id)
              })
            }}
          >
            <X size={15} />
          </div>
        </div>
      </div>
    {/each}
  </div>
  <textarea
    class="editor-textarea"
    bind:value={activeFile.fileContent}
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
  }

  .editor-tab {
    background-color: #282828;
    border-bottom: 2px solid #1e1e1e;
  }

  .editor-tab-item {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    position: relative;
    cursor: pointer;
    font-size: 0.85rem;
    padding: 0.6rem 0.65rem 0.6rem 0.7rem;
    color: #9d9d9d;
  }

  .editor-tab-item.active::after {
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

  .editor-tab-actions {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 15px;
    height: 15px;
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

  .editor-tab-actions:hover .editor-tab-indicator {
    opacity: 0;
  }

  .editor-tab-close {
    position: absolute;
    top: 0;
    left: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 16px;
    height: 16px;
    border-radius: 3px;
    background-color: #1e1e1e;
    color: #9d9d9d;
    opacity: 0;
    cursor: pointer;
  }

  .editor-tab-actions:hover .editor-tab-close {
    opacity: 1;
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
