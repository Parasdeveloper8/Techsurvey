// Function to fetch data from the API and render the graph
fetch('http://localhost:7300/getUserNumber')  // Update with your actual API endpoint
    .then(response => response.json())
    .then(data => {
        const emailCount = data.email_count;

        // Data for the chart
        const chartData = {
            labels: ['Distinct Emails'],  // Label for the bar
            datasets: [{
                label: 'Number of Distinct Emails',
                data: [emailCount],  // Data point for the graph
                backgroundColor: 'rgba(75, 192, 192, 0.2)',  // Bar color
                borderColor: 'rgba(75, 192, 192, 1)',  // Border color
                borderWidth: 1
            }]
        };

        // Chart options
        const chartOptions = {
            responsive: true,
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        };

        // Render the chart
        const ctx = document.getElementById('emailCountChart').getContext('2d');
        const emailCountChart = new Chart(ctx, {
            type: 'bar',  // Change to 'line', 'pie', etc., depending on your preference
            data: chartData,
            options: chartOptions
        });
    })
    .catch(error => console.error('Error fetching email count:', error));
