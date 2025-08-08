<script>
  import TreeNode from './TreeNode.svelte'
  import ContextMenu from './ContextMenu.svelte'

  import { GetNodeTree, MoveToTrash } from '../../../../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from '../../../stores/global.js'
  import { onMount, onDestroy } from 'svelte'
  import { Move } from 'lucide-svelte'

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
      parent: parentNode,
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
      parent: parentNode,
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

  const handleOpenRenameInput = () => {
    if (contextMenuTargetNode === null) {
      return
    }

    let targetFullPath = contextMenuTargetNode.path
    let parentNode = parentMap.get(targetFullPath)

    let dirPath = parentNode.path
    // Change view node to rename node
    contextMenuTargetNode.parent = parentNode
    contextMenuTargetNode.state = 'rename'
    contextMenuTargetNode.oldPath = contextMenuTargetNode.path

    forceTreeUpdate()
  }

  const handleMoveToTrash = async () => {
    if (contextMenuTargetNode === null) {
      return
    }

    let moveToTrashNode = contextMenuTargetNode
    try {
      await MoveToTrash(moveToTrashNode.path)
      removeNode(moveToTrashNode)
      forceTreeUpdate()
    } catch (err) {
      console.error('Error move to trash:', err)
    }
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
      sortNodeChildren(node)
      for (const child of node.children) {
        addNodeField(child)
      }
    }

    return node
  }

  const sortNodeChildren = (node) => {
    // Sort: directories first, then files, each alphabetically by name
    node.children.sort((a, b) => {
      if (a.type !== b.type) {
        return a.type === 'directory' ? -1 : 1
      }
      return a.name.localeCompare(b.name, undefined, { sensitivity: 'base' })
    })
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
      sortNodeChildren={sortNodeChildren}
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
    <ContextMenu
      contextMenuX={contextMenuX}
      contextMenuY={contextMenuY}
      handleOpenCreateFileInput={handleOpenCreateFileInput}
      handleOpenCreateDirectoryInput={handleOpenCreateDirectoryInput}
      handleOpenRenameInput={handleOpenRenameInput}
      handleMoveToTrash={handleMoveToTrash}
    />
  {/if}
</div>

<style>
  .directory-tree {
    width: 100%;
    overflow-x: hidden;
    overflow-y: scroll;
    box-sizing: border-box;
  }

  .directory-tree::-webkit-scrollbar {
    width: 8px;
  }

  .directory-tree::-webkit-scrollbar-track {
    background: transparent;
  }

  .directory-tree::-webkit-scrollbar-thumb {
    background-color: rgba(255, 255, 255, 0.2);
    border-radius: 4px;
  }
</style>
