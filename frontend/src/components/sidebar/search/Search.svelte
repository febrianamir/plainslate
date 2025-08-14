<script>
  import { SearchInFiles } from '../../../../wailsjs/go/usecase/Usecase.js'
  import SearchResultItem from './SearchResultItem.svelte'

  let query = $state('')
  let results = $state([])
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
  <div class="search-bar">
    <input
      class="search-input"
      bind:value={query}
      oninput={handleInput}
      placeholder="Search in folder..."
    />
  </div>

  <div class="search-result">
    {#each results as file}
      <SearchResultItem
        filename={file.file_path.split('/').pop()}
        fileMatches={file.matches}
        query={query}
        path={file.file_path}
      />
    {/each}
  </div>
</div>

<style>
  .search {
    display: flex;
    flex-direction: column;
    overflow-x: hidden;
    padding: 0.5rem 0 0 0.25rem;
  }

  .search-bar {
    padding-right: 0.75rem;
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
    flex: 1;
    padding: 0.5rem 0.75rem 0 0;
    font-size: 0.85rem;
    overflow-x: hidden;
    overflow-y: scroll;
  }

  .search-result::-webkit-scrollbar {
    width: 8px;
  }

  .search-result::-webkit-scrollbar-track {
    background: transparent;
  }

  .search-result::-webkit-scrollbar-thumb {
    background-color: rgba(255, 255, 255, 0.2);
    border-radius: 4px;
  }
</style>
