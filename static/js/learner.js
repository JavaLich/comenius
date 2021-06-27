// Certificate File Input
document.getElementById("certificate-file-input").addEventListener('change', (event) => {
    console.log(event.target.value);
    let fileExtension = event.target.value.split('.').pop();
    if (!["jpeg", "jpg", "png", "pdf"].includes(fileExtension)) {
        event.target.files[0] = null;
        event.target.value = "";
        document.getElementById("certificate-file-name-label").innerHTML = "File extension must be .jpeg, .jpg, .png, or .pdf";
        return;
    }
    document.getElementById("certificate-file-name-label").innerHTML = event.target.value.split("\\").pop();
})

// Actively Funding
function loadLearnerDetails() {
    var url_string = window.location.href;
    let username = url_string.split("/").pop()
    fetch(`/learner_details?username=${username}`)
        .then(response => response.json())
        .then(data => {
            // Weekly money raised
            document.getElementById("weekly-contributions").innerText = "$" + (data["MoneyRaisedWeek"]/100).toFixed(2)

            // Total money raised
            document.getElementById("total-contributions").innerText = "$" + (data["TotalContributionsReceived"]/100).toFixed(2)

            // Contribution history
            setUpHistory(data["ContributionHistory"].reverse())

            // Certificate data
            certData = data["CertificateList"];
            for (let i = 0; i < certData.length; i++) {
                let color = "red";
                if (100 * certData[i]["raisedAmount"]/certData[i]["price"] > 80) {
                    color = "green";
                }
                else if (100 * certData[i]["raisedAmount"]/certData[i]["price"] > 30) {
                    color = "orange";
                }
                let entry = `
                <div class="active-course-listing">
                    <img src="${certData[i]["courseImageURL"]}"/>
                    <div class="course-info">
                        <h3>${certData[i]["name"]}</h3>
                        <p>${certData[i]["platform"]}</p>
                        <div class="meter ${color} nostripes">
                            <span style="width: ${100 * certData[i]["raisedAmount"]/certData[i]["price"]}%"></span>
                        </div>
                        $${(certData[i]["raisedAmount"]/100).toFixed(2)} funded out of $${(certData[i]["price"]/100).toFixed(2)} 
                    </div>
                </div>`;
                document.getElementById("active-listing-section").insertAdjacentHTML("beforeend", entry);
            }
        })
}

loadLearnerDetails()

// History
function getDateString(daysBeforeToday) {
    const date = new Date(); // defaults to today
    date.setDate(date.getDate() - daysBeforeToday);
    return date.toDateString();
}

function setUpHistory(data) {
    var labelList = [];
    for (var i = 7; i > 0; i--) {
        labelList.push(getDateString(i));
    }

    var ctx = document.getElementById('myChart').getContext('2d');
    var myChart = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: labelList,
            datasets: [{
                label: 'Contributions Received ($)',
                data: data,
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(255, 206, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(255, 159, 64, 0.2)',
                    'rgba(38, 158, 191, 0.2)'
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)',
                    'rgba(255, 206, 86, 1)',
                    'rgba(75, 192, 192, 1)',
                    'rgba(153, 102, 255, 1)',
                    'rgba(255, 159, 64, 1)',
                    'rgba(38, 158, 191, 1)'
                ],
                borderWidth: 1
            }]
        },
        options: {
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });
}

function uploadCert() {
    let course_url = document.getElementById("course-url").value;
    let fileElement = document.getElementById("certificate-file-input");
    let formData = new FormData();
    formData.append("certificate_file", fileElement.files[0]);
    formData.append("course_url", course_url)
    fetch("../certificate", {
        body: formData,
        method: "post"
    }).then(response => response.json())
        .then(result => { console.log(result) })
}