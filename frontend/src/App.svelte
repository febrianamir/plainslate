<script>
  import SetRootPath from './components/SetRootPath.svelte'
  import Editor from './components/Editor.svelte'
  import Sidebar from './components/sidebar/Sidebar.svelte'
</script>

<div class="main-section">
  <p>plainslate</p>
  <div class="section-1">
    <SetRootPath />
  </div>
  <div class="section-2">
    <Sidebar />
    <Editor />
  </div>
</div>

<style>
  .section-2 {
    display: flex;
    flex-direction: row;
    overflow: hidden;
    width: 100%;
    height: 100vh;
  }
</style>
