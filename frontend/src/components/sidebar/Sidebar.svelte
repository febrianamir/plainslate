<script>
  import DirectoryTree from '../DirectoryTree.svelte'

  let width = 250
  let isResizing = false

  const startResize = e => {
    isResizing = true
    e.preventDefault()
  }

  const onMouseMove = e => {
    if (isResizing) {
      width = Math.max(200, e.clientX) // Minimum width = 200px
    }
  }

  const stopResize = () => {
    isResizing = false
  }

  // Register global mouse events
  window.addEventListener('mousemove', onMouseMove)
  window.addEventListener('mouseup', stopResize)
</script>

<div class="sidebar">
  <div class="sidebar-content" style="width: {width}px">
    <DirectoryTree />
  </div>
  <div class="sidebar-resizer" on:mousedown={startResize}></div>
</div>

<style>
  .sidebar {
    display: flex;
    height: 100%;
    background-color: #1e1e1e;
    color: #9d9d9d;
  }

  .sidebar-content {
    padding-top: 1rem;
  }

  .sidebar-resizer {
    width: 5px;
    cursor: ew-resize;
    height: 100%;
    flex-shrink: 0;
    background-color: #1e1e1e;
  }
</style>
