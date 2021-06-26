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

// History
function getDateString(daysBeforeToday) {
    const date = new Date(); // defaults to today
    date.setDate(date.getDate() - daysBeforeToday);
    return date.toDateString();
}

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
            data: [12, 19, 3, 5, 2, 3, 6],
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