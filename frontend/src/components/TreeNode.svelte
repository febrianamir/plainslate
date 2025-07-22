<script>
  import { Folder, FolderOpen, File } from 'lucide-svelte'
  import TreeNode from './TreeNode.svelte'
  import { openedFile } from '../stores/global.js'

  export let node
  let expanded = node.is_root

  const onClickNode = () => {
    if (node.type === 'directory') {
      toggleExpand()
    }
    if (node.type === 'file') {
      openFile()
    }
  }

  const toggleExpand = () => {
    expanded = !expanded
  }

  const openFile = () => {
    openedFile.set(node.path)
  }
</script>

<div class="node">
  <div on:click={onClickNode} class="node-content">
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
        <TreeNode node={child} />
      {/each}
    </div>
  {/if}
</div>

<style>
  .node {
    font-family: monospace;
    margin-left: 1rem;
  }

  .node-content {
    display: flex;
    column-gap: 0.25rem;
    min-height: 1rem;
    margin-bottom: 0.3rem;
    cursor: pointer;
  }
</style>
