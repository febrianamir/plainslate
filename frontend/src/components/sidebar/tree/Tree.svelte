<script>
  import TreeNode from './TreeNode.svelte'
  import ContextMenu from './ContextMenu.svelte'
  import {
    openedFilesUpdateFile,
    openedFilesCheckExist,
    openedFilesClose,
  } from '../../../state/openedFile.svelte.js'
  import { clipboard, cleanClipboard } from '../../../state/clipboard.svelte.js'
  import {
    GetRootNodeTree,
    GetNodeTree,
    MoveToTrash,
    RenamePath,
    CopyFile,
    CopyDirectory,
    CheckPath,
  } from '../../../../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from '../../../stores/global.js'
  import { onMount, onDestroy } from 'svelte'
  import { parseFilepath, getLastDirName, getParentDirPath } from '../../../lib/utils.js'

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
          const result = await GetRootNodeTree()
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

      if (moveToTrashNode.type === 'file') {
        openedFilesClose(moveToTrashNode.path)
      }

      if (moveToTrashNode.type === 'directory') {
        let paths = getNodePaths(moveToTrashNode)
        paths.forEach(function (v, i) {
          openedFilesClose(v)
        })
      }

      removeNode(moveToTrashNode)
      indexTreeParents()
    } catch (err) {
      console.error('Error move to trash:', err)
    }
  }

  function addToClipboard(clipboardType) {
    if (contextMenuTargetNode === null) {
      return
    }

    let filename = contextMenuTargetNode.path.split('/').pop()

    clipboard.clipboardType = clipboardType
    clipboard.path = contextMenuTargetNode.path
    clipboard.filename = filename
    clipboard.node = contextMenuTargetNode
  }

  async function paste() {
    if (contextMenuTargetNode === null) {
      return
    }

    if (clipboard.clipboardType === '') {
      return
    }

    let targetFullPath = contextMenuTargetNode.path
    let destParentNode = contextMenuTargetNode
    if (contextMenuTargetNode.type === 'file') {
      destParentNode = parentMap.get(targetFullPath)
    }
    let sourcePath = clipboard.path
    let destPath = destParentNode.path + '/' + clipboard.filename

    if (clipboard.clipboardType === 'CUT' && sourcePath !== destPath) {
      const req = {
        oldPath: sourcePath,
        newPath: destPath,
      }
      await RenamePath(req)

      // Remove node from source path
      removeNode(clipboard.node)

      if (clipboard.node.type === 'file') {
        cutFile(sourcePath, destPath, clipboard.node, destParentNode)
      }

      if (clipboard.node.type === 'directory') {
        cutDirectory(sourcePath, destPath, clipboard.node, destParentNode)
      }
    }

    if (clipboard.clipboardType === 'COPY') {
      if (clipboard.node.type === 'file') {
        await copyFile(sourcePath, destPath, destParentNode)
      }

      if (clipboard.node.type === 'directory') {
        await copyDirectory(sourcePath, destPath, destParentNode)
      }
    }

    // Clean clipboard
    cleanClipboard()
  }

  async function cutFile(sourcePath, destPath, node, destParentNode) {
    // Insert new node to destination path
    node.path = destPath
    insertNode(destParentNode, node)

    // Update opened file path
    if (openedFilesCheckExist(sourcePath)) {
      openedFilesUpdateFile(sourcePath, {
        id: destPath,
        filePath: destPath,
        fileName: node.name,
      })
    }

    // Reorder the destination parent
    sortNodeChildren(destParentNode)

    // Reindex tree parents
    indexTreeParents()
  }

  async function cutDirectory(sourcePath, destPath, node, destParentNode) {
    try {
      // Insert new node to destination path
      const getNodeTreeReq = {
        path: destPath,
      }
      const result = await GetNodeTree(getNodeTreeReq)
      let newDirectoryNode = addNodeField(result)
      insertNode(destParentNode, newDirectoryNode)

      // Update opened file path
      let paths = getNodePaths(node)
      let newPaths = getNodePaths(newDirectoryNode)
      paths.forEach(function (v, i) {
        let op = v
        let np = newPaths[i]
        let file = parseFilepath(np)

        if (openedFilesCheckExist(op)) {
          openedFilesUpdateFile(op, {
            id: np,
            filePath: np,
            fileName: file.name + '.' + file.extension,
          })
        }
      })

      // Reorder the destination parent
      sortNodeChildren(destParentNode)

      // Reindex tree parents
      indexTreeParents()
    } catch (err) {
      console.error('Error update tree node after paste:', err)
    }
  }

  async function copyFile(sourcePath, destPath, destParentNode) {
    try {
      let file = parseFilepath(destPath)
      let safeDestPath = await generateSafeDestFilePath(
        destPath,
        file.name + '.' + file.extension,
        0
      )

      const req = {
        sourcePath: sourcePath,
        destPath: safeDestPath,
      }
      await CopyFile(req)

      // Add new node to destination path
      let newFile = parseFilepath(safeDestPath)
      let newNode = {
        state: 'view',
        name: newFile.name + '.' + newFile.extension,
        path: safeDestPath,
        type: 'file',
        children: [],
      }
      insertNode(destParentNode, newNode)

      // Reorder the destination parent
      sortNodeChildren(destParentNode)

      // Reindex tree parents
      indexTreeParents()
    } catch (err) {
      console.error('Error copy to destination:', err)
    }
  }

  async function copyDirectory(sourcePath, destPath, destParentNode) {
    try {
      let originalDirName = getLastDirName(destPath)
      let safeDestPath = await generateSafeDestDirPath(destPath, originalDirName, 0)

      const req = {
        sourcePath: sourcePath,
        destPath: safeDestPath,
      }
      await CopyDirectory(req)

      // Get new directory tree
      const getNodeTreeReq = {
        path: safeDestPath,
      }
      const result = await GetNodeTree(getNodeTreeReq)
      let newDirectoryNode = addNodeField(result)

      // Add new node to destination path
      insertNode(destParentNode, newDirectoryNode)

      // Reorder the destination parent
      sortNodeChildren(destParentNode)

      // Reindex tree parents
      indexTreeParents()
    } catch (err) {
      console.error('Error copy to destination:', err)
    }
  }

  async function generateSafeDestFilePath(destPath, originalFilename, iteration) {
    // Check duplicate file
    const checkPathReq = {
      path: destPath,
    }
    const isDuplicate = await CheckPath(checkPathReq)

    if (isDuplicate) {
      // Create new destination path
      iteration++
      let file = parseFilepath(destPath)
      let newDestPath =
        file.dirpath + originalFilename + ' (' + iteration + ')' + '.' + file.extension
      newDestPath = await generateSafeDestFilePath(newDestPath, originalFilename, iteration)
      return newDestPath
    }
    return destPath
  }

  async function generateSafeDestDirPath(destPath, originalDir, iteration) {
    // Check duplicate dir
    const checkPathReq = {
      path: destPath,
    }
    const isDuplicate = await CheckPath(checkPathReq)

    if (isDuplicate) {
      // Create new destination path
      iteration++
      let dirpath = getParentDirPath(destPath)
      let newDestPath = dirpath + originalDir + ' (' + iteration + ')'
      newDestPath = await generateSafeDestDirPath(newDestPath, originalDir, iteration)
      return newDestPath
    }
    return destPath
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

  function getNodePaths(node, paths = []) {
    paths.push(node.path)
    if (node.children) {
      for (const child of node.children) {
        getNodePaths(child, paths)
      }
    }
    return paths
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
      addToClipboard={addToClipboard}
      paste={paste}
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
