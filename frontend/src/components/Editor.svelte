<script>
  import { OpenFile, SaveFile } from '../../wailsjs/go/usecase/Usecase.js'
  import { openedFile } from '../stores/global.js'
  import { onMount, onDestroy } from 'svelte'
  import { get } from 'svelte/store'

  export let filename = 'Untitled'
  export let value = '# Hello\nThis is a markdown note.'
  export let placeholder = 'Write your markdown here...'
  export let rows = 35
  export let cols = 100

  let unsubOpenedFile
  let errorMessage = ''
  let successMessage = ''

  const handleInput = e => {
    value = e.target.value
    dispatchEvent(new CustomEvent('input', { detail: value }))
  }

  const handleSave = async () => {
    errorMessage = ''
    successMessage = ''
    try {
      let saveFilePath = get(openedFile)
      await SaveFile(saveFilePath, value)
      successMessage = 'Directory saved!'
    } catch (err) {
      errorMessage = err.message || 'Failed to save directory.'
    }
  }

  const onKeyDown = async e => {
    if ((e.ctrlKey || e.metaKey) && e.key === 's') {
      // Ctrl+S or Command+S for MacOS
      e.preventDefault()
      await handleSave()
    }
  }

  onMount(() => {
    unsubOpenedFile = openedFile.subscribe(async path => {
      if (path && path.trim() !== '') {
        // Open file
        try {
          const result = await OpenFile(path)
          value = result
        } catch (err) {
          console.error('Error fetching file:', err)
          value = ''
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
      {filename}
    </div>
  </div>
  <textarea
    class="editor-textarea"
    bind:value={value}
    rows={rows}
    cols={cols}
    placeholder={placeholder}
    on:input={handleInput}
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
    display: inline-block;
    position: relative;
    cursor: pointer;
    font-size: 0.85rem;
    padding: 0.5rem 0.75rem;
    color: #9d9d9d;
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
