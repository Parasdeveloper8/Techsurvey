const loader = document.getElementById('loader');
loader.style.display = 'block';
document.addEventListener("DOMContentLoaded", () => {
    const pointsContainer = document.createElement("div"); // Create a container for the points
    pointsContainer.style.textAlign = "center";
    pointsContainer.style.fontSize = "20px"; // Set some basic styling
    pointsContainer.style.marginTop = "20px";
    pointsContainer.style.fontWeight = "bold";
    pointsContainer.textContent = "Fetching your points..."; // Initial text before fetching
    document.body.appendChild(pointsContainer); // Add the container to the body

    // Fetch the points from the server
    fetch("http://localhost:4700/getpoints")
        .then((response) => {
            if (!response.ok) {
                throw new Error("Failed to fetch points");
            }
            return response.json();
        })
        .then((data) => {
            if (data.tables_with_email !== undefined) {
                pointsContainer.textContent = `Your Points are ${data.tables_with_email}`;
                loader.style.display = 'none';
            } else {
                pointsContainer.textContent = "Could not retrieve your points.";
            }
        })
        .catch((error) => {
            console.error("Error fetching points:", error);
            pointsContainer.textContent = "Error fetching your points. Please try again later.";
        });
});
