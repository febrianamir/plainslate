<script>
  import { ChevronDown, ChevronRight } from 'lucide-svelte'
  import { handleEnter } from '../../../../src/lib/utils'

  export let filename
  export let path
  export let fileMatches

  let isExpanded = true

  function toggleExpand() {
    isExpanded = !isExpanded
  }
</script>

<div class="search-result-item">
  <div
    class="search-result-file"
    on:click={toggleExpand}
    on:keydown={(e) => {
      handleEnter(e, toggleExpand)
    }}
  >
    <div class="search-result-file-icon">
      {#if isExpanded}
        <ChevronDown size={16} strokeWidth={3} />
      {:else}
        <ChevronRight size={16} strokeWidth={3} />
      {/if}
    </div>
    <div>
      {filename}
    </div>
  </div>
  <div class="search-result-matches" class:active={isExpanded}>
    {#each fileMatches as match}
      <div class="search-result-match">{match}</div>
    {/each}
  </div>
</div>

<style>
  .search-result-file,
  .search-result-match {
    user-select: none;
    white-space: nowrap;
    overflow: hidden;
    padding: 0.1rem 0;
  }

  .search-result-file {
    display: flex;
    gap: 0.1rem;
    cursor: default;
    font-weight: 600;
  }

  .search-result-matches {
    display: none;
  }

  .search-result-matches.active {
    display: block;
  }

  .search-result-match {
    cursor: pointer;
    padding-left: 0.75rem;
  }
</style>
