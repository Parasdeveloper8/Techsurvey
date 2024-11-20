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

function ordinal(rank) {
    // Special cases for numbers ending in 11, 12, 13
    if (rank % 100 >= 11 && rank % 100 <= 13) {
        return rank + "th";
    }
    // Determine suffix based on last digit
    switch (rank % 10) {
        case 1: return rank + "st";
        case 2: return rank + "nd";
        case 3: return rank + "rd";
        default: return rank + "th";
    }
}

function createLeaderboard(leaderboardData) {
    const container = document.getElementById("leaderboard-container");

    // Clear any existing content in the container
    container.innerHTML = "";

    // Create a table element
    const table = document.createElement("table");
    table.style.width = "100%";
    table.style.borderCollapse = "collapse";

    // Create the header row with proper alignment
    const headerRow = document.createElement("tr");
    const headers = ["Rank", "Username", "Points"];
    headers.forEach(headerText => {
        const th = document.createElement("th");
        th.textContent = headerText;
        th.style.padding = "12px";
        th.style.backgroundColor = "#0066cc";
        th.style.color = "white";
        th.style.textAlign = "center"; // Center align header text
        th.style.fontSize = "16px";
        th.style.border = "1px solid #ddd";
        headerRow.appendChild(th);
    });
    table.appendChild(headerRow);

    // Add data rows to the table
    leaderboardData.forEach((item, index) => {
        const row = document.createElement("tr");

        // Rank column with ordinal suffix
        const rankCell = document.createElement("td");
        rankCell.textContent = ordinal(index + 1); // Use the ordinal function
        rankCell.style.padding = "12px";
        rankCell.style.textAlign = "center"; // Align center for rank
        rankCell.style.border = "1px solid #ddd";
        row.appendChild(rankCell);

        // Username column
        const usernameCell = document.createElement("td");
        usernameCell.textContent = item.username; // Assuming each item has a username property
        usernameCell.style.padding = "12px";
        usernameCell.style.textAlign = "left"; // Align left for username
        usernameCell.style.border = "1px solid #ddd";
        row.appendChild(usernameCell);

        // Points column
        const pointsCell = document.createElement("td");
        pointsCell.textContent = item.points; // Assuming each item has a points property
        pointsCell.style.padding = "12px";
        pointsCell.style.textAlign = "center"; // Align center for points
        pointsCell.style.border = "1px solid #ddd";
        row.appendChild(pointsCell);

        table.appendChild(row);
    });

    // Add the table to the container
    container.appendChild(table);

    // Add styling for the leaderboard container
    container.style.margin = "20px auto";
    container.style.padding = "20px";
    container.style.maxWidth = "800px";
    container.style.borderRadius = "8px";
    container.style.boxShadow = "0 0 10px rgba(0, 0, 0, 0.1)";
    container.style.backgroundColor = "#f9f9f9";
}
