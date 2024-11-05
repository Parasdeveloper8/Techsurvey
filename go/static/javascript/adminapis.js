document.addEventListener("DOMContentLoaded", () => {
    fetch("http://localhost:6900/getdata")
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => {
        console.log(`Data received:`, data);
        const container = document.getElementById("data-container");
        const ul = document.createElement("ul");

        for (let i = 0; i < data.length; i++) {
          const li = document.createElement("li");
          li.textContent = JSON.stringify(data[i], null, 2); // Prettify JSON output
          ul.appendChild(li);
        }
        
        container.appendChild(ul);
      })
      .catch(error => console.error('Error fetching data:', error));
  });

  
document.addEventListener("DOMContentLoaded", () => {
    fetch("http://localhost:6900/getfeedback")
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => {
        console.log(`Data received:`, data);
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
      .catch(error => console.error('Error fetching data:', error));
  });