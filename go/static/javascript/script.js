document.getElementById('registerForm').addEventListener('submit', function(event) {
    const password = document.getElementById('password').value;
    const checkPassword = document.getElementById('checkPassword').value;
    const passwordError = document.getElementById('passwordError');
    const btn = document.getElementById('sbtn');
    
    if (password !== checkPassword) {
        event.preventDefault();  // Prevent form submission
        btn.disabled = true;
        passwordError.textContent = "Passwords do not match.";
        passwordError.style.display = "block";
    } else {
        // If passwords match, remove error and enable the button
        btn.disabled = false;
        passwordError.textContent = "";
        passwordError.style.display = "none";
    }
});

