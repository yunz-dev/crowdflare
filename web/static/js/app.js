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

let addFlarePopup = L.popup();
let latlng = null;

map.on("click", (e) => {
  latlng = e.latlng;
  addFlarePopup
    .setLatLng(e.latlng)
    .setContent(`<button @click="add=true">Fire a flare here 🎆<button>`)
    .openOn(map);
});

// dummy_flares = [
//   { id: "1", lat: -33.91717944990013, lng: 151.2298107147217 },
//   { id: "2", lat: -33.91724622440734, lng: 151.23012721538547 },
//   { id: "3", lat: -33.917406483011156, lng: 151.2306743860245 },
//   { id: "4", lat: -33.91726403093376, lng: 151.22993677854538 },
//   { id: "5", lat: -33.91751371612105, lng: 151.23352171604117 },
// ];
fetch("/api/flares")
  .then((response) => response.json())
  .then((json) => {
    if (json == null) {
      return;
    }
    for (const flare of json) {
      let marker = L.marker([flare.Lat, flare.Lng], {
        bubblingMouseEvents: false,
        fillColor: "blue",
        stroke: false,
        fillOpacity: 0.4,
        radius: 10,
      });
      marker.addTo(map);
      popup = L.popup(null, { maxWidth: 800 });
      popup.setContent(
        `
        <div class="flare-popup" x-data="{upvotes: ${flare.Upvotes}, downvotes: ${flare.Downvotes}}">
            <h1 class="text-lg">${flare.OwnerId}</hi>
            <h2>${flare.Rating}/5</h2>
            <div class="flex flex-row items-center justify-around">
                <button
                    hx-put="/api/flare/upvote"
                    hx-vals='{"id": "${flare.ID}"}'
                    hx-trigger="click once"
                    hx-swap="none"
                    @click.once="upvotes++">⬆️</button>
                <h3 x-text="upvotes"></h3>
                <button
                    hx-put="/api/flare/downvote"
                    hx-vals='{"id": "${flare.ID}"}'
                    hx-trigger="click once"
                    hx-swap="none"
                    @click.once="downvotes++">⬇️</button>
                <h3 x-text="downvotes"></h3>
            </div>
        </div>
        `,
      );
      marker.bindPopup(popup);
      marker.on("click", () => {
        for (const e of document.querySelectorAll(".flare-popup")) {
          htmx.process(e);
        }
      });

      // e = marker._icon;
      // if (!e) {
      //   continue;
      // }

      // e.setAttribute("hx-get", "/component/flare");
      // e.setAttribute("hx-vals", `{"id": ${flare.id}}`);
      // e.setAttribute("hx-target", "#popup");
      // e.setAttribute("x-on:click", "popup=true");
      // htmx.process(e);
    }
  });

// navigator.geolocation.getCurrentPosition(
//   (pos) => {
//     let mark = L.marker([pos.coords.latitude, pos.coords.longitude]);
//     mark.bindTooltip("You are here").openTooltip();
//     mark.addTo(map);
//   },
// );
