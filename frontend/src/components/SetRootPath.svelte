<script>
  import { SetRootPath } from '../../wailsjs/go/usecase/Usecase.js';
  import { rootPath } from '../stores/global.js';

  let path = '';
  let errorMessage = '';
  let successMessage = '';

  async function savePath() {
    errorMessage = '';
    successMessage = '';
    try {
      await SetRootPath(path);
      successMessage = 'Directory saved!';
      rootPath.set(path)
    } catch (err) {
      errorMessage = err.message || 'Failed to save directory.';
    }
  }
</script>

<div style="display: flex; gap: 0.5rem; align-items: center;">
  <input
    type="text"
    bind:value={path}
    placeholder="Enter directory path"
    style="flex: 1; padding: 0.5rem;"
  />
  <button on:click={savePath}>Save</button>
</div>

{#if errorMessage}
  <p style="color: red; margin-top: 0.5rem;">{errorMessage}</p>
{/if}

{#if successMessage}
  <p style="color: green; margin-top: 0.5rem;">{successMessage}</p>
{/if}