class Snake {
    constructor(id){
        this.id = id;
    }

    getId() {
        return this.id;
    }
}

// Get the canvas element by its ID
var canvas = document.getElementById("myCanvas");
const cellSize = 50; // Each cell is 50*50 pixels
const cols = canvas.width / cellSize; 
const rows = canvas.height / cellSize; 
var position;
var row;
var col;

const emptyValue = 0;
const appleValue = 1;

const leadValue = 0;

// Function to generate a random integer
function getRandomInt(max) {
    return Math.floor(Math.random() * max);
}

// Check if canvas exists
if (canvas) {
    // Get the 2D drawing context
    var ctx = canvas.getContext("2d");
    let screenArray = Array.from({ length: rows }, () => Array(cols).fill(emptyValue));
    
    // Define the radius of the circle
    var radius = cellSize / 2; // Radius is half of cell size

    // Generate random positions within the array for apple and snake
    do {
        position = getRandomInt(rows * cols); 
        row = Math.floor(position / cols);
        col = position % cols;
        console.log("Row for snake head is " + row);
        console.log("Col for snake head is " + col);
        if (row === 0 || row === rows - 1 || col === 0 || col === cols - 1) {
            console.log("Trying again");
        }
    } while (row === 0 || row === rows - 1 || col === 0 || col === cols - 1);

    screenArray[row][col] = new Snake(0); // Set position for snake head (id = 0)

    let newSegment = findNewSnakeSegment(screenArray);
    if (newSegment) {
        screenArray[newSegment.row][newSegment.col] = new Snake(newSegment.newId);
    }

    do {
        position = getRandomInt(rows * cols);
        row = Math.floor(position / cols);
        col = position % cols;
    } while (screenArray[row][col] instanceof Snake);
    screenArray[row][col] = appleValue; // Set position for apple (1)

    // Function to draw the snake parts
    function drawSnake() {
        for (let row = 0; row < rows; row++) {
            for (let col = 0; col < cols; col++) {
                if (screenArray[row][col] === appleValue) {
                    // Calculate the position for the circle (apple)
                    const x = col * cellSize + cellSize / 2;
                    const y = row * cellSize + cellSize / 2;

                    // Draw the circle (apple)
                    ctx.beginPath();
                    ctx.arc(x, y, radius, 0, 2 * Math.PI);
                    ctx.fillStyle = 'red';
                    ctx.fill();
                    ctx.stroke();
                } else if (screenArray[row][col] instanceof Snake) {
                    // Calculate the position for the square (snake segment)
                    const x = col * cellSize;
                    const y = row * cellSize;

                    // Draw the square
                    ctx.beginPath();
                    ctx.rect(x, y, cellSize, cellSize);
                    ctx.fillStyle = 'green'; // Snake color
                    ctx.fill();
                    ctx.stroke();

                    if (screenArray[row][col].getId() === leadValue) {
                        const x2 = col * cellSize + cellSize / 2;
                        const y2 = row * cellSize + cellSize / 2;
                        ctx.beginPath();
                        ctx.arc(x2, y2, cellSize / 5, 0, 2 * Math.PI);
                        ctx.fillStyle = "black";
                        ctx.fill();
                        ctx.stroke();
                    }
                }
            }
        }
    }

    function update() {
        // Clear the canvas
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        
        // Draw the apple and snake
        drawSnake();

        // Request the browser to call the update function on the next frame
        // Uncomment the following line if you want continuous updates
        // requestAnimationFrame(update); 
    }

    function findNewSnakeSegment(screenArray) {
        let highestId = -1;
        let highestRow = -1;
        let highestCol = -1;
    
        // Iterate to find the snake segment with the highest ID
        for (let row = 0; row < rows; row++) {
            for (let col = 0; col < cols; col++) {
                if (screenArray[row][col] instanceof Snake) {
                    if (screenArray[row][col].getId() > highestId) {
                        highestId = screenArray[row][col].getId();
                        highestRow = row;
                        highestCol = col;
                    }
                }
            }
        }
    
        // Possible directions to add a new segment (up, right, down, left)
        const directions = [
            { row: 0, col: 1 } ,  // Right
            { row: -1, col: 0 }, // Up
            { row: 1, col: 0 },  // Down
            { row: 0, col: -1 }, // Left
        ];
    
        let newRow, newCol;
    
        // Iterate over the directions to find a free position
        for (let i = 0; i < directions.length; i++) {
            newRow = highestRow + directions[i].row;
            newCol = highestCol + directions[i].col;
    
            // Check if the new position is within bounds and is free
            if (
                newRow >= 0 && newRow < rows &&
                newCol >= 0 && newCol < cols &&
                screenArray[newRow][newCol] === 0 // Check if it's free (0)
            ) {
                return {
                    row: newRow,
                    col: newCol,
                    newId: highestId + 1
                };
            }
        }
    
        // If there are no free positions (rare situation), return null or some other logic
        return null; // Or return a default value if the snake is completely enclosed
    }

    // Start the initial update
    update();
} else {
    console.error("Canvas not found.");
}
