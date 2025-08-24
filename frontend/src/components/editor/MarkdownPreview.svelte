<script>
  import { marked } from 'marked'
  import { markedHighlight } from 'marked-highlight'
  import hljs from 'highlight.js'
  import 'highlight.js/styles/atom-one-dark.css'

  let { content = '', className = '' } = $props()

  // Configure marked with syntax highlighting
  marked.use(
    markedHighlight({
      langPrefix: 'hljs language-',
      highlight(code, lang) {
        const language = hljs.getLanguage(lang) ? lang : 'plaintext'
        return hljs.highlight(code, { language }).value
      },
    })
  )

  // Configure marked options
  marked.setOptions({
    breaks: true, // Convert \n to <br>
    gfm: true, // GitHub Flavored Markdown
    headerIds: true,
    mangle: false,
    sanitize: false,
  })

  // Convert markdown to HTML
  let htmlContent = $derived(marked.parse(content || ''))
</script>

<div class="markdown-preview {className}">
  {@html htmlContent}
</div>

<style>
  .markdown-preview {
    padding: 2rem;
    font-family: -apple-system, BlinkMacSystemFont, 'Inter', 'Roboto', sans-serif;
    line-height: 1.6;
    color: #e0e0e0;
    background: #1e1e1e;
    height: 100%;
    overflow-y: scroll;
    user-select: text;
    -webkit-user-select: text;
  }

  /* Scrollbar styling */
  .markdown-preview::-webkit-scrollbar {
    width: 8px;
  }

  .markdown-preview::-webkit-scrollbar-track {
    background: transparent;
  }

  .markdown-preview::-webkit-scrollbar-thumb {
    background-color: rgba(255, 255, 255, 0.2);
    border-radius: 4px;
  }

  .markdown-preview::-webkit-scrollbar-thumb:hover {
    background-color: rgba(255, 255, 255, 0.3);
  }

  /* Typography */
  :global(.markdown-preview h1) {
    font-size: 2rem;
    margin: 1.5rem 0 0.5rem 0;
    padding-bottom: 0.3rem;
    border-bottom: 1px solid #404040;
    color: #ffffff;
  }

  :global(.markdown-preview h2) {
    font-size: 1.5rem;
    margin: 1.2rem 0 0.4rem 0;
    color: #ffffff;
    border-bottom: 1px solid #333333;
    padding-bottom: 0.2rem;
  }

  :global(.markdown-preview h3) {
    font-size: 1.25rem;
    margin: 1rem 0 0.3rem 0;
    color: #ffffff;
  }

  :global(.markdown-preview h4),
  :global(.markdown-preview h5),
  :global(.markdown-preview h6) {
    margin: 0.8rem 0 0.3rem 0;
    color: #ffffff;
  }

  :global(.markdown-preview p) {
    margin: 0.8rem 0;
  }

  :global(.markdown-preview a) {
    color: #6da96f;
    text-decoration: none;
  }

  :global(.markdown-preview a:hover) {
    text-decoration: underline;
  }

  /* Code blocks */
  :global(.markdown-preview pre) {
    background: #282828;
    border: 1px solid #404040;
    border-radius: 6px;
    padding: 1rem;
    margin: 1rem 0;
    overflow-x: auto;
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.85rem;
  }

  :global(.markdown-preview code) {
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.85rem;
  }

  :global(.markdown-preview p code) {
    background: #2a2a2a;
    padding: 0.2rem 0.4rem;
    border-radius: 3px;
    border: 1px solid #404040;
  }

  :global(.markdown-preview pre code.hljs) {
    background: transparent !important;
    padding: 0;
  }

  /* Lists */
  :global(.markdown-preview ul),
  :global(.markdown-preview ol) {
    margin: 0.8rem 0;
    padding-left: 2rem;
  }

  :global(.markdown-preview li) {
    margin: 0.3rem 0;
  }

  /* Blockquotes */
  :global(.markdown-preview blockquote) {
    margin: 1rem 0;
    padding: 0 1rem;
    border-left: 4px solid #6da96f;
    background: #252525;
    color: #b0b0b0;
  }

  /* Tables */
  :global(.markdown-preview table) {
    border-collapse: collapse;
    width: 100%;
    margin: 1rem 0;
  }

  :global(.markdown-preview th),
  :global(.markdown-preview td) {
    border: 1px solid #404040;
    padding: 0.5rem 1rem;
    text-align: left;
  }

  :global(.markdown-preview th) {
    background: #2a2a2a;
    font-weight: 600;
  }

  :global(.markdown-preview tr:nth-child(even)) {
    background: #252525;
  }

  /* Images */
  :global(.markdown-preview img) {
    max-width: 100%;
    height: auto;
    border-radius: 4px;
    margin: 0.5rem 0;
  }

  /* Horizontal rules */
  :global(.markdown-preview hr) {
    border: none;
    height: 1px;
    background: #404040;
    margin: 2rem 0;
  }

  /* Task lists (GitHub flavored markdown) */
  :global(.markdown-preview input[type='checkbox']) {
    margin-right: 0.5rem;
  }

  /* Empty state */
  .markdown-preview:empty::before {
    content: 'Start writing markdown to see the preview...';
    color: #666;
    font-style: italic;
  }
</style>
