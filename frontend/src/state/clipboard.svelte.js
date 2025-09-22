export let clipboard = $state({
  clipboardType: '', // CUT, COPY
  path: '', // Clipboard path
  filename: '', // Clipboard filename
  node: {}, // Clipboard node
})

export function cleanClipboard() {
  clipboard.clipboardType = ''
  clipboard.path = ''
  clipboard.filename = ''
  clipboard.node = {}
}
