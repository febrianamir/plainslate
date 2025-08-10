<script>
  import Search from './search/Search.svelte'
  import Tree from './tree/Tree.svelte'
  import SetRootPath from '../SetRootPath.svelte'
  import { FileText, Search as SearchIcon } from 'lucide-svelte'
  import { onDestroy, onMount } from 'svelte'
  import { handleEnter } from '../../../src/lib/utils.js'

  let sidebarWidth = 250
  let isSidebarResizing = false
  let activeTab = 'explorer'

  function startResize(e) {
    e.preventDefault()
    isSidebarResizing = true
  }

  function onMouseMove(e) {
    if (isSidebarResizing) {
      sidebarWidth = Math.max(200, e.clientX) // Minimum width = 200px
    }
  }

  function handleActiveTab(tabName) {
    activeTab = tabName
  }

  function stopResize() {
    isSidebarResizing = false
  }

  onMount(() => {
    window.addEventListener('mousemove', onMouseMove)
    window.addEventListener('mouseup', stopResize)
  })

  onDestroy(() => {
    window.removeEventListener('mousemove', onMouseMove)
    window.removeEventListener('mouseup', stopResize)
  })
</script>

<div class="sidebar">
  <div class="sidebar-content" style="width: {sidebarWidth}px">
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
            handleEnter(e, () => handleActiveTab('explorer'))
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
            handleEnter(e, () => handleActiveTab('search'))
          }}
        >
          <SearchIcon size={24} strokeWidth={2.5} />
        </div>
      </div>
      <div class="sidebar-activity-content">
        <div class="sidebar-explorer {activeTab === 'explorer' ? 'active' : ''}">
          <Tree />
        </div>
        <div class="sidebar-search {activeTab === 'search' ? 'active' : ''}">
          <Search />
        </div>
      </div>
    </div>
  </div>
  <div class="sidebar-resizer" class:active={isSidebarResizing} on:mousedown={startResize}></div>
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
    height: 100%;
    overflow: hidden;
    padding-top: 1rem;
  }

  .sidebar-appname {
    padding: 0 1rem 1rem 1rem;
  }

  .sidebar-activity {
    display: flex;
    flex: 1;
    min-height: 0;
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
    display: flex;
    flex-direction: column;
    flex: 1;
    min-width: 0;
  }

  .sidebar-explorer,
  .sidebar-search {
    display: none;
    flex: 1;
    min-height: 0;
  }

  .sidebar-explorer.active,
  .sidebar-search.active {
    display: flex;
    flex-direction: column;
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
