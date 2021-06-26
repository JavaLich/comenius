function initialize(){
}

// Redirects to correct login page
function login(loginPage){
    window.location.href = `/login.html?type=${loginPage}`;
}

// Function for initializing on page load
initialize();