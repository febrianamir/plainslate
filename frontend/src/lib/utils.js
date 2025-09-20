export async function handleEnter(e, callback) {
  if (e.key === 'Enter') {
    e.preventDefault()
    await callback()
  }
}

export function debounce(func, delay) {
  let timeoutId
  return function(...args) {
    clearTimeout(timeoutId)
    timeoutId = setTimeout(() => func.apply(this, args), delay)
  }
}

export function parseFilepath(filepath) {
  const lastDotIndex = filepath.lastIndexOf('.');
  const lastSlashIndex = Math.max(filepath.lastIndexOf('/'), filepath.lastIndexOf('\\'));

  if (lastDotIndex === -1 || lastDotIndex < lastSlashIndex) {
    // No extension found
    return {
      name: filepath.substring(lastSlashIndex + 1),
      extension: ''
    };
  }

  const dirpath = filepath.substring(0, lastSlashIndex + 1)
  const filename = filepath.substring(lastSlashIndex + 1, lastDotIndex);
  const extension = filepath.substring(lastDotIndex + 1);

  return { dirpath: dirpath, name: filename, extension: extension };
}
