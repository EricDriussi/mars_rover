let roverId;
let planet;

document.addEventListener('DOMContentLoaded', function () {
    getRandomRover();
});

function getRandomRover() {
    fetch('/api/randomRover', {
        method: 'POST',
    })
        .then(response => response.json())
        .then(data => {
            console.log('Random Rover:', data);
            roverId = data.Rover.Id;
            planet = data.Planet;
            const canvas = document.getElementById('canvas');
            const ctx = canvas.getContext('2d');
            ctx.clearRect(0, 0, canvas.width, canvas.height);
            drawPlanetAndRover(data.Planet, data.Rover);
        })
        .catch(error => {
            console.error('Error getting random rover:', error);
            alert('Error getting random rover. Check the console for details.');
        });
}

function moveRover() {
    const commands = document.getElementById('commands').value;
    if (!roverId) {
        alert('Rover ID not available. Call getRandomRover first.');
        return;
    }

    fetch('/api/moveSequence', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({id: roverId, commands}),
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            } else {
                return response.json().then(errors => {
                    throw new Error(errors.join('\n'));
                });
            }
        })
        .then(data => {
            console.log('Rover moved successfully:', data);
            drawPlanetAndRover(planet, data)
        })
        .catch(error => {
            console.error('Error moving rover:', error);
            alert('Error moving rover. Check the console for details.');
        });
}

function drawPlanetAndRover(planet, rover) {
    const canvas = document.getElementById('canvas');

    if (canvas.getContext) {
        const ctx = canvas.getContext('2d');
        const cellSize = 20;
        canvas.width = planet.Width * cellSize;
        canvas.height = planet.Height * cellSize;

        // Draw the planet
        ctx.fillStyle = 'white';
        ctx.fillRect(0, 0, planet.Width * cellSize, planet.Height * cellSize);

        // Draw the borders
        ctx.strokeStyle = 'lightgrey';
        for (let x = 0; x < planet.Width; x++) {
            for (let y = 0; y < planet.Height; y++) {
                ctx.strokeRect(x * cellSize, y * cellSize, cellSize, cellSize);
            }
        }

        // Draw the rover
        ctx.fillStyle = 'red';
        // TODO: sometimes it's not there
        ctx.beginPath();
        const roverX = rover.Coordinate.X * cellSize;
        const roverY = canvas.height - (rover.Coordinate.Y + 1) * cellSize;

        if (rover.Direction === 'N') {
            ctx.moveTo(roverX, roverY);
            ctx.lineTo(roverX + cellSize, roverY);
            ctx.lineTo(roverX + cellSize / 2, roverY - cellSize);
        } else if (rover.Direction === 'S') {
            ctx.moveTo(roverX, roverY);
            ctx.lineTo(roverX + cellSize, roverY);
            ctx.lineTo(roverX + cellSize / 2, roverY + cellSize);
        } else if (rover.Direction === 'E') {
            ctx.moveTo(roverX, roverY);
            ctx.lineTo(roverX, roverY + cellSize);
            ctx.lineTo(roverX + cellSize, roverY + cellSize / 2);
        } else if (rover.Direction === 'W') {
            ctx.moveTo(roverX, roverY);
            ctx.lineTo(roverX, roverY + cellSize);
            ctx.lineTo(roverX - cellSize, roverY + cellSize / 2);
        }
        ctx.fill();

        // Draw the obstacles
        ctx.fillStyle = 'black';
        // TODO: why is this sometimes null?
        planet.Obstacles.forEach(obstacle => {
            obstacle.Coordinate.forEach(coordinate => {
                ctx.fillRect(coordinate.X * cellSize, canvas.height - (coordinate.Y + 1) * cellSize, cellSize, cellSize);
            });
        });
    }
}
