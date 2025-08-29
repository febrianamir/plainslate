<script>
  import { Folder, FolderOpen, File } from 'lucide-svelte'
  import TreeNode from './TreeNode.svelte'
  import { openedFilesUpdateFile, getOpenedFiles } from '../../../state/openedFile.svelte.js'
  import { onMount, onDestroy, tick } from 'svelte'
  import { CreateDirectory, RenamePath, SaveFile } from '../../../../wailsjs/go/usecase/Usecase.js'
  import { openFile } from '../../../lib/services/fileService.js'
  import { handleEnter } from '../../../../src/lib/utils.js'

  let openedFiles = getOpenedFiles()

  let {
    parentMap,
    node = $bindable(),
    isShowContextMenu = false,
    closeContextMenu = () => {},
    onRightClick = () => {},
    removeNode = () => {},
    sortNodeChildren = () => {},
    indexTreeParents = () => {},
    depth = 0,
  } = $props()

  let nodeActive = $state(false)
  $effect(() => {
    if (openedFiles.activeId) {
      nodeActive = openedFiles.activeId === node.path
    } else {
      nodeActive = false
    }
  })

  let inputRef = $state()
  $effect(() => {
    if (node.state !== 'view') {
      tick().then(() => {
        inputRef?.focus()
      })
    }
  })

  let nodeEl = $state()

  function onClick() {
    if (isShowContextMenu) {
      return closeContextMenu()
    }

    if (node.type === 'directory' && node.state === 'view') {
      toggleExpandFolder()
    }
    if (node.type === 'file' && node.state === 'view') {
      openFile(node.path)
    }
  }

  function toggleExpandFolder() {
    node.expanded = !node.expanded
  }

  function createFile() {
    let dirPath = node.path.substring(0, node.path.lastIndexOf('/'))
    node.path = dirPath + '/' + node.name
    node.state = 'view'

    openFile(node.path)
    if (node.parent) {
      let parentNode = node.parent
      delete node.parent
      sortNodeChildren(parentNode)
    }
    indexTreeParents()
  }

  async function createFolder() {
    try {
      let dirPath = node.path.substring(0, node.path.lastIndexOf('/')) + '/' + node.name
      const req = {
        path: dirPath,
      }
      await CreateDirectory(req)
      node.path = dirPath
      node.state = 'view'

      if (node.parent) {
        let parentNode = node.parent
        delete node.parent
        sortNodeChildren(parentNode)
      }
      indexTreeParents()
    } catch (err) {
      console.error('Error creating directory:', err)
    }
  }

  async function rename() {
    if (!node.oldPath || node.oldPath === '') {
      return
    }

    try {
      let newPath = node.path.substring(0, node.path.lastIndexOf('/')) + '/' + node.name

      if (node.oldPath !== newPath) {
        const req = {
          oldPath: node.oldPath,
          newPath: newPath,
        }
        await RenamePath(req)
      }
      node.path = newPath
      node.state = 'view'

      // Sort node after rename
      let parentNode = parentMap.get(node.oldPath)
      if (parentNode) {
        sortNodeChildren(parentNode)
      }

      // Update opened file data
      if (node.oldPath) {
        openedFilesUpdateFile(node.oldPath, {
          id: node.path,
          filepath: node.path,
          filename: node.name,
        })
        delete node.oldPath
      }
      indexTreeParents()
    } catch (err) {
      console.error('Error renaming path:', err)
    }
  }

  function onClickOutside(e) {
    if (node.state === 'view') {
      return
    }

    if (!nodeEl.contains(e.target)) {
      switch (node.state) {
        case 'create':
          removeNode(node)
          break
        case 'rename':
          node.state = 'view'
          node.path = node.oldPath
          delete node.oldPath
          delete node.parent
      }
      indexTreeParents()
    }
  }

  onMount(() => {
    window.addEventListener('click', onClickOutside, true)
  })

  onDestroy(() => {
    window.removeEventListener('click', onClickOutside, true)
  })
</script>

<div class="node">
  <div
    role="button"
    tabindex="0"
    bind:this={nodeEl}
    class="node-content"
    class:active={nodeActive}
    class:edit={node.state !== 'view'}
    style="--depth: {depth + 1}"
    onclick={onClick}
    oncontextmenu={(e) => onRightClick(node, e)}
    onkeydown={(e) => {
      const target = e.target

      // Prevent file open when editing an input or any editable content
      let isEditingContent =
        target instanceof HTMLInputElement ||
        target instanceof HTMLTextAreaElement ||
        target.isContentEditable
      if (isEditingContent) {
        return
      }

      handleEnter(e, () => openFile(node.path))
    }}
  >
    <div class="node-icon">
      {#if node.type === 'directory'}
        {#if node.expanded}
          <FolderOpen size={16} />
        {:else}
          <Folder size={16} />
        {/if}
      {:else}
        <File size={16} />
      {/if}
    </div>
    <div class="node-text">
      {#if node.state === 'view'}
        {node.name}
      {:else}
        <input
          class="node-input"
          type="text"
          bind:this={inputRef}
          bind:value={node.name}
          onkeydown={(e) => {
            handleEnter(e, () => {
              if (node.state === 'create' && node.type === 'file') {
                return createFile()
              }

              if (node.state === 'create' && node.type === 'directory') {
                return createFolder()
              }

              if (node.state === 'rename') {
                return rename()
              }
            })
          }}
        />
      {/if}
    </div>
  </div>

  {#if node.expanded && node.children}
    <div>
      {#each node.children as child, i}
        <TreeNode
          parentMap={parentMap}
          sortNodeChildren={sortNodeChildren}
          removeNode={removeNode}
          indexTreeParents={indexTreeParents}
          bind:node={node.children[i]}
          onRightClick={onRightClick}
          isShowContextMenu={isShowContextMenu}
          closeContextMenu={closeContextMenu}
          depth={depth + 1}
        />
      {/each}
    </div>
  {/if}
</div>

<style>
  .node {
    min-width: 0;
    max-width: 100%;
    overflow: hidden;
  }

  .node-content {
    min-width: 0;
    max-width: 100%;
    position: relative;
    padding: 0.35rem 0 0.25rem 0;
    padding-left: calc(var(--depth, 0) * 1rem);
    font-size: 0.85rem;
    display: flex;
    column-gap: 0.25rem;
    min-height: 1rem;
    cursor: default;
  }

  .node-content:hover {
    background-color: #303030;
  }

  .node-content.edit {
    background-color: #141414;
  }

  .node-content.edit:hover {
    background-color: #141414;
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

  .node-text {
    flex: 1;
    min-width: 0;
    text-overflow: ellipsis;
    overflow: hidden;
  }

  .node-input {
    all: unset; /* Resets most browser default styles */
    width: 100%;
    font-size: 0.85rem;
    background-color: #141414;
    color: #9d9d9d;
    cursor: text;
    text-overflow: ellipsis;
    overflow: hidden;
  }
</style>
