<script>
  import Editor from './components/Editor.svelte'
  import Sidebar from './components/sidebar/Sidebar.svelte'
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
