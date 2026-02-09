function toggleTheme() {
  document.body.classList.toggle("theme-purple")

  if (document.body.classList.contains("theme-purple")) {
    localStorage.setItem("theme", "purple")
  } else {
    localStorage.setItem("theme", "green")
  }
}

// Appliquer le thème au chargement
document.addEventListener("DOMContentLoaded", () => {
  const theme = localStorage.getItem("theme")
  if (theme === "purple") {
    document.body.classList.add("theme-purple")
  }
})
