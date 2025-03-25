let map = L.map("map", {
  zoom: 17,
  center: L.latLng(-33.9173534616886, 151.23126985972831),
  maxBounds: L.latLngBounds(
    L.latLng(-33.86692728515977, 151.17067658807477),
    L.latLng(-33.94499762920698, 151.28045227050784),
  ),
});

L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
  maxZoom: 19,
  minZoom: 14,
  attribution:
    '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
}).addTo(map);

let popup = L.popup();
let latlng = null;

map.on("click", function (e) {
  latlng = e.latlng;
  popup
    .setLatLng(e.latlng)
    .setContent(`Add flare <button @click="popup=true">✅<button>`)
    .openOn(map);
});

// navigator.geolocation.getCurrentPosition(
//   (pos) => {
//     let mark = L.marker([pos.coords.latitude, pos.coords.longitude]);
//     mark.bindTooltip("You are here").openTooltip();
//     mark.addTo(map);
//   },
// );
