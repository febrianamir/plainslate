<script>
  import { OpenFile } from '../../wailsjs/go/usecase/Usecase.js';
  import { openedFile } from '../stores/global.js';
  import { createEventDispatcher, onMount, onDestroy } from 'svelte';

  export let value = '# Hello\nThis is a markdown note.';
  export let placeholder = 'Write your markdown here...';
  export let rows = 25;
  export let cols = 100;
  let unsubOpenedFile;

  onMount(() => {
    unsubOpenedFile = openedFile.subscribe(async (path) => {
      if (path && path.trim() !== "") {
        try {
          const result = await OpenFile(path)
          value = result
        } catch (err) {
          console.error("Error fetching file:", err)
          value = ''
        }
      }
    })
  })

  onDestroy(() => {
    unsubOpenedFile?.()
  })

  const handleInput = (event) => {
    value = event.target.value;
    dispatchEvent(new CustomEvent('input', { detail: value }));
  };
</script>

<textarea
  class="editor-textarea"
  bind:value
  rows={rows}
  cols={cols}
  placeholder={placeholder}
  on:input={handleInput}
></textarea>

<style>
  .editor-textarea {
    background-color: rgb(34, 34, 34);
    color: rgb(210, 210, 210);
    width: 100%;
    height: 100%;
    min-height: 400px;
    font-family: monospace;
    font-size: 1rem;
    padding: 1rem;
    box-sizing: border-box;
    resize: vertical;
  }
</style>