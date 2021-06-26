function setLoginType(type) {
    if (type === "Learner") {
        document.getElementById("learner-selector").classList.add("selected")
        document.getElementById("contributor-selector").classList.remove("selected")
    } else if (type === "Contributor") {
        document.getElementById("contributor-selector").classList.add("selected")
        document.getElementById("learner-selector").classList.remove("selected")
    }
}