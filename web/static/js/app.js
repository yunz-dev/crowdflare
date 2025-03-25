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
    .setContent(`<button @click="popup=true">Add flare 🎆<button>`)
    .openOn(map);
});

dummy_flares = [
  { id: "1", lat: -33.91717944990013, lng: 151.2298107147217 },
  { id: "2", lat: -33.91724622440734, lng: 151.23012721538547 },
  { id: "3", lat: -33.917406483011156, lng: 151.2306743860245 },
];

for (const flare of dummy_flares) {
  let marker = L.circle([flare.lat, flare.lng]);
  marker.setStyle({
    fillColor: "red",
    stroke: false,
    fillOpacity: 0.4,
    radius: 200,
  });
  marker.addTo(map);

  e = marker._path;
  if (!e) {
    continue;
  }

  e.setAttribute("hx-get", "/component/flare");
  e.setAttribute("hx-vals", `{"id": ${flare.id}}`);
  e.setAttribute("hx-target", "#popup");
  e.setAttribute("x-on:click", "popup=true");
  htmx.process(e);
}

// navigator.geolocation.getCurrentPosition(
//   (pos) => {
//     let mark = L.marker([pos.coords.latitude, pos.coords.longitude]);
//     mark.bindTooltip("You are here").openTooltip();
//     mark.addTo(map);
//   },
// );
