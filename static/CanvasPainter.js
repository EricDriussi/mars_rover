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
        obstacles.forEach(obstacle => {
            const isSmall = obstacle.length === 1;
            if (isSmall) {
                this.#ctx.fillStyle = this.#randomColorBlackGreenOrBlue()
            } else {
                this.#ctx.fillStyle = 'black'
            }
            obstacle.forEach(coordinate => {
                this.#drawObstacle(coordinate);
            });
        });
    }

    #randomColorBlackGreenOrBlue() {
        const random = Math.random();
        if (random < 0.1) {
            return 'green';
        } else if (random < 0.2) {
            return 'blue';
        } else if (random < 0.3) {
            return 'yellow';
        } else if (random < 0.4) {
            return 'orange';
        } else if (random < 0.5) {
            return 'purple';
        } else if (random < 0.6) {
            return 'pink';
        } else if (random < 0.7) {
            return 'brown';
        } else {
            return 'gray';
        }
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
                this.#clearIfRedOrWhite(x, y);
            }
        }
    }

    #clearIfRedOrWhite(x, y) {
        const pixelX = x * this.#cellSize + this.#cellSize / 2;
        const pixelY = y * this.#cellSize + this.#cellSize / 2;
        const pixelData = this.#ctx.getImageData(pixelX, pixelY, 1, 1).data;

        if (this.#isRedOrWhite(pixelData)) {
            this.#drawEmptyCell(x * this.#cellSize, y * this.#cellSize);
            this.#drawCellBorder(x * this.#cellSize, y * this.#cellSize);
        }
    }

    #isRedOrWhite(pixelData) {
        const pixelIsRed = pixelData[0] === 255 && pixelData[1] === 0 && pixelData[2] === 0;
        const pixelIsWhite = pixelData[0] === 255 && pixelData[1] === 255 && pixelData[2] === 255;
        return pixelIsRed || pixelIsWhite;
    }
}
