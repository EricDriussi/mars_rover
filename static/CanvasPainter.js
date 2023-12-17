export class CanvasPainter {
    #canvas;
    #ctx;
    #cellSize;
    #borderSize;

    constructor(canvas, cellSize, borderSize) {
        this.#canvas = canvas;
        this.#ctx = canvas.getContext('2d');
        this.#cellSize = cellSize;
        this.#borderSize = borderSize
    }

    drawPlanet(planet) {
        this.#setCanvasSize(planet.Width, planet.Height);
        this.#drawEmptyPlanet();
        this.#drawCellsBorders(planet);
    }

    #drawEmptyPlanet() {
        this.#ctx.fillStyle = 'white';
        this.#ctx.fillRect(0, 0, this.#canvas.width, this.#canvas.height);
    }

    #setCanvasSize(width, height) {
        this.#canvas.width = width * this.#cellSize;
        this.#canvas.height = height * this.#cellSize;
    }

    #drawCellsBorders(planet) {
        for (let x = 0; x < planet.Width; x++) {
            for (let y = 0; y < planet.Height; y++) {
                this.#drawCellBorder(x * this.#cellSize, y * this.#cellSize);
            }
        }
    }

    #drawCellBorder(x, y) {
        this.#ctx.strokeStyle = 'lightgrey';
        this.#ctx.strokeRect(x, y, this.#cellSize, this.#cellSize);
    }

    drawRover(direction, coordinate) {
        this.#clearCellsWithoutObstacles();
        this.#ctx.fillStyle = 'red';
        this.#ctx.save();
        this.#ctx.beginPath();

        this.#centerOnCell(coordinate);
        this.#pointDirection(direction);
        this.#drawTriangle();

        this.#ctx.restore();
        this.#ctx.fill();
    }

    #pointDirection(direction) {
        if (direction === 'N') {
            this.#ctx.rotate(Math.PI);
        } else if (direction === 'S') {
            this.#ctx.rotate(0);
        } else if (direction === 'E') {
            this.#ctx.rotate(3 * Math.PI / 2);
        } else if (direction === 'W') {
            this.#ctx.rotate(Math.PI / 2);
        }
    }

    #drawTriangle() {
        this.#ctx.moveTo(-this.#cellSize / 2, -this.#cellSize / 2);
        this.#ctx.lineTo(this.#cellSize / 2, -this.#cellSize / 2);
        this.#ctx.lineTo(0, this.#cellSize / 2);
        this.#ctx.lineTo(-this.#cellSize / 2, -this.#cellSize / 2);
        this.#ctx.closePath();
    }

    #drawEmptyCell(roverXGridPosition, roverYGridPosition) {
        this.#ctx.fillStyle = 'white';
        this.#ctx.fillRect(roverXGridPosition, roverYGridPosition, this.#cellSize, this.#cellSize);
    }

    #centerOnCell(coordinate) {
        const roverXGridPosition = coordinate.X * this.#cellSize;
        const roverYGridPosition = this.#canvas.height - (coordinate.Y + 1) * this.#cellSize;
        this.#ctx.translate(roverXGridPosition + this.#cellSize / 2, roverYGridPosition + this.#cellSize / 2);
    }

    drawObstacles(obstacles) {
        this.#ctx.fillStyle = 'black';
        obstacles.forEach(obstacle => {
            obstacle.Coordinate.forEach(coordinate => {
                this.#drawObstacle(coordinate);
            });
        });
    }

    #drawObstacle(coordinate) {
        this.#ctx.fillRect(
            coordinate.X * this.#cellSize + this.#borderSize,
            this.#canvas.height - (coordinate.Y + 1) * this.#cellSize + this.#borderSize,
            this.#cellSize - 2 * this.#borderSize,
            this.#cellSize - 2 * this.#borderSize
        );
    }

    #clearCellsWithoutObstacles() {
        const widthInCells = this.#canvas.width / this.#cellSize;
        const heightInCells = this.#canvas.height / this.#cellSize;

        for (let x = 0; x < widthInCells; x++) {
            for (let y = 0; y < heightInCells; y++) {
                this.#clearIfNotBlack(x, y);
            }
        }
    }

    #clearIfNotBlack(x, y) {
        // Get the pixel data from the middle of the cell to avoid issues with borders
        const pixelX = x * this.#cellSize + this.#cellSize / 2;
        const pixelY = y * this.#cellSize + this.#cellSize / 2;
        const pixelData = this.#ctx.getImageData(pixelX, pixelY, 1, 1).data;

        const pixelIsNotBlack = pixelData[0] !== 0 || pixelData[1] !== 0 || pixelData[2] !== 0; // RGB
        if (pixelIsNotBlack) {
            this.#drawEmptyCell(x * this.#cellSize, y * this.#cellSize);
            this.#drawCellBorder(x * this.#cellSize, y * this.#cellSize);
        }
    }
}
