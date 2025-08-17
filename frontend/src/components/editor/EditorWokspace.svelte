<script>
  import { getOpenedFiles } from '../../state/openedFile.svelte.js'
  import appicon from '../../assets/images/appicon.png'
  import Editor from './Editor.svelte'

  let openedFiles = getOpenedFiles()
</script>

<div class="editor-workspace">
  {#if openedFiles.files && openedFiles.files.length > 0}
    <Editor />
  {:else}
    <div class="editor-placeholder">
      <img src={appicon} alt="PlainSlate" class="editor-placeholder-image" />
    </div>
  {/if}
</div>

<style>
  .editor-workspace {
    min-width: 0;
    flex: 1;
  }

  .editor-placeholder {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }

  .editor-placeholder-image {
    display: block;
    width: 150px;
    height: 150px;
    user-select: none;
    -webkit-user-drag: none;
  }
</style>
