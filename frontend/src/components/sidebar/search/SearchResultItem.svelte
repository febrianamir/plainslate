<script>
  import { ChevronDown, ChevronRight } from 'lucide-svelte'
  import { handleEnter } from '../../../../src/lib/utils'
  import { onDestroy, onMount } from 'svelte'
  import { openFile } from '../../../lib/services/fileService'

  let { filename, path, fileMatches, query } = $props()

  let clientX = 0
  let clientY = 0

  let isExpanded = $state(true)
  let isShowTooltip = $state(false)
  let tooltipX = $state(0)
  let tooltipY = $state(0)
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

  // Function to split line into parts with matched query highlighted
  function highlightMatch(line, query) {
    if (!query) return [{ text: line, match: false }]

    const regex = new RegExp(`(${escapeRegExp(query)})`, 'gi')
    let parts = []
    let lastIndex = 0
    let match

    while ((match = regex.exec(line)) !== null) {
      if (match.index > lastIndex) {
        parts.push({ text: line.slice(lastIndex, match.index), match: false })
      }
      parts.push({ text: match[0], match: true })
      lastIndex = match.index + match[0].length
    }
    if (lastIndex < line.length) {
      parts.push({ text: line.slice(lastIndex), match: false })
    }
    return parts
  }

  // Escape regex special chars in query
  function escapeRegExp(string) {
    return string.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')
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
    role="button"
    tabindex="0"
    onclick={() => {
      clearTimeout(tooltipShowTimer)
      toggleExpand()
    }}
    onkeydown={(e) => {
      clearTimeout(tooltipShowTimer)
      handleEnter(e, toggleExpand)
    }}
    onmouseenter={(e) => {
      clearTimeout(tooltipShowTimer)
      tooltipShowTimer = setTimeout(() => {
        openTooltip(e)
      }, 1000)
    }}
    onmouseleave={() => {
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
    <div class="search-result-filename">
      {filename}
    </div>
    <div class="search-result-file-count">
      {fileMatches.length}
    </div>
  </div>
  <div class:active={isExpanded} class="search-result-matches">
    {#each fileMatches as match}
      <div
        role="button"
        tabindex="0"
        class="search-result-match"
        onclick={() => openFile(path)}
        onkeydown={(e) => {
          handleEnter(e, () => openFile(path))
        }}
      >
        {#each highlightMatch(match, query) as part}
          {#if part.match}
            <span class="search-result-highlight">{part.text}</span>
          {:else}
            {part.text}
          {/if}
        {/each}
      </div>
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

  .search-result-filename {
    flex: 1;
  }

  .search-result-file-count {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 18px;
    height: 18px;
    font-weight: 400;
    font-size: 0.65rem;
    border-radius: 20px;
    background: #2d2d2d;
  }

  .search-result-matches {
    display: none;
  }

  .search-result-matches.active {
    display: block;
  }

  .search-result-match {
    cursor: pointer;
    padding: 0.15rem 0;
    padding-left: 0.75rem;
  }

  .search-result-match:hover {
    background-color: #2b2b2b;
  }

  .search-result-highlight {
    background-color: #c9e6c1;
    color: #1e1e1e;
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
    border: 1px solid #9d9d9d;
    transform: translateY(-4px);
  }
</style>
