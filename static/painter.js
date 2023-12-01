export class CanvasPainter {
    constructor(canvas, cellSize) {
        this.canvas = canvas;
        this.ctx = canvas.getContext('2d');
        this.cellSize = cellSize;
    }

    drawPlanet(planet) {
        this.#setCanvasSize(planet);
        this.#drawEmptyPlanet();
        this.#drawCellBorders(planet);
    }

    #drawEmptyPlanet() {
        this.ctx.fillStyle = 'white';
        this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height);
    }

    #setCanvasSize(planet) {
        this.canvas.width = planet.Width * this.cellSize;
        this.canvas.height = planet.Height * this.cellSize;
    }

    #drawCellBorders(planet) {
        this.ctx.strokeStyle = 'lightgrey';
        for (let x = 0; x < planet.Width; x++) {
            for (let y = 0; y < planet.Height; y++) {
                this.ctx.strokeRect(x * this.cellSize, y * this.cellSize, this.cellSize, this.cellSize);
            }
        }
    }

    drawRover(rover) {
        this.ctx.save();
        this.ctx.beginPath();

        this.ctx.fillStyle = 'red';
        this.#centerOnCell(rover);
        this.#pointDirection(rover.Direction);
        this.#drawTriangle();

        this.ctx.restore();
        this.ctx.fill();
    }

    #centerOnCell(rover) {
        const roverXGridPosition = rover.Coordinate.X * this.cellSize;
        const roverYGridPosition = this.canvas.height - (rover.Coordinate.Y + 1) * this.cellSize;
        this.ctx.translate(roverXGridPosition + this.cellSize / 2, roverYGridPosition + this.cellSize / 2);
    }

    #pointDirection(direction) {
        if (direction === 'N') {
            this.ctx.rotate(Math.PI);
        } else if (direction === 'S') {
            this.ctx.rotate(0);
        } else if (direction === 'E') {
            this.ctx.rotate(3 * Math.PI / 2);
        } else if (direction === 'W') {
            this.ctx.rotate(Math.PI / 2);
        }
    }

    #drawTriangle() {
        this.ctx.moveTo(-this.cellSize / 2, -this.cellSize / 2);
        this.ctx.lineTo(this.cellSize / 2, -this.cellSize / 2);
        this.ctx.lineTo(0, this.cellSize / 2);
        this.ctx.lineTo(-this.cellSize / 2, -this.cellSize / 2);
        this.ctx.closePath();
    }

    drawObstacles(planet) {
        this.ctx.fillStyle = 'black';
        planet.Obstacles.forEach(obstacle => {
            obstacle.Coordinate.forEach(coordinate => {
                this.#drawObstacle(coordinate);
            });
        });
    }

    #drawObstacle(coordinate) {
        this.ctx.fillRect(
            coordinate.X * this.cellSize,
            this.canvas.height - (coordinate.Y + 1) * this.cellSize,
            this.cellSize,
            this.cellSize
        );
    }

    drawPlanetAndRover(planet, rover) {
        this.drawPlanet(planet);
        this.drawObstacles(planet);
        this.drawRover(rover);
    }
}
