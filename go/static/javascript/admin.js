
document.addEventListener("keydown", function(event) {
    if (event.altKey && event.key === 'p') {
        event.preventDefault();
        const adminOption = document.getElementById("adminOption");
        const adminOverlay = document.getElementById("adminOverlay");
        const isVisible = adminOption.style.display === "block";
        adminOption.style.display = isVisible ? "none" : "block";
        adminOverlay.style.display = isVisible ? "none" : "block";
    }
});

function redirectToAdmin() {
    window.location.href = "/admin"; // Redirects to the admin page
}

function closeAdmin() {
    document.getElementById("adminOption").style.display = "none";
    document.getElementById("adminOverlay").style.display = "none";
}
