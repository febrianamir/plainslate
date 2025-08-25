<script>
  import Search from './search/Search.svelte'
  import Tree from './tree/Tree.svelte'
  import SetRootPath from '../SetRootPath.svelte'
  import { FileText, Search as SearchIcon } from 'lucide-svelte'
  import { onDestroy, onMount } from 'svelte'
  import { handleEnter } from '../../../src/lib/utils.js'

  let isShowSidebar = $state(true)
  let sidebarWidth = $state(250)
  let isSidebarResizing = $state(false)
  let activeTab = $state('explorer')

  function startResize(e) {
    e.preventDefault()
    isSidebarResizing = true
  }

  function onMouseMove(e) {
    if (isSidebarResizing) {
      sidebarWidth = Math.max(200, e.clientX) // Minimum width = 200px
    }
  }

  function toggleSidebar() {
    isShowSidebar = !isShowSidebar
  }

  function handleActiveTab(tabName) {
    activeTab = tabName
  }

  function stopResize() {
    isSidebarResizing = false
  }

  function onKeyDown(e) {
    if ((e.ctrlKey || e.metaKey) && e.key === 'b') {
      e.preventDefault()
      toggleSidebar()
    }
  }

  onMount(() => {
    window.addEventListener('mousemove', onMouseMove)
    window.addEventListener('mouseup', stopResize)
    window.addEventListener('keydown', onKeyDown)
  })

  onDestroy(() => {
    window.removeEventListener('mousemove', onMouseMove)
    window.removeEventListener('mouseup', stopResize)
    window.removeEventListener('keydown', onKeyDown)
  })
</script>

<div class:active={isShowSidebar} class="sidebar">
  <div class="sidebar-content" style="width: {sidebarWidth}px">
    <p class="sidebar-appname">PlainSlate</p>
    <SetRootPath />
    <div class="sidebar-activity">
      <div class="sidebar-activity-bar">
        <div
          class="sidebar-explorer-button {activeTab === 'explorer' ? 'active' : ''}"
          role="button"
          tabindex="0"
          onclick={(e) => {
            e.preventDefault()
            handleActiveTab('explorer')
          }}
          onkeydown={(e) => {
            handleEnter(e, () => handleActiveTab('explorer'))
          }}
        >
          <FileText size={24} />
        </div>
        <div
          class="sidebar-search-button {activeTab === 'search' ? 'active' : ''}"
          role="button"
          tabindex="0"
          onclick={(e) => {
            e.preventDefault()
            handleActiveTab('search')
          }}
          onkeydown={(e) => {
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
  <div
    class="sidebar-resizer"
    role="button"
    tabindex="0"
    class:active={isSidebarResizing}
    onmousedown={startResize}
  ></div>
</div>

<style>
  .sidebar {
    overflow: hidden;
    display: none;
    height: 100%;
    background-color: #1e1e1e;
    color: #9d9d9d;
  }

  .sidebar.active {
    display: flex;
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
