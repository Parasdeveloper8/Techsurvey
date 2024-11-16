const loader = document.getElementById('loader');
loader.style.display = 'block';

document.addEventListener("DOMContentLoaded", () => {
    // Fetch leaderboard data from the server
    fetch("http://localhost:7300/leaderboarddata")
        .then((response) => response.json())
        .then((data) => {
            if (data && Array.isArray(data)) {
                // Call the function to display leaderboard data
                createLeaderboard(data);
                loader.style.display = 'none'; // Hide loader once data is loaded
            } else {
                console.error("Leaderboard data not found.");
            }
        })
        .catch((error) => {
            console.error("Error fetching leaderboard data:", error);
        });
});

function createLeaderboard(leaderboardData) {
    const container = document.getElementById("leaderboard-container");

    // Create a table element
    const table = document.createElement("table");
    table.style.width = "100%";
    table.style.borderCollapse = "collapse";
    table.style.fontFamily = "Arial, sans-serif";

    // Create the header row
    const headerRow = document.createElement("tr");
    const headers = ["Rank", "Username", "Points"];
    headers.forEach(headerText => {
        const th = document.createElement("th");
        th.textContent = headerText;
        th.style.padding = "10px";
        th.style.backgroundColor = "#4CAF50";
        th.style.color = "white";
        th.style.textAlign = "left";
        th.style.fontSize = "16px";
        th.style.border = "1px solid #ddd";
        headerRow.appendChild(th);
    });
    table.appendChild(headerRow);

    // Add data rows to the table
    leaderboardData.forEach((item, index) => {
        const row = document.createElement("tr");

        // Rank column
        const rankCell = document.createElement("td");
        rankCell.textContent = index + 1;
        rankCell.style.padding = "8px";
        rankCell.style.textAlign = "center";
        rankCell.style.border = "1px solid #ddd";
        row.appendChild(rankCell);

        // Username column
        const usernameCell = document.createElement("td");
        usernameCell.textContent = item.username; // Assuming each item has a username property
        usernameCell.style.padding = "8px";
        usernameCell.style.border = "1px solid #ddd";
        row.appendChild(usernameCell);

        // Points column
        const pointsCell = document.createElement("td");
        pointsCell.textContent = item.points; // Assuming each item has a points property
        pointsCell.style.padding = "8px";
        pointsCell.style.textAlign = "center";
        pointsCell.style.border = "1px solid #ddd";
        row.appendChild(pointsCell);

        table.appendChild(row);
    });

    // Add the table to the container
    container.appendChild(table);

    // Add some basic styling for the leaderboard container
    container.style.margin = "20px auto";
    container.style.padding = "20px";
    container.style.maxWidth = "800px";
    container.style.borderRadius = "8px";
    container.style.boxShadow = "0 0 10px rgba(0, 0, 0, 0.1)";
    container.style.backgroundColor = "#f9f9f9";
    container.style.textAlign = "center";
}
