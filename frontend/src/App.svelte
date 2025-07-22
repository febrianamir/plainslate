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
    <div class="editor">
      <Editor />
    </div>
  </div>
</div>

<style>
  .section-2 {
    display: flex;
    height: 100vh;
  }

  .section-2 .directory-tree {
    width: 20%;
  }

  .section-2 .editor {
    width: 80%;
  }
</style>
