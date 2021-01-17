export function uniq(array) {
  const temp = []
  const index = []
  const l = array.length
  for (let i = 0; i < l; i++) {
    for (let j = i + 1; j < l; j++) {
      if (array[i] === array[j]) {
        i++
        j = i
      }
    }
    temp.push(array[i])
    index.push(i)
  }
  // console.log(index)
  return temp
}
