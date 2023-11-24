function getRandomRover() {
    fetch('/api/randomRover', {
        method: 'GET',
    })
    .then(response => response.json())
    .then(data => {
        console.log('Random Rover:', data);
        alert('Random Rover generated! Check the console for details.');
    })
    .catch(error => {
        console.error('Error getting random rover:', error);
        alert('Error getting random rover. Check the console for details.');
    });
}

function moveRover() {
    const commands = document.getElementById('commands').value;

    fetch('/api/moveSequence', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ commands }),
    })
    .then(response => {
        if (response.ok) {
            console.log('Rover moved successfully!');
            alert('Rover moved successfully!');
        } else {
            return response.json();
        }
    })
    .then(errors => {
        if (errors) {
            console.error('Error moving rover:', errors);
            alert('Error moving rover. Check the console for details.');
        }
    })
    .catch(error => {
        console.error('Error moving rover:', error);
        alert('Error moving rover. Check the console for details.');
    });
}
