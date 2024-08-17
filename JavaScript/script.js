// Get the canvas element by its ID
var canvas = document.getElementById("myCanvas");
const cellSize = 20; // Each cell is 20*20 pixels
const cols = canvas.width / cellSize; 
const rows = canvas.height / cellSize; 
var position;
var row;
var col;

const emptyValue = 0;
const appleValue = 1;
const snakeValue = 2;

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

    position = getRandomInt(rows * cols); 
    row = Math.floor(position / cols);
    col = position % cols;
    console.log(row)
    console.log(col)

    screenArray[row][col] = snakeValue; // Set position for snake (2)

    do{
        position = getRandomInt(rows * cols);
        row = Math.floor(position / cols);
        col = position % cols;
    }while(screenArray[row][col] == snakeValue);
    screenArray[row][col] = appleValue; // Set position for apple (1)


    // Function to draw the snake parts
    function drawSnake() {
        for (let row = 0; row < rows; row++) {
            for (let col = 0; col < cols; col++) {
                if (screenArray[row][col] === appleValue) {
                    // Calculate the position for the circle
                    const x = col * cellSize + cellSize / 2;
                    const y = row * cellSize + cellSize / 2;

                    // Draw the circle (apple)
                    ctx.beginPath();
                    ctx.arc(x, y, radius, 0, 2 * Math.PI);
                    ctx.fillStyle = 'red';
                    ctx.fill();
                    ctx.stroke();
                }
                else if (screenArray[row][col] === snakeValue) {
                    // Calculate the position for the square (snake segment)
                    const x = col * cellSize;
                    const y = row * cellSize;

                    // Draw the square
                    ctx.beginPath();
                    ctx.rect(x, y, cellSize, cellSize);
                    ctx.fillStyle = 'green'; // Snake color
                    ctx.fill();
                    ctx.stroke();
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

    // Start the initial update
    update();
} else {
    console.error("Canvas not found.");
}
