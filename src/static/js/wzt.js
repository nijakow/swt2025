
function addZettelToWarenkorb(zettelId) {
    // fetch("/api/add?id=" + zettelId, { method: 'POST' });
    /*
     * Diese Zeilen wurden mithilfe von GitHub Copilot generiert.
     */
    const currentUrl = window.location.href;
    const redirectUrl = encodeURIComponent(currentUrl);
    window.location.href = `/api/add?id=${zettelId}`;
}

function removeZettelFromWarenkorb(zettelId) {
    const currentUrl = window.location.href;
    const redirectUrl = encodeURIComponent(currentUrl);
    window.location.href = `/api/remove?id=${zettelId}`;
}
