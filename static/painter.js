export class CanvasPainter {
    constructor(canvas, cellSize) {
        this.canvas = canvas;
        this.ctx = canvas.getContext('2d');
        this.cellSize = cellSize;
    }

    drawPlanet(planet) {
        this.canvas.width = planet.Width * this.cellSize;
        this.canvas.height = planet.Height * this.cellSize;

        this.ctx.fillStyle = 'white';
        this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height);
        this.drawCellBorders(planet);
    }

    drawCellBorders(planet) {
        this.ctx.strokeStyle = 'lightgrey';
        for (let x = 0; x < planet.Width; x++) {
            for (let y = 0; y < planet.Height; y++) {
                this.ctx.strokeRect(x * this.cellSize, y * this.cellSize, this.cellSize, this.cellSize);
            }
        }
    }

    drawRover(rover) {
        this.ctx.fillStyle = 'red';
        this.ctx.beginPath();
        const roverX = rover.Coordinate.X * this.cellSize;
        const roverY = this.canvas.height - (rover.Coordinate.Y + 1) * this.cellSize;

        // Save the current state of the context
        this.ctx.save();

        // Translate to the center of the rover
        this.ctx.translate(roverX + this.cellSize / 2, roverY + this.cellSize / 2);

        // Rotate the context based on the rover's direction
        if (rover.Direction === 'N') {
            this.ctx.rotate(Math.PI);
        } else if (rover.Direction === 'S') {
            this.ctx.rotate(0);
        } else if (rover.Direction === 'E') {
            this.ctx.rotate(3 * Math.PI / 2);
        } else if (rover.Direction === 'W') {
            this.ctx.rotate(Math.PI / 2);
        }

        // Draw the rover at the origin (0, 0)
        this.ctx.moveTo(-this.cellSize / 2, -this.cellSize / 2);
        this.ctx.lineTo(this.cellSize / 2, -this.cellSize / 2);
        this.ctx.lineTo(0, this.cellSize / 2);
        this.ctx.lineTo(-this.cellSize / 2, -this.cellSize / 2);
        this.ctx.closePath();

        // Restore the context to its original state
        this.ctx.restore();

        this.ctx.fill();
    }

    drawObstacles(planet) {
        this.ctx.fillStyle = 'black';
        planet.Obstacles.forEach(obstacle => {
            obstacle.Coordinate.forEach(coordinate => {
                this.ctx.fillRect(
                    coordinate.X * this.cellSize,
                    this.canvas.height - (coordinate.Y + 1) * this.cellSize,
                    this.cellSize,
                    this.cellSize
                );
            });
        });
    }

    drawPlanetAndRover(planet, rover) {
        this.drawPlanet(planet);
        this.drawObstacles(planet);
        this.drawRover(rover);
    }
}
