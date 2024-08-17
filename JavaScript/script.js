var canvas = document.getElementById("myCanvas");

// Check if canva exists
if (canvas) {
    // Get the 2D drawing context
    var ctx = canvas.getContext("2d");

    // Define the center coordinates of the circle
    var centerX = canvas.width / 2;
    var centerY = canvas.height / 2;
    var speedX = 6;
    var speedY = 2;
    const coefficientOfRestitution = 1

    // Define the radius of the circle
    var radius = 50;

    // Function to update and draw the content of the canvas
    function update() {
        // Clear the canvas
        ctx.clearRect(0, 0, canvas.width, canvas.height);

        // Update the circle's coordinates (e.g., moving it in a circular path)
        centerX += speedX;
        centerY += speedY;
        
        if (centerX + radius >= canvas.width){
            centerX = canvas.width - radius;
            speedX *= -coefficientOfRestitution;
        }else if(centerX - radius <= 0){
            centerX = radius;
            speedX *= -coefficientOfRestitution;
        }

        if (centerY + radius >= canvas.height){
            centerY = canvas.height - radius;
            speedY *= -coefficientOfRestitution;
        }else if(centerY - radius <= 0){
            centerY = radius;
            speedY *= -coefficientOfRestitution;
        }
        // Draw the circle
        ctx.beginPath();
        ctx.arc(centerX, centerY, radius, 0, 2 * Math.PI);
        ctx.fillStyle = "purple";
        ctx.fill();


        // Request the browser to call the update function on the next frame
        requestAnimationFrame(update);
    }

    // Start the animation
    update();
} else {
    console.error("Canvas not found.");
}
