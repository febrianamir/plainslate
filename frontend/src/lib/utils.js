export async function handleEnter(e, callback) {
  if (e.key === 'Enter') {
    e.preventDefault()
    await callback()
  }
}
