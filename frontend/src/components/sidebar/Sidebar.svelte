<script>
  import Tree from './tree/Tree.svelte'
  import SetRootPath from '../SetRootPath.svelte'
  import { FileText, Search } from 'lucide-svelte'

  let width = 250
  let isResizing = false
  let activeTab = 'explorer'

  const startResize = (e) => {
    isResizing = true
    e.preventDefault()
  }

  const onMouseMove = (e) => {
    if (isResizing) {
      width = Math.max(200, e.clientX) // Minimum width = 200px
    }
  }

  const handleActiveTab = (tabName) => {
    activeTab = tabName
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
    <div class="sidebar-activity">
      <div class="sidebar-activity-bar">
        <div
          class="sidebar-explorer-button {activeTab === 'explorer' ? 'active' : ''}"
          on:click={(e) => {
            e.preventDefault()
            handleActiveTab('explorer')
          }}
          on:keydown={(e) => {
            if (e.key === 'Enter') {
              e.preventDefault()
              handleActiveTab('explorer')
            }
          }}
        >
          <FileText size={24} />
        </div>
        <div
          class="sidebar-search-button {activeTab === 'search' ? 'active' : ''}"
          on:click={(e) => {
            e.preventDefault()
            handleActiveTab('search')
          }}
          on:keydown={(e) => {
            if (e.key === 'Enter') {
              e.preventDefault()
              handleActiveTab('search')
            }
          }}
        >
          <Search size={24} strokeWidth={2.5} />
        </div>
      </div>
      <div class="sidebar-activity-content">
        <div class="sidebar-explorer {activeTab === 'explorer' ? 'active' : ''}">
          <Tree />
        </div>
        <div class="sidebar-search {activeTab === 'search' ? 'active' : ''}"></div>
      </div>
    </div>
  </div>
  <div class="sidebar-resizer" class:active={isResizing} on:mousedown={startResize}></div>
</div>

<style>
  .sidebar {
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
    overflow: hidden;
    padding-top: 1rem;
  }

  .sidebar-appname {
    padding: 0 1rem 1rem 1rem;
  }

  .sidebar-activity {
    display: flex;
  }

  .sidebar-activity-bar {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.75rem;
    width: 50px;
    padding: 0.4rem 0;
  }

  .sidebar-explorer-button,
  .sidebar-search-button {
    cursor: pointer;
  }

  .sidebar-explorer-button.active,
  .sidebar-search-button.active {
    color: #6da96f;
  }

  .sidebar-activity-content {
    flex: 1;
  }

  .sidebar-explorer,
  .sidebar-search {
    display: none;
  }

  .sidebar-explorer.active,
  .sidebar-search.active {
    display: block;
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
