<script>
  import { onMount, onDestroy } from 'svelte'
  import { EditorView, basicSetup } from 'codemirror'
  import { markdown } from '@codemirror/lang-markdown'
  import { oneDark } from '@codemirror/theme-one-dark'
  import { keymap } from '@codemirror/view'

  // Props
  let {
    value = '',
    placeholder = 'Write your markdown here...',
    onSave = () => {},
    onChange = () => {},
  } = $props()

  let editorElement
  let editorView

  onMount(() => {
    // Create the editor view
    editorView = new EditorView({
      doc: value,
      extensions: [
        basicSetup,
        markdown(),
        oneDark,
        keymap.of([
          {
            key: 'Ctrl-s',
            mac: 'Cmd-s',
            run: () => {
              onSave()
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
      ],
      parent: editorElement,
    })
  })

  onDestroy(() => {
    if (editorView) {
      editorView.destroy()
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
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.9rem;
  }

  :global(.cm-focused) {
    outline: none;
  }
</style>
