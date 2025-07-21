<script>
  import TreeNode from './TreeNode.svelte';

  export let node;
  let expanded = node.is_root;

  const toggle = () => {
    if (node.type === 'directory') expanded = !expanded;
  };
</script>

<div class="node" style="margin-left: 1rem;">
  <div on:click={toggle} style="cursor: pointer;">
    {#if node.type === 'directory'}
      {expanded ? '📂' : '📁'} {node.name}
    {:else}
      📄 {node.name}
    {/if}
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
    line-height: 1.5;
  }
</style>
