document.getElementById('registerForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const password = document.getElementById('password').value;
    const checkPassword = document.getElementById('checkPassword').value;
    const passwordError = document.getElementById('passwordError');

    if (password !== checkPassword) {
        passwordError.textContent = "Passwords do not match.";
        passwordError.style.display = "block";
    } else {
        passwordError.style.display = "none";
        alert("Registration successful!");
        // Here you can add code to submit the form or send the data to the server
    }
});
