<script>
  export let contextMenuX = 0
  export let contextMenuY = 0

  export let handleOpenCreateFileInput = () => {}
  export let handleOpenCreateDirectoryInput = () => {}
</script>

<div class="tree-context-menu" style="top:{contextMenuY}px; left:{contextMenuX}px;">
  <div
    role="button"
    tabindex="0"
    class="tree-context-item"
    on:click={handleOpenCreateFileInput}
    on:keydown={(e) => {
      if (e.key === 'Enter' || e.key === ' ') {
        e.preventDefault()
        handleOpenCreateFileInput()
      }
    }}
  >
    New File
  </div>
  <div
    role="button"
    tabindex="0"
    class="tree-context-item"
    on:click={handleOpenCreateDirectoryInput}
    on:keydown={(e) => {
      if (e.key === 'Enter' || e.key === ' ') {
        e.preventDefault()
        handleOpenCreateDirectoryInput()
      }
    }}
  >
    New Directory
  </div>
</div>

<style>
  .tree-context-menu {
    position: fixed;
    z-index: 1000;
    background: #282828;
    border: 1px solid #323232;
    border-radius: 2px;
    font-size: 0.8rem;
    padding: 0.2rem 0;
  }

  .tree-context-item {
    padding: 0.35rem 1.2rem;
    cursor: pointer;
  }

  .tree-context-item:hover {
    color: #282828;
    background-color: #c9e6c1;
  }
</style>
