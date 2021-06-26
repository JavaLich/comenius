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

initialize();