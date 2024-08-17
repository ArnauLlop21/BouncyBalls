var canvas = document.getElementById("myCanvas");

// Check if canva exists
if (canvas) {
    // Obtain the 2d container to draw
    var ctx = canvas.getContext("2d");

    // Define the circle coordinates
    var centerX = canvas.width / 2;
    var centerY = canvas.height / 2;

    // Define the circle radius
    var radius = 50;

    // Begin a new path to draw
    ctx.beginPath();

    // Draw the circle
    ctx.arc(centerX, centerY, radius, 0, 2 * Math.PI);

    // Fill the circle with a certain color.
    ctx.fillStyle = "blue";
    ctx.fill();
} else {
    console.error("El canvas no s'ha trobat.");
}
