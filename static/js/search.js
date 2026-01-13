const input = document.getElementById("search");
const suggestions = document.getElementById("suggestions");

input.addEventListener("input", () => {
    const q = input.value.trim();
    if (q.length < 2) {
        suggestions.innerHTML = "";
        return;
    }

    fetch("/artists?search=" + encodeURIComponent(q))
        .then(res => res.text())
        .then(html => {
            const names = [...html.matchAll(/<h3>(.*?)<\/h3>/g)]
                .map(m => m[1])
                .slice(0, 5);

            suggestions.innerHTML = names.map(n =>
                `<div class="suggestion" onclick="selectSuggestion('${n}')">${n}</div>`
            ).join("");
        });
});

function selectSuggestion(name) {
    input.value = name;
    suggestions.innerHTML = "";
}
