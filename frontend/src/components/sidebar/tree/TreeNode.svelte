<script>
  import { Folder, FolderOpen, File } from 'lucide-svelte'
  import TreeNode from './TreeNode.svelte'
  import { openedFile } from '../../../stores/global.js'
  import { onMount, onDestroy } from 'svelte'

  export let node
  let expanded = node.is_root

  // Context menu properties
  let showContextMenu = false
  let contextMenuX = 0
  let contextMenuY = 0

  onMount(() => {
    window.addEventListener('click', handleCloseContextMenu)
  })

  onDestroy(() => {
    window.removeEventListener('click', handleCloseContextMenu)
  })

  const onClick = () => {
    if (node.type === 'directory') {
      handleToggleExpand()
    }
    if (node.type === 'file') {
      handleOpenFile()
    }
  }

  const onRightClick = e => {
    e.preventDefault()
    handleOpenContextMenu(e)
  }

  const handleToggleExpand = () => {
    expanded = !expanded
  }

  const handleOpenFile = () => {
    openedFile.set(node.path)
  }

  const handleOpenContextMenu = e => {
    contextMenuX = e.clientX
    contextMenuY = e.clientY
    showContextMenu = true
  }

  const handleCloseContextMenu = () => {
    showContextMenu = false
  }

  const handleCreateNewFile = () => {}

  const handleCreateNewDirectory = () => {}
</script>

<div class="node">
  <div
    role="button"
    tabindex="0"
    class="node-content"
    on:click={onClick}
    on:contextmenu={onRightClick}
    on:keydown={e => {
      if (e.key === 'Enter' || e.key === ' ') {
        e.preventDefault()
        onClick()
      }
    }}
  >
    <div class="node-icon">
      {#if node.type === 'directory'}
        {#if expanded}
          <FolderOpen size={16} />
        {:else}
          <Folder size={16} />
        {/if}
      {:else}
        <File size={16} />
      {/if}
    </div>
    <div class="node-text">
      {node.name}
    </div>
  </div>

  {#if showContextMenu}
    <div class="node-context-menu" style="top:{contextMenuY}px; left:{contextMenuX}px;">
      <div
        role="button"
        tabindex="0"
        class="node-context-item"
        on:click={handleCreateNewFile}
        on:keydown={e => {
          if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault()
            handleCreateNewFile()
          }
        }}
      >
        New File
      </div>
      <div
        role="button"
        tabindex="0"
        class="node-context-item"
        on:click={handleCreateNewDirectory}
        on:keydown={e => {
          if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault()
            handleCreateNewDirectory()
          }
        }}
      >
        New Directory
      </div>
    </div>
  {/if}

  {#if expanded && node.children}
    <div>
      {#each node.children as child}
        <TreeNode node={child} />
      {/each}
    </div>
  {/if}
</div>

<style>
  .node {
    margin-left: 1rem;
  }

  .node-content {
    font-size: 0.85rem;
    display: flex;
    column-gap: 0.25rem;
    min-height: 1rem;
    margin-bottom: 0.3rem;
    cursor: pointer;
  }

  .node-context-menu {
    position: fixed;
    z-index: 1000;
    background: #282828;
    border: 1px solid #323232;
    border-radius: 2px;
    font-size: 0.8rem;
    padding: 0.2rem 0;
  }

  .node-context-item {
    padding: 0.35rem 1.2rem;
    cursor: pointer;
  }

  .node-context-item:hover {
    color: #282828;
    background-color: #c9e6c1;
  }
</style>
