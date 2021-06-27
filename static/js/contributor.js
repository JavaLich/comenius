
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

function donate(){
    let selectCert = document.getElementById("cert-select");
    let strCert = selectCert.options[selectCert.selectedIndex].text;
    let certID = selectCert.options[selectCert.selectedIndex].cert_id;

    let selectBank = document.getElementById("bank-select");
    let strBank = selectBank.options[selectBank.selectedIndex].text;

    let amount = document.getElementById("amount").value;

    let user = getCookie("user");

    let url = "../donate"
    data = {
        "certificate": strCert,
        "bank": strBank,
        "amount": amount,
        "user": user,
        "certID": certID
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
    .then(result =>{
        console.log(result);
    });
}