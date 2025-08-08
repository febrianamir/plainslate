<script>
  import Tree from './tree/Tree.svelte'
  import SetRootPath from '../SetRootPath.svelte'

  let width = 250
  let isResizing = false

  const startResize = (e) => {
    isResizing = true
    e.preventDefault()
  }

  const onMouseMove = (e) => {
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
    <p class="sidebar-appname">PlainSlate</p>
    <SetRootPath />
    <Tree />
  </div>
  <div class="sidebar-resizer" class:active={isResizing} on:mousedown={startResize}></div>
</div>

<style>
  .sidebar {
    box-sizing: border-box;
    overflow: hidden;
    display: flex;
    height: 100%;
    background-color: #1e1e1e;
    color: #9d9d9d;
  }

  .sidebar-content {
    display: flex;
    flex-direction: column;
    max-width: 100%;
    width: 100%;
    box-sizing: border-box;
    overflow: hidden;
    padding-top: 1rem;
  }

  .sidebar-appname {
    padding: 0 1rem 1rem 1rem;
  }

  .sidebar-resizer {
    width: 2px;
    cursor: ew-resize;
    height: 100%;
    flex-shrink: 0;
    background-color: #1e1e1e;
  }

  .sidebar-resizer:hover,
  .sidebar-resizer.active {
    background-color: #9d9d9d;
  }
</style>
