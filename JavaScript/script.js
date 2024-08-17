// Get the canvas element by its ID
var canvas = document.getElementById("myCanvas");
const rows = 150;
const cols = 120;
const cellSize = 5; // Fix the variable name to be consistent

// Function to generate a random integer
function getRandomInt(max) {
    return Math.floor(Math.random() * max);
}

// Check if canvas exists
if (canvas) {
    // Get the 2D drawing context
    var ctx = canvas.getContext("2d");
    let screenArray = Array.from({ length: rows }, () => Array(cols).fill(0));
    
    // Define the radius of the circle
    var radius = cellSize / 2; // Updated to be consistent with cellSize

    // Generate a random position within the array
    var position = getRandomInt(rows * cols); // Use rows * cols to get a valid index
    const row = Math.floor(position / cols);
    const col = position % cols;

    screenArray[row][col] = 1;

    // Function to draw the circles (apples)
    function drawApples() {
        for (let row = 0; row < rows; row++) {
            for (let col = 0; col < cols; col++) {
                if (screenArray[row][col] === 1) {
                    // Calculate the position for the circle
                    const x = col * cellSize + cellSize / 2;
                    const y = row * cellSize + cellSize / 2;

                    // Draw the circle
                    ctx.beginPath();
                    ctx.arc(x, y, radius, 0, 2 * Math.PI);
                    ctx.fillStyle = 'blue'; // Circle color
                    ctx.fill();
                    ctx.stroke(); // Optional: draw the border of the circle
                }
            }
        }
    }

    function update() {
        // Clear the canvas
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        
        // Draw the apple circles
        drawApples();

        // Request the browser to call the update function on the next frame
        requestAnimationFrame(update); // Uncomment this line if you want continuous updates
    }

    // Start the animation
    update();
} else {
    console.error("Canvas not found.");
}
