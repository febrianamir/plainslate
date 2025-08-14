<script>
  import { getOpenedFiles } from '../../state/openedFile.svelte'
  import Editor from './Editor.svelte'

  let openedFiles = getOpenedFiles()
</script>

<div class="editor-workspace">
  {#if openedFiles.files && openedFiles.files.length > 0}
    <Editor />
  {:else}
    No File Opened
  {/if}
</div>

<style>
  .editor-workspace {
    min-width: 0;
    flex: 1;
  }
</style>
