document.addEventListener("DOMContentLoaded", () => {
  const code = prompt("HI Admin Enter security code to proceed further");
  
  // Check the security code
  if (code !== "8890Aaaa@") {
    alert("Invalid security code. Redirecting...");
    window.location.href = "/";
    return; // Exit the function if the code is invalid
  }
  fetch("http://localhost:7300/getdata")
  .then(response => {
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    return response.json();
  })
  .then(data => {
    // Convert the data to a string and remove quotes and curly braces
    let cleanedData = JSON.stringify(data).replace(/["{}]/g, '');
  
    console.log(`Data received:`, cleanedData); // Log the cleaned data
    const container = document.getElementById("data-container");
    const ul = document.createElement("ul");
  
    // If data is an array, iterate over the cleaned data
    if (Array.isArray(data)) {
      data.forEach(item => {
        const li = document.createElement("li");
        li.textContent = JSON.stringify(item, null, 2).replace(/["{}]/g, ''); // Clean individual items
        ul.appendChild(li);
      });
    } else {
      // If data is not an array, handle it accordingly
      const li = document.createElement("li");
      li.textContent = cleanedData; // Use cleanedData for non-array data
      ul.appendChild(li);
    }
    
    container.appendChild(ul);
  })
  .catch(error => console.error('Error fetching data:', error));
  

  // Fetch feedback from /getfeedback
  fetch("http://localhost:7300/getfeedback")
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then(data => {
      console.log(`Feedback received:`, data);
      const container = document.getElementById("feedback");

      // Create table header
      const thead = document.createElement("thead");
      const headerRow = document.createElement("tr");
      ["Email", "Comment", "Time"].forEach(headerText => {
        const th = document.createElement("th");
        th.textContent = headerText;
        headerRow.appendChild(th);
      });
      
      thead.appendChild(headerRow);
      container.appendChild(thead);

      // Create table body
      const tbody = document.createElement("tbody");

      data.forEach(item => {
        const tr = document.createElement("tr");

        // Create a cell for each data item: email, comment, and time
        const emailCell = document.createElement("td");
        emailCell.textContent = item.emailofuser;
        tr.appendChild(emailCell);

        const commentCell = document.createElement("td");
        commentCell.textContent = item.comment;
        tr.appendChild(commentCell);

        const timeCell = document.createElement("td");
        timeCell.textContent = item.time;
        tr.appendChild(timeCell);

        tbody.appendChild(tr);
      });

      container.appendChild(tbody);
    })
    .catch(error => console.error('Error fetching feedback:', error));
});
