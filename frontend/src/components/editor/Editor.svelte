<script>
  import { debounce } from '../../lib/utils.js'
  import { OpenOrCreateFile, SaveFile } from '../../../wailsjs/go/usecase/Usecase.js'
  import {
    getOpenedFiles,
    getActiveFile,
    openedFilesSelect,
  } from '../../state/openedFile.svelte.js'
  import { onMount, onDestroy } from 'svelte'
  import { X, Eye, EyeOff } from 'lucide-svelte'
  import { closeFile } from '../../lib/services/fileService.js'
  import CodeMirrorEditor from './CodeMirrorEditor.svelte'
  import MarkdownPreview from './MarkdownPreview.svelte'

  let openedFiles = getOpenedFiles()
  let rows = 35
  let cols = 100
  let placeholder = 'Write your markdown here...'
  let isShowPreview = $state(false)

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

  function togglePreview() {
    isShowPreview = !isShowPreview
  }

  async function saveFile() {
    // Only save changed files
    if (activeFile.hasUnsavedChanges) {
      try {
        let saveFilePath = activeFile.filePath
        const req = {
          filePath: saveFilePath,
          content: activeFile.fileContent,
        }
        await SaveFile(req)
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
      // Ctrl/Cmd+S to save
      e.preventDefault()
      await saveFile()
    }

    // Ctrl/Cmd+Alt+P to toggle preview
    if ((e.ctrlKey || e.metaKey) && e.altKey && e.key === 'p') {
      e.preventDefault()
      togglePreview()
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
  <div class="editor-header">
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
            {file.fileName}
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

    <div class="preview-controls">
      <button
        class="control-btn"
        class:active={isShowPreview}
        onclick={togglePreview}
        title="Toggle Preview (Ctrl+Alt+P)"
      >
        {#if isShowPreview}
          <EyeOff size={16} />
        {:else}
          <Eye size={16} />
        {/if}
      </button>
    </div>
  </div>
  <div class="editor-content">
    <!-- Editor panel -->
    <div class="editor-panel" class:hidden={isShowPreview}>
      <CodeMirrorEditor
        value={activeFile.fileContent || ''}
        placeholder={placeholder}
        onSave={saveFile}
        onChange={(newValue) => {
          activeFile.fileContent = newValue
          handleInput()
        }}
      />
    </div>

    <!-- Preview panel -->
    <div class="preview-panel" class:active={isShowPreview}>
      <MarkdownPreview fileContent={activeFile.fileContent || ''} className="editor-preview" />
    </div>
  </div>
</div>

<style>
  .editor {
    display: flex;
    flex-direction: column;
    height: 100vh;
    overflow: hidden;
  }

  .editor-header {
    display: flex;
    justify-content: space-between;
    background-color: #282828;
    border-bottom: 2px solid #1e1e1e;
  }

  .editor-tab {
    background-color: #282828;
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

  /* Preview controls */
  .preview-controls {
    padding: 0.125rem 0.7rem;
  }

  .control-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 32px;
    height: 32px;
    border: none;
    border-radius: 4px;
    background: #1e1e1e;
    color: #9d9d9d;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  /* Content layout */
  .editor-content {
    flex: 1;
    display: flex;
    overflow: hidden;
  }

  .editor-panel {
    flex: 1;
    overflow: scroll;
  }

  .editor-panel.hidden {
    display: none;
  }

  .preview-panel {
    display: none;
    width: 100%;
  }

  .preview-panel.active {
    display: block;
  }
</style>
