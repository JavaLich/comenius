function loadContributorDetails() {
    var url_string = window.location.href;
    let username = url_string.split("/").pop()
    fetch(`/contributor_details?username=${username}`)
        .then(response => response.json())
        .then(data => {
            // Donation history
            donationHistory = data["ContributionList"]
            for (let i = donationHistory.length - 1; i >= 0; i--) {
                let entry = `
                            <tr>
                              <td>${donationHistory[i]["Date"].substring(0, 9)}</td>
                              <td>$${(donationHistory[i]["Amount"] / 100).toFixed(2)}</td>
                            </tr>`;
                document.getElementById("payment-table").insertAdjacentHTML("beforeend", entry);
            }

            // Total impact
            document.getElementById("total-impact").innerHTML = data["TotalImpact"] + " people";

            // Total money raised
            document.getElementById("amount-raised").innerHTML = "$" + (data["TotalMoneyRaised"] / 100).toFixed(2);

            // People Impacted
            let peopleImpacted = data["PeopleImpacted"];
            let keys = Object.keys(peopleImpacted)
            for (let i = 0; i < keys.length; i++) {
                entry = `
                <tr>
                    <td>${keys[i]}</td>
                    <td>$${(peopleImpacted[keys[i]]/100).toFixed(2)}</td>
                </tr>
                `;
                document.getElementById("impact-table").insertAdjacentHTML("beforeend", entry);
            }
        })
}

loadContributorDetails()

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

function donate() {
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
