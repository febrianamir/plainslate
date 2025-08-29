<script>
  import { debounce } from '../../../lib/utils.js'
  import { SearchInFiles } from '../../../../wailsjs/go/usecase/Usecase.js'
  import { CaseSensitive } from 'lucide-svelte'
  import SearchResultItem from './SearchResultItem.svelte'

  let query = $state('')
  let isCaseSensitive = $state(false)
  let results = $state([])

  const debouncedSearchInFiles = debounce(async (query, isCaseSensitive) => {
    if (query.trim()) {
      try {
        const req = {
          query: query,
          isCaseSensitive: isCaseSensitive,
        }
        const resp = await SearchInFiles(req)
        sortSearchResultFiles(resp)
        results = resp
      } catch (err) {
        console.log('Error search in files:', err)
      }
    } else {
      results = []
    }
  }, 200)

  function toggleCaseSensitive() {
    isCaseSensitive = !isCaseSensitive
  }

  $effect(() => {
    debouncedSearchInFiles(query, isCaseSensitive)
  })

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
    <input class="search-input" bind:value={query} placeholder="Search in folder..." />
    <div
      role="button"
      tabindex="0"
      class:active={isCaseSensitive}
      class="search-case-sensitive"
      onclick={toggleCaseSensitive}
      onkeydown={(e) => {
        handleEnter(e, toggleCaseSensitive)
      }}
    >
      <CaseSensitive size={24} strokeWidth={1.75} />
    </div>
  </div>

  <div class="search-result">
    {#each results as file}
      <SearchResultItem
        filename={file.file_path.split('/').pop()}
        fileMatches={file.matches}
        query={query}
        path={file.file_path}
        isCaseSensitive={isCaseSensitive}
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
    position: relative;
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

  .search-case-sensitive {
    position: absolute;
    top: 0;
    right: 0.75rem;
    padding: 0.35rem 0.5rem;
    cursor: pointer;
  }

  .search-case-sensitive:hover,
  .search-case-sensitive.active {
    color: #c9e6c1;
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
