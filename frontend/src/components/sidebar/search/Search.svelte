<script>
  import { SearchInFiles } from '../../../../wailsjs/go/usecase/Usecase.js'
  import SearchResultItem from './SearchResultItem.svelte'

  let query = ''
  let results = []
  let searchDebounceTimer

  function handleInput() {
    clearTimeout(searchDebounceTimer)
    searchDebounceTimer = setTimeout(async () => {
      if (query.trim()) {
        try {
          const result = await SearchInFiles(query)
          sortSearchResultFiles(result)
          results = result
        } catch (err) {
          console.log('Error search in files:', err)
        }
      } else {
        results = []
      }
    }, 200) // Debounce delay
  }

  function sortSearchResultFiles(results) {
    // Sort: alphabetically by name
    results.sort((a, b) => {
      let aFilename = a.file_path.split('/').pop()
      let bFilename = b.file_path.split('/').pop()
      return aFilename.localeCompare(bFilename, undefined, { sensitivity: 'base' })
    })
  }
</script>

<div class="search">
  <input
    class="search-input"
    bind:value={query}
    on:input={handleInput}
    placeholder="Search in folder..."
  />

  <div class="search-result">
    {#each results as file}
      <SearchResultItem
        filename={file.file_path.split('/').pop()}
        fileMatches={file.matches}
        path={file.file_path}
      />
    {/each}
  </div>
</div>

<style>
  .search {
    overflow-x: hidden;
    overflow-y: scroll;
    padding: 0.5rem 0.75rem 0.5rem 0.25rem;
  }

  .search-input {
    all: unset; /* Reset default styles */
    box-sizing: border-box;
    display: block;
    font-size: 0.85rem;
    padding: 0.5rem 0.75rem;
    width: 100%;
    background-color: #303336;
    text-overflow: ellipsis;
    overflow: hidden;
    border-radius: 15px;
  }

  .search-result {
    padding: 0.5rem 0;
    font-size: 0.85rem;
  }
</style>
