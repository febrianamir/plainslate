<script>
  import { Folder, FolderOpen, File } from 'lucide-svelte'
  import TreeNode from './TreeNode.svelte'
  import { openedFile } from '../../../stores/global.js'
  import { onMount, onDestroy } from 'svelte'

  export let node
  export let showContextMenu = false
  export let handleCloseContextMenu = () => {}
  export let onRightClick = () => {}

  export let depth = 0
  let expanded = node.is_root
  let nodeActive = false
  $: nodeActive = $openedFile === node.path

  const onClick = () => {
    if (showContextMenu) {
      return handleCloseContextMenu()
    }

    if (node.type === 'directory') {
      handleToggleExpand()
    }
    if (node.type === 'file') {
      handleOpenFile()
    }
  }

  const handleToggleExpand = () => {
    expanded = !expanded
  }

  const handleOpenFile = () => {
    openedFile.set(node.path)
  }
</script>

<div class="node">
  <div
    role="button"
    tabindex="0"
    class="node-content"
    class:active={nodeActive}
    style="padding-left: {depth + 1}rem"
    on:click={onClick}
    on:contextmenu={(e) => onRightClick(node, e)}
    on:keydown={(e) => {
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

  {#if expanded && node.children}
    <div>
      {#each node.children as child}
        <TreeNode
          node={child}
          onRightClick={onRightClick}
          showContextMenu={showContextMenu}
          handleCloseContextMenu={handleCloseContextMenu}
          depth={depth + 1}
        />
      {/each}
    </div>
  {/if}
</div>

<style>
  .node-content {
    position: relative;
    padding: 0.35rem 0 0.25rem 0;
    font-size: 0.85rem;
    display: flex;
    column-gap: 0.25rem;
    min-height: 1rem;
    cursor: default;
  }

  .node-content:hover {
    background-color: #303030;
  }

  .node-content.active {
    color: #6da96f;
  }

  .node-content.active::before {
    content: '';
    position: absolute;
    width: 2px;
    top: 0;
    left: 0;
    bottom: 0;
    background-color: #6da96f;
  }
</style>
