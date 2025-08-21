<script>
  import { onMount, onDestroy } from 'svelte'
  import { EditorView, basicSetup } from 'codemirror'
  import { markdown } from '@codemirror/lang-markdown'
  import { oneDark } from '@codemirror/theme-one-dark'
  import { keymap } from '@codemirror/view'
  import { languages } from '@codemirror/language-data'
  import { StateEffect } from '@codemirror/state'

  let {
    value = '',
    placeholder = 'Write your markdown here...',
    onSave = () => {},
    onChange = () => {},
  } = $props()

  let editorElement
  let editorView
  let isLineWrapActive = $state(true)
  let currentExtensions = []

  function createExtensions() {
    return [
      basicSetup,
      ...(isLineWrapActive ? [EditorView.lineWrapping] : []),
      markdown({
        codeLanguages: languages,
      }),
      oneDark,
      EditorView.theme({
        '.cm-scroller': {
          '&::-webkit-scrollbar': {
            width: '8px',
            height: '8px',
          },
          '&::-webkit-scrollbar-track': {
            background: 'transparent',
          },
          '&::-webkit-scrollbar-corner': {
            background: 'transparent',
          },
          '&::-webkit-scrollbar-thumb': {
            backgroundColor: 'rgba(255, 255, 255, 0.2)',
            borderRadius: '4px',
          },
          '&::-webkit-scrollbar-thumb:hover': {
            backgroundColor: 'rgba(255, 255, 255, 0.3)',
          },
        },
      }),
      keymap.of([
        {
          key: 'Ctrl-s', // Save file
          mac: 'Cmd-s',
          run: () => {
            onSave()
            return true
          },
        },
        {
          key: 'Alt-z', // Toggle line wrap
          run: (view) => {
            console.log('wrapped')
            isLineWrapActive = !isLineWrapActive
            return true
          },
        },
      ]),
      EditorView.updateListener.of((update) => {
        if (update.docChanged) {
          const newValue = update.state.doc.toString()
          onChange(newValue)
        }
      }),
    ]
  }

  onMount(() => {
    currentExtensions = createExtensions()

    // Create the editor view
    editorView = new EditorView({
      doc: value,
      extensions: currentExtensions,
      parent: editorElement,
    })
  })

  onDestroy(() => {
    if (editorView) {
      editorView.destroy()
    }
  })

  $effect(() => {
    if (editorView) {
      const newExtensions = createExtensions()
      editorView.dispatch({
        effects: StateEffect.reconfigure.of(newExtensions),
      })
      currentExtensions = newExtensions
    }
  })

  // Update editor content when value prop changes
  $effect(() => {
    if (editorView && value !== editorView.state.doc.toString()) {
      editorView.dispatch({
        changes: {
          from: 0,
          to: editorView.state.doc.length,
          insert: value,
        },
      })
    }
  })
</script>

<div bind:this={editorElement} class="codemirror-wrapper"></div>

<style>
  .codemirror-wrapper {
    flex: 1;
    overflow: auto;
  }

  :global(.cm-editor) {
    height: 100%;
  }

  :global(.cm-scroller) {
    padding-bottom: 7rem;
  }

  :global(.cm-content) {
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.85rem;
  }

  :global(.cm-focused) {
    outline: none;
  }
</style>
