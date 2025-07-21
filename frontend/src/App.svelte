<script>
  import SetRootPath from "./components/SetRootPath.svelte";
  import DirectoryTree from "./components/DirectoryTree.svelte";
  import Editor from "./components/Editor.svelte";
</script>

<div class="main-section">
  <p>plainslate</p>
  <div class="section-1">
    <SetRootPath />
  </div>
  <div class="section-2">
    <div class="directory-tree">
      <DirectoryTree />
    </div>
    <div class="editor">
      <Editor />
    </div>
  </div>
</div>

<style>
  .section-2 {
    display: flex;
  }

  .section-2 .directory-tree {
    width: 20%;
  }

  .section-2 .editor {
    width: 80%;
  }
</style>
