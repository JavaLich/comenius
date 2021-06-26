function initialize(){
    var url_string = window.location.href;
    var url = new URL(url_string);
    var type = url.searchParams.get("type");
    if (type === "learner") {
        document.getElementById("learner-selector").classList.add("selected")
        document.getElementById("contributor-selector").classList.remove("selected")
    } else if (type === "contributor") {
        document.getElementById("contributor-selector").classList.add("selected")
        document.getElementById("learner-selector").classList.remove("selected")
    }
}

function setLoginType(type) {
    window.location.href = "./login.html?type="+type;
    if (type === "learner") {
        document.getElementById("learner-selector").classList.add("selected")
        document.getElementById("contributor-selector").classList.remove("selected")
    } else if (type === "contributor") {
        document.getElementById("contributor-selector").classList.add("selected")
        document.getElementById("learner-selector").classList.remove("selected")
    }
}

function login(){
    var url_string = window.location.href;
    var cur_url = new URL(url_string);
    var type = cur_url.searchParams.get("type");
    let username = document.getElementById("username");
    let password = document.getElementById("password");
    let url = "./login"
    data = {
        "username": username,
        "password": password,
        "type": type
    }
    fetch(url, {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
          'Content-Type': 'application/json'
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify(data) // body data type must match "Content-Type" header
    }).then(response =>response.json())
    .then(result =>console.log(result))
}

initialize();