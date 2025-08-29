<script>
  import EditorWokspace from './components/editor/EditorWokspace.svelte'
  import Sidebar from './components/sidebar/Sidebar.svelte'
  import { GetConfig } from '../wailsjs/go/usecase/Usecase.js'
  import { rootPath } from './stores/global.js'
  import { onMount } from 'svelte'

  onMount(async () => {
    const config = await GetConfig()
    if (config?.rootPath !== '') {
      rootPath.set(config.rootPath)
    }
  })
</script>

<div class="app">
  <Sidebar />
  <EditorWokspace />
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
