console.log("hello world");

document.addEventListener("DOMContentLoaded", function() {
    console.log(getPhotos());
})

async function getPhotos() {
    const url = "/api/photos";
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`Response: ${response.status}`);
        }
        const result = await response.json();
        console.log(result);
        for (const photo of result) {
            const gallery = document.querySelector(".grid");
            console.log(gallery)
            const img = document.createElement("img");
            img.src = photo.url;
            img.alt = photo.name;
            gallery.appendChild(img);
        }
    } catch (error) {
        console.error(error.message)
    }
}

console.log("test");
