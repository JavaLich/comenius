document.getElementById("certificate-file-input").addEventListener('change', (event) => {
    console.log(event.target.value);
    let fileExtension = event.target.value.split('.').pop();
    if (!["gif", "jpeg", "jpg", "png", "pdf"].includes(fileExtension)) {
        event.target.files[0] = null;
        event.target.value = "";
        return;
    }
    document.getElementById("certificate-file-name-label").innerHTML = event.target.value.split("\\").pop();
})
