export class CanvasPainter {
    constructor(canvas) {
        this.canvas = canvas;
        this.ctx = canvas.getContext('2d');
        this.cellSize = 20;
    }

    drawPlanet(planet) {
        this.canvas.width = planet.Width * this.cellSize;
        this.canvas.height = planet.Height * this.cellSize;

        // Draw the planet
        this.ctx.fillStyle = 'white';
        this.ctx.fillRect(0, 0, planet.Width * this.cellSize, planet.Height * this.cellSize);

        // Draw the borders
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

        if (rover.Direction === 'N') {
            this.ctx.moveTo(roverX, roverY);
            this.ctx.lineTo(roverX + this.cellSize, roverY);
            this.ctx.lineTo(roverX + this.cellSize / 2, roverY - this.cellSize);
        } else if (rover.Direction === 'S') {
            this.ctx.moveTo(roverX, roverY);
            this.ctx.lineTo(roverX + this.cellSize, roverY);
            this.ctx.lineTo(roverX + this.cellSize / 2, roverY + this.cellSize);
        } else if (rover.Direction === 'E') {
            this.ctx.moveTo(roverX, roverY);
            this.ctx.lineTo(roverX, roverY + this.cellSize);
            this.ctx.lineTo(roverX + this.cellSize, roverY + this.cellSize / 2);
        } else if (rover.Direction === 'W') {
            this.ctx.moveTo(roverX, roverY);
            this.ctx.lineTo(roverX, roverY + this.cellSize);
            this.ctx.lineTo(roverX - this.cellSize, roverY + this.cellSize / 2);
        }
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
