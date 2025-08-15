export async function handleEnter(e, callback) {
  if (e.key === 'Enter') {
    e.preventDefault()
    await callback()
  }
}

export function debounce(func, delay) {
  let timeoutId
  return function (...args) {
    clearTimeout(timeoutId)
    timeoutId = setTimeout(() => func.apply(this, args), delay)
  }
}
