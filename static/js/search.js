const input = document.getElementById("search")
const box = document.getElementById("suggestions")

input.addEventListener("input", () => {
  const q = input.value.trim()
  if (q.length < 2) {
    box.innerHTML = ""
    return
  }

  fetch("/artists?search=" + q)
    .then(r => r.text())
    .then(t => {
      const res = []
      const parts = t.split("<h3>")

      for (let i = 1; i < parts.length && res.length < 5; i++) {
        const name = parts[i].split("</h3>")[0]
        res.push(name)
      }

      box.innerHTML = res.map(n =>
        `<div class="suggestion" data-name="${n}">${n}</div>`
      ).join("")
    })
})

box.addEventListener("click", e => {
  if (e.target.dataset.name) {
    input.value = e.target.dataset.name
    box.innerHTML = ""
  }
})
