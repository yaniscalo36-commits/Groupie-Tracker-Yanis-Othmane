
// Recherche et filtres dynamiques


const searchInput = document.getElementById("search")
const suggestionsBox = document.getElementById("suggestions")
const artistCards = document.querySelectorAll(".artist-card")

// Recherche dynamique (client-side)
if (searchInput) {
  searchInput.addEventListener("input", () => {
    const query = searchInput.value.toLowerCase().trim()
    suggestionsBox.innerHTML = ""

    if (query.length < 2) return

    let count = 0

    artistCards.forEach(card => {
      const name = card.dataset.name.toLowerCase()

      if (name.includes(query) && count < 5) {
        const div = document.createElement("div")
        div.className = "suggestion"
        div.textContent = card.dataset.name

        div.addEventListener("click", () => {
          searchInput.value = card.dataset.name
          suggestionsBox.innerHTML = ""
        })

        suggestionsBox.appendChild(div)
        count++
      }
    })
  })
}


// Tri dynamique des artistes

function sortArtists(type) {
  const container = document.querySelector(".artists-grid")
  if (!container) return

  const cards = Array.from(container.children)

  cards.sort((a, b) => {
    let A = a.dataset[type]
    let B = b.dataset[type]

    // Tri numérique
    if (type === "year" || type === "members") {
      return Number(A) - Number(B)
    }

    // Tri alphabétique
    return A.localeCompare(B)
  })

  cards.forEach(card => container.appendChild(card))
}


// Réinitialisation des filtres

function resetAll() {
  window.location.href = "/artists"
}
