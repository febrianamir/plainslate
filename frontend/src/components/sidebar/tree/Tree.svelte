<script>
  import TreeNode from './TreeNode.svelte'

  import { GetNodeTree } from '../../../../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from '../../../stores/global.js'
  import { onMount, onDestroy } from 'svelte'

  let unsubRootPath
  let tree = null

  // Context menu properties
  let showContextMenu = false
  let contextMenuX = 0
  let contextMenuY = 0
  let contextMenuTargetNode = null

  onMount(() => {
    window.addEventListener('click', handleCloseContextMenu)
    unsubRootPath = rootPath.subscribe(async (dir) => {
      if (dir && dir.trim() !== '') {
        try {
          const result = await GetNodeTree()
          tree = result
        } catch (err) {
          console.error('Error fetching node tree:', err)
          tree = null
        }
      }
    })
  })

  onDestroy(() => {
    window.removeEventListener('click', handleCloseContextMenu)
    unsubRootPath?.()
  })

  const onRightClick = (node, e) => {
    e.preventDefault()

    if (!showContextMenu) {
      return handleOpenContextMenu(node, e)
    }
    handleCloseContextMenu(node, e)
  }

  const handleOpenContextMenu = (node, e) => {
    contextMenuTargetNode = node
    contextMenuX = e.clientX
    contextMenuY = e.clientY
    showContextMenu = true
  }

  const handleCloseContextMenu = () => {
    contextMenuTargetNode = null
    contextMenuX = 0
    contextMenuY = 0
    showContextMenu = false
  }

  const handleCreateNewFile = () => {}

  const handleCreateNewDirectory = () => {}
</script>

<div class="directory-tree">
  {#if tree}
    <TreeNode
      node={tree}
      onRightClick={onRightClick}
      showContextMenu={showContextMenu}
      handleCloseContextMenu={handleCloseContextMenu}
    />
  {/if}

  {#if showContextMenu}
    <div class="tree-context-menu" style="top:{contextMenuY}px; left:{contextMenuX}px;">
      <div
        role="button"
        tabindex="0"
        class="tree-context-item"
        on:click={handleCreateNewFile}
        on:keydown={(e) => {
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
        class="tree-context-item"
        on:click={handleCreateNewDirectory}
        on:keydown={(e) => {
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
