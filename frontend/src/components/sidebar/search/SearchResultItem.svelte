<script>
  import { ChevronDown, ChevronRight } from 'lucide-svelte'
  import { handleEnter } from '../../../../src/lib/utils'
  import { onDestroy, onMount } from 'svelte'

  export let filename
  export let path
  export let fileMatches

  let clientX = 0
  let clientY = 0

  let isExpanded = true
  let isShowTooltip = false
  let tooltipX = 0
  let tooltipY = 0
  let tooltipShowTimer

  function toggleExpand() {
    isExpanded = !isExpanded
  }

  function openTooltip(e) {
    tooltipX = clientX + 10
    tooltipY = clientY + 10
    isShowTooltip = true
  }

  function closeTooltip(e) {
    isShowTooltip = false
  }

  function onMouseMove(e) {
    clientX = e.clientX
    clientY = e.clientY
  }

  onMount(() => {
    window.addEventListener('mousemove', onMouseMove)
  })

  onDestroy(() => {
    window.removeEventListener('mousemove', onMouseMove)
  })
</script>

<div class="search-result-item">
  <div
    class="search-result-file"
    on:click={() => {
      clearTimeout(tooltipShowTimer)
      toggleExpand()
    }}
    on:keydown={(e) => {
      clearTimeout(tooltipShowTimer)
      handleEnter(e, toggleExpand)
    }}
    on:mouseenter={(e) => {
      clearTimeout(tooltipShowTimer)
      tooltipShowTimer = setTimeout(() => {
        openTooltip(e)
      }, 1000)
    }}
    on:mouseleave={() => {
      clearTimeout(tooltipShowTimer)
      closeTooltip()
    }}
  >
    {#if isShowTooltip}
      <div class="search-tooltip" style="top:{tooltipY}px; left:{tooltipX}px;">{path}</div>
    {/if}
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
  .search-result-item {
    position: relative;
  }

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
    padding-top: 0.25rem;
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

  .search-tooltip {
    position: fixed;
    z-index: 1000;
    top: 0;
    left: 0;
    background: #181818;
    color: #9d9d9d;
    padding: 4px 8px;
    font-weight: 400;
    font-size: 12px;
    white-space: nowrap;
    /* border-radius: 4px; */
    border: 1px solid #9d9d9d;
    transform: translateY(-4px);
  }
</style>
