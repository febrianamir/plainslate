<script>
  import Editor from './components/Editor.svelte'
  import Sidebar from './components/sidebar/Sidebar.svelte'
  import { GetConfig } from '../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from './stores/global.js'
  import { onMount } from 'svelte'

  onMount(async () => {
    const config = await GetConfig()
    if (config?.root_path !== '') {
      rootPath.set(config.root_path)
    }
  })
</script>

<div class="app">
  <Sidebar />
  <Editor />
</div>

<style>
  .app {
    display: flex;
    flex-direction: row;
    overflow: hidden;
    width: 100%;
    height: 100vh;
  }
</style>
