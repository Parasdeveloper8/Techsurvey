document.addEventListener("DOMContentLoaded", () => {
    fetch("http://localhost:7300/getupdate")
      .then(response => response.json())
      .then(data => {
        console.log(data); // Check structure of data
        for (let i = 0; i < data.length; i++) {
          // Create a container for the message with styling
          const messageContainer = document.createElement("div");
          messageContainer.style.border = "1px solid #ddd";
          messageContainer.style.borderRadius = "8px";
          messageContainer.style.padding = "15px";
          messageContainer.style.margin = "10px auto";
          messageContainer.style.maxWidth = "500px";
          messageContainer.style.boxShadow = "0px 4px 8px rgba(0, 0, 0, 0.1)";
          messageContainer.style.backgroundColor = "#f9f9f9";
  
          // Create the message element with bold font
          const messageElement = document.createElement("p");
          messageElement.textContent = data[i].message || "No message available";
          messageElement.style.fontSize = "18px";
          messageElement.style.fontWeight = "600";
          messageElement.style.color = "#333";
          messageElement.style.marginBottom = "5px";
  
          // Create the time element in smaller, lighter font
          const timeElement = document.createElement("small");
          timeElement.textContent = data[i].time || "No time available";
          timeElement.style.fontSize = "14px";
          timeElement.style.color = "#777";
  
          // Append message and time elements to the container
          messageContainer.appendChild(messageElement);
          messageContainer.appendChild(timeElement);
  
          // Append the container to the body
          document.body.appendChild(messageContainer);
        }
      })
      .catch(error => console.error("Error:", error));
  });
  
  
  