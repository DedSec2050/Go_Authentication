# Fetch-based Login Script for JWT Authentication

This README explains the functionality of a JavaScript snippet designed to send a login request to a backend server using **fetch** and process the response containing a **JWT token**.

---

## **Overview**

This code demonstrates how to:
1. Send a POST request to an authentication endpoint.
2. Provide **email** and **password** as login credentials.
3. Handle the server's response and extract the **JWT token** upon successful authentication.
4. Log and store the JWT token in the browser's local storage for future use.

---

## **Prerequisites**

1. A backend server running at `http://localhost:8080` with a login endpoint at `/api/auth/login`.
2. The backend must:
   - Accept `email` and `password` in the request body as JSON.
   - Respond with a valid JWT token in a JSON object (e.g., `{"token": "jwt-token-here"}`) on successful authentication.
   - Return an appropriate error response for invalid credentials.

---

## **Code Explanation**

### **Code Block**
```javascript
// User credentials
const email = "test@example.com";
const password = "password";

// Send POST request to login endpoint
fetch("http://localhost:8080/api/auth/login", {
    method: "POST",  // HTTP method for login
    headers: {
        "Content-Type": "application/json",  // Sending data as JSON
    },
    body: JSON.stringify({
        email: email,  // Email from user input
        password: password  // Password from user input
    })
})
    .then(response => {
        if (!response.ok) {
            throw new Error("Login failed. Invalid credentials or error on the server.");
        }
        return response.json();  // Parse JSON response
    })
    .then(data => {
        console.log("Login successful. JWT Token:", data.token);  // Access the JWT token from response
        // Store the token for future 
