<script>
  import TreeNode from './TreeNode.svelte'

  import { GetNodeTree } from '../../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from '../stores/global.js'
  import { onMount, onDestroy } from 'svelte'

  let unsubRootPath
  let tree = null

  onMount(() => {
    unsubRootPath = rootPath.subscribe(async dir => {
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
    unsubRootPath?.()
  })
</script>

<div class="directory-tree">
  {#if tree}
    <TreeNode node={tree} />
  {/if}
</div>
