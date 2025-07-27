<script>
  import { SetRootPath } from '../../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from '../stores/global.js'

  let path = ''
  let errorMessage = ''
  let successMessage = ''

  const handleSavePath = async () => {
    errorMessage = ''
    successMessage = ''
    try {
      await SetRootPath(path)
      successMessage = 'Directory saved!'
      rootPath.set(path)
    } catch (err) {
      errorMessage = err.message || 'Failed to save directory.'
    }
  }

  const onKeyDown = async e => {
    if (e.key === 'Enter') {
      await handleSavePath()
    }
  }
</script>

<div class="config-root-path">
  <input
    type="text"
    bind:value={path}
    placeholder="Enter directory path"
    style="flex: 1; padding: 0.5rem;"
    on:keydown={onKeyDown}
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
