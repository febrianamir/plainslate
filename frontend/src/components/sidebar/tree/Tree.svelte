<script>
  import TreeNode from './TreeNode.svelte'

  import { GetNodeTree } from '../../../../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from '../../../stores/global.js'
  import { onMount, onDestroy } from 'svelte'

  let unsubRootPath
  let depth = 0
  let tree = null
  let parentMap = new Map()
  // forceTreeUpdate to force reactivity after changes to tree
  const forceTreeUpdate = () => {
    tree = structuredClone(tree)
    parentMap = indexParents(tree)
  }

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
          tree = addNodeField(result)
          parentMap = indexParents(tree)
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

  const handleOpenCreateFileInput = () => {
    if (contextMenuTargetNode === null) {
      return
    }

    let targetFullPath = contextMenuTargetNode.path
    let parentNode = contextMenuTargetNode
    if (contextMenuTargetNode.type === 'file') {
      parentNode = parentMap.get(targetFullPath)
    }

    let dirPath = parentNode.path
    let newNode = {
      state: 'create',
      name: '',
      path: dirPath + '/' + 'new_file.md',
      type: 'file',
      children: [],
    }

    insertNode(parentNode, newNode)
    contextMenuTargetNode.expanded = true
    forceTreeUpdate()
  }

  const handleOpenCreateDirectoryInput = () => {
    if (contextMenuTargetNode === null) {
      return
    }

    let targetFullPath = contextMenuTargetNode.path
    let parentNode = contextMenuTargetNode
    if (contextMenuTargetNode.type === 'file') {
      parentNode = parentMap.get(targetFullPath)
    }

    let dirPath = parentNode.path
    let newNode = {
      state: 'create',
      name: '',
      path: dirPath + '/' + 'new_directory',
      type: 'directory',
      children: [],
    }

    insertNode(parentNode, newNode)
    contextMenuTargetNode.expanded = true
    forceTreeUpdate()
  }

  const insertNode = (targetNode, newNode) => {
    if (!targetNode.children) {
      targetNode.children = []
    }
    targetNode.children = [newNode, ...(targetNode.children || [])]
  }

  const removeNode = (targetNode) => {
    const parent = parentMap.get(targetNode.path)
    if (!parent || !parent.children) {
      return
    }
    parent.children = parent.children.filter((n) => {
      return n !== targetNode
    })
  }

  const addNodeField = (node) => {
    node.expanded = node.is_root || false
    node.state = 'view'
    if (node.children && Array.isArray(node.children)) {
      for (const child of node.children) {
        addNodeField(child)
      }
    }
    return node
  }

  const indexParents = (node, parent = null, map = new Map()) => {
    map.set(node.path, parent)
    if (node.children) {
      for (const child of node.children) {
        indexParents(child, node, map)
      }
    }
    return map
  }
</script>

<div class="directory-tree">
  {#if tree}
    <TreeNode
      removeNode={removeNode}
      forceTreeUpdate={forceTreeUpdate}
      node={tree}
      onRightClick={onRightClick}
      showContextMenu={showContextMenu}
      handleCloseContextMenu={handleCloseContextMenu}
      depth={depth}
    />
  {/if}

  {#if showContextMenu}
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
  {/if}
</div>

<style>
  .directory-tree {
    width: 100%;
    overflow: hidden;
    box-sizing: border-box;
  }

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
