// TODO: refactor and test this file
import {CanvasPainter} from './painter.js';

let roverId;
let planet;

document.addEventListener('DOMContentLoaded', () => getRandomRover()); // NOSONAR

document.addEventListener('keydown', async (event) => { // NOSONAR
    switch (event.key) {
        case 'ArrowUp':
        case 'k':
            await moveRover('f');
            break;
        case 'ArrowDown':
        case 'j':
            await moveRover('b');
            break;
        case 'ArrowLeft':
        case 'h':
            await moveRover('l');
            break;
        case 'ArrowRight':
        case 'l':
            await moveRover('r');
            break;
        default:
            break;
    }
});

export async function getRandomRover() {
    const response = await fetch('/api/randomRover', {
        method: 'POST',
    });

    const data = await response.json();
    if (!response.ok) {
        console.error('Error getting random rover:', data);
        displayErrors(data);
    }

    roverId = data.Rover.Id;
    planet = data.Planet;
    const canvas = document.getElementById('canvas');
    const ctx = canvas.getContext('2d');
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    const roverDrawer = new CanvasPainter(canvas, 20);
    roverDrawer.drawPlanetAndRover(data.Planet, data.Rover);
}

export async function moveRoverByCommands() {
    const commands = document.getElementById('commands').value;
    await moveRover(commands);
}

async function moveRover(commands) {
    if (!roverId) {
        displayErrors('Rover ID not available. Call getRandomRover first.');
        return;
    }
    const moveData = await callMoveEndpoint(commands);
    const canvas = document.getElementById('canvas');
    const roverDrawer = new CanvasPainter(canvas, 20);
    roverDrawer.drawPlanetAndRover(planet, moveData.Rover);
    displayErrors(moveData.Errors);
}

async function callMoveEndpoint(commands) {
    const response = await fetch('/api/moveSequence', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({id: roverId, commands}),
    });

    const data = await response.json();
    if (!response.ok) {
        console.error('Error calling endpoint:', data);
        displayErrors(data);
    }

    return data;
}


// TODO: movement errors should be treated differently than bad requests (sending invalid commands)
function displayErrors(errors) {
    const errorBox = document.getElementById('error-box');
    const errorList = document.getElementById('error-list');

    errorList.innerHTML = '';

    if (errors && errors.length > 0) {
        errorBox.classList.remove('hidden');
        errors.forEach(error => {
            const listItem = document.createElement('li');
            listItem.textContent = error;
            errorList.appendChild(listItem);
        });
    } else {
        errorBox.classList.add('hidden');
    }
}

window.random = getRandomRover;
window.move = moveRoverByCommands;
