<script>
  import { handleEnter } from '../../../../src/lib/utils.js'

  let {
    contextMenuX = 0,
    contextMenuY = 0,
    showCreateFileInput = () => {},
    showCreateFolderInput = () => {},
    showRenameInput = () => {},
    moveItemToTrash = () => {},
    addToClipboard = () => {},
    paste = () => {},
  } = $props()
</script>

<div class="context-menu" style="top:{contextMenuY}px; left:{contextMenuX}px;">
  <div
    role="button"
    tabindex="0"
    class="context-item"
    onclick={showCreateFileInput}
    onkeydown={(e) => {
      handleEnter(e, showCreateFileInput)
    }}
  >
    New File
  </div>
  <div
    role="button"
    tabindex="0"
    class="context-item"
    onclick={showCreateFolderInput}
    onkeydown={(e) => {
      handleEnter(e, showCreateFolderInput)
    }}
  >
    New Directory
  </div>
  <div class="context-separator"></div>
  <div
    role="button"
    tabindex="0"
    class="context-item"
    onclick={showRenameInput}
    onkeydown={(e) => {
      handleEnter(e, showRenameInput)
    }}
  >
    Rename
  </div>
  <div
    role="button"
    tabindex="0"
    class="context-item"
    onclick={moveItemToTrash}
    onkeydown={(e) => {
      handleEnter(e, moveItemToTrash)
    }}
  >
    Move To Trash
  </div>
  <div class="context-separator"></div>
  <div
    role="button"
    tabindex="0"
    class="context-item"
    onclick={(e) => addToClipboard('CUT')}
    onkeydown={(e) => {
      handleEnter(e, addToClipboard('CUT'))
    }}
  >
    Cut
  </div>
  <div
    role="button"
    tabindex="0"
    class="context-item"
    onclick={(e) => paste()}
    onkeydown={(e) => {
      handleEnter(e, paste())
    }}
  >
    Paste
  </div>
</div>

<style>
  .context-menu {
    position: fixed;
    z-index: 10000;
    background: #282828;
    border: 1px solid #323232;
    border-radius: 2px;
    font-size: 0.8rem;
    padding: 0.2rem 0;
  }

  .context-item {
    padding: 0.35rem 1.2rem;
    cursor: pointer;
  }

  .context-item:hover {
    color: #282828;
    background-color: #c9e6c1;
  }

  .context-separator {
    height: 1px;
    background-color: #1e1e1e;
  }
</style>
