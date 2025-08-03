<script>
  import { Folder, FolderOpen, File } from 'lucide-svelte'
  import TreeNode from './TreeNode.svelte'
  import { openedFile } from '../../../stores/global.js'
  import { onMount, onDestroy, tick } from 'svelte'

  export let node
  export let showContextMenu = false
  export let handleCloseContextMenu = () => {}
  export let onRightClick = () => {}
  export let removeNode = () => {}
  export let forceTreeUpdate = () => {}

  export let depth = 0

  let nodeActive = false
  $: nodeActive = $openedFile === node.path

  let inputRef
  $: if (node.state === 'create') {
    tick().then(() => {
      inputRef?.focus()
    })
  }

  let nodeEl

  const onClick = () => {
    if (showContextMenu) {
      return handleCloseContextMenu()
    }

    if (node.type === 'directory') {
      handleToggleExpand()
    }
    if (node.type === 'file' && node.state === 'view') {
      handleOpenFile()
    }
  }

  const handleToggleExpand = () => {
    node.expanded = !node.expanded
    forceTreeUpdate()
  }

  const handleOpenFile = () => {
    openedFile.set(node.path)
  }

  const handleCreateFile = () => {
    let dirPath = node.path.substring(0, node.path.lastIndexOf('/'))
    node.path = dirPath + '/' + node.name
    node.state = 'view'
    handleOpenFile()
    forceTreeUpdate()
  }

  const onClickOutside = (e) => {
    if (node.state !== 'create') {
      return
    }
    if (!nodeEl.contains(e.target)) {
      removeNode(node)
      forceTreeUpdate()
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
    on:click={onClick}
    on:contextmenu={(e) => onRightClick(node, e)}
    on:keydown={(e) => {
      // Prevent open file when editing input node
      if (document.activeElement.tagName === 'INPUT' || document.activeElement.isContentEditable) {
        return
      }

      if (e.key === 'Enter' || e.key === ' ') {
        e.preventDefault()
        handleOpenFile()
      }
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
          on:keydown={(e) => {
            if (e.key === 'Enter') {
              e.preventDefault()
              handleCreateFile()
            }
          }}
        />
      {/if}
    </div>
  </div>

  {#if node.expanded && node.children}
    <div>
      {#each node.children as child}
        <TreeNode
          removeNode={removeNode}
          forceTreeUpdate={forceTreeUpdate}
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
