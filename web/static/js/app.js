var map = L.map("map").setView([-33.9173534616886, 151.23126985972831], 17);

L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
  maxZoom: 19,
  attribution:
    '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
}).addTo(map);

map.on('click', function(e) {
    console.log(e.latlng.lat);
    console.log(e.latlng.lng);
    popup = document.getElementById("popup");
    popup.classList.toggle("hidden");
})
