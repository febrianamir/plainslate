<script>
  import TreeNode from './TreeNode.svelte'
  import ContextMenu from './ContextMenu.svelte'
  import { openedFilesClose } from '../../../state/openedFile.svelte.js'
  import { GetNodeTree, MoveToTrash } from '../../../../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from '../../../stores/global.js'
  import { onMount, onDestroy } from 'svelte'

  let unsubRootPath
  let depth = 0
  let tree = $state(null)

  // svelte-ignore non_reactive_update
  let parentMap = new Map()
  function indexTreeParents() {
    parentMap = indexParents(tree)
  }

  // Context menu properties
  let isShowContextMenu = $state(false)
  let contextMenuX = $state(0)
  let contextMenuY = $state(0)
  let contextMenuTargetNode = null

  onMount(() => {
    window.addEventListener('click', closeContextMenu)
    unsubRootPath = rootPath.subscribe(async (dir) => {
      if (dir && dir.trim() !== '') {
        try {
          const result = await GetNodeTree()
          tree = addNodeField(result)
          indexTreeParents()
        } catch (err) {
          console.error('Error fetching node tree:', err)
          tree = null
        }
      }
    })
  })

  onDestroy(() => {
    window.removeEventListener('click', closeContextMenu)
    unsubRootPath?.()
  })

  function onRightClick(node, e) {
    e.preventDefault()

    if (!isShowContextMenu) {
      return openContextMenu(node, e)
    }
    closeContextMenu(node, e)
  }

  function openContextMenu(node, e) {
    contextMenuTargetNode = node
    contextMenuX = e.clientX
    contextMenuY = e.clientY
    isShowContextMenu = true
  }

  function closeContextMenu() {
    contextMenuTargetNode = null
    contextMenuX = 0
    contextMenuY = 0
    isShowContextMenu = false
  }

  function showCreateFileInput() {
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
    indexTreeParents()
  }

  function showCreateFolderInput() {
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
    indexTreeParents()
  }

  function showRenameInput() {
    if (contextMenuTargetNode === null) {
      return
    }

    let targetFullPath = contextMenuTargetNode.path
    let parentNode = parentMap.get(targetFullPath)
    let dirPath = parentNode.path

    // Change view node to rename node
    contextMenuTargetNode.state = 'rename'
    contextMenuTargetNode.oldPath = contextMenuTargetNode.path
    indexTreeParents()
  }

  async function moveItemToTrash() {
    if (contextMenuTargetNode === null) {
      return
    }

    let moveToTrashNode = contextMenuTargetNode
    try {
      const req = {
        path: moveToTrashNode.path,
      }
      await MoveToTrash(req)
      openedFilesClose(moveToTrashNode.path)
      removeNode(moveToTrashNode)
      indexTreeParents()
    } catch (err) {
      console.error('Error move to trash:', err)
    }
  }

  function insertNode(targetNode, newNode) {
    if (!targetNode.children) {
      targetNode.children = []
    }
    targetNode.children = [newNode, ...(targetNode.children || [])]
  }

  function removeNode(targetNode) {
    const parent = parentMap.get(targetNode.path)
    if (!parent || !parent.children) {
      return
    }
    parent.children = parent.children.filter((n) => {
      return n !== targetNode
    })
  }

  function addNodeField(node) {
    node.expanded = node.isRoot || false
    node.state = 'view'

    if (node.children && Array.isArray(node.children)) {
      sortNodeChildren(node)
      for (const child of node.children) {
        addNodeField(child)
      }
    }

    return node
  }

  function sortNodeChildren(node) {
    // Sort: directories first, then files, each alphabetically by name
    node.children.sort((a, b) => {
      if (a.type !== b.type) {
        return a.type === 'directory' ? -1 : 1
      }
      return a.name.localeCompare(b.name, undefined, { sensitivity: 'base' })
    })
  }

  function indexParents(node, parent = null, map = new Map()) {
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
      parentMap={parentMap}
      sortNodeChildren={sortNodeChildren}
      removeNode={removeNode}
      indexTreeParents={indexTreeParents}
      bind:node={tree}
      onRightClick={onRightClick}
      isShowContextMenu={isShowContextMenu}
      closeContextMenu={closeContextMenu}
      depth={depth}
    />
  {/if}

  {#if isShowContextMenu}
    <ContextMenu
      contextMenuX={contextMenuX}
      contextMenuY={contextMenuY}
      showCreateFileInput={showCreateFileInput}
      showCreateFolderInput={showCreateFolderInput}
      showRenameInput={showRenameInput}
      moveItemToTrash={moveItemToTrash}
    />
  {/if}
</div>

<style>
  .directory-tree {
    width: 100%;
    min-height: 0;
    overflow-x: hidden;
    overflow-y: scroll;
    padding-bottom: 14rem;
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
