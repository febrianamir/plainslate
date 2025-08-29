<script>
  import { SetRootPath } from '../../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from '../stores/global.js'
  import { onDestroy, onMount } from 'svelte'
  import { handleEnter } from '../../src/lib/utils.js'

  let path = $state('')
  let errorMessage = $state('')
  let successMessage = $state('')

  let unsubRootPath

  async function setRootPath() {
    errorMessage = ''
    successMessage = ''
    try {
      const req = {
        rootPath: path,
      }
      await SetRootPath(path)
      successMessage = 'Directory saved!'
      rootPath.set(path)
    } catch (err) {
      errorMessage = err.message || 'Failed to save directory.'
    }
  }

  async function onKeyDown(e) {
    await handleEnter(e, setRootPath)
  }

  onMount(() => {
    unsubRootPath = rootPath.subscribe(async (dir) => {
      // Set input value from rootPath store
      if (dir) {
        path = dir
      }
    })
  })

  onDestroy(() => {
    unsubRootPath?.()
  })
</script>

<div class="config-root-path">
  <input
    type="text"
    bind:value={path}
    placeholder="Enter directory path"
    style="flex: 1; padding: 0.5rem;"
    onkeydown={onKeyDown}
  />
</div>

{#if errorMessage}
  <p style="color: red; margin-top: 0.5rem;">{errorMessage}</p>
{/if}

{#if successMessage}
  <p style="color: green; margin-top: 0.5rem;">{successMessage}</p>
{/if}

<style>
  .config-root-path {
    display: flex;
    gap: 0.5rem;
    align-items: center;
    padding: 0 1rem 1rem 1rem;
  }
</style>
