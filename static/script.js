// TODO: refactor and test this file
import {CanvasPainter} from './painter.js';

let roverId;
let planet;

document.addEventListener('DOMContentLoaded', () => getRandomRover());

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

export async function moveRover() {
    if (!roverId) {
        displayErrors('Rover ID not available. Call getRandomRover first.');
        return;
    }
    const commands = document.getElementById('commands').value;
    const response = await fetch('/api/moveSequence', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({id: roverId, commands}),
    });

    const data = await response.json();
    if (!response.ok) {
        console.error('Error moving rover:', data);
        displayErrors(data);
    }

    console.log('Rover moved successfully:', data);
    const canvas = document.getElementById('canvas');
    const roverDrawer = new CanvasPainter(canvas, 20);
    roverDrawer.drawPlanetAndRover(planet, data.Rover);
    displayErrors(data.Errors);
}


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
window.move = moveRover;
