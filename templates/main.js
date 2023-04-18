const search = document.getElementById('myInput');
const matchList = document.getElementById('match-list');
let youknow = []

// Search data.json and filter it    

const searchNames = async searchText => {
  const res = await fetch('./data.json')
  const data = await res.json();
  console.log(data)
  data.forEach(element => {
    element.creationDate = element.creationDate.toString()
  });
  // Get matches to current text input
  let matches = data.filter((name) => {
    const regex = new RegExp(`${searchText}`, 'gi');
    return name.name.match(regex)
  });


  if (searchText.length === 0) {
    matches = [];
    matchList.innerHTML = '';
  }

  outputHtml(matches);
};

// Show Results in HTML
const outputHtml = matches => {
  if (matches.length >= 0) {

    const html = matches.map((match) => {
      let box = ""
      let boxloc = ""
      let news = ""
      youknow = []
      for (let b = 0; b < match.members.length; b++) {
        box += `<li style="color: white;">${(match.members[b])}</li>`
      }
      for (let c = 0; c < match.locations.length; c++) {
        news = match.locations[c].replaceAll('_', ' ');
        finalstring = news.replaceAll('-', ' ');
        boxloc += `<li style="color: white;">${(finalstring)}</li>`
        youknow.push(finalstring)
      }
      console.log(youknow)
      return `
            <div id="search-list-box" class="result result-body mb-1">
            <div class="search-list-left">
            <h3 style="color: rgb(211, 21, 84);;">'CLEAR' button top-right AVAILABLE.</h3>
            <br>
            <img id="search-list-image" src="${match.image}"/>
            </div>
            <div class="search-list-right">
            <h3>${match.name} ${match.creationDate}  </h4>
            <h5>Premier Album: ${match.firstAlbum}</h5>
            <br>
            <h4 style="color: white;"> MEMBERS:${box} </h3>
            <br>
            <h4 style="color: white;"> LOCATIONS:${boxloc} </h3>
            <div class="map"></div>
            </div>
            </div>
            
        `

    }).join('');

    matchList.innerHTML = html;

    initMap();

  }
}
let maps = [];


function initMap() {
  maps = [];

  console.log("maps initialized");

  var geocoder = null;
  var bounds = null;

  document.querySelectorAll(".map").forEach((element) => {
    let map = new google.maps.Map(element, {
      center: { lat: -34.397, lng: 150.644 },
      zoom: 8
    });
    maps.push(map);

    geocoder = new google.maps.Geocoder();
    /* Déclaration de l'objet de type LatLngBounds qui définira les limites de la carte */
    bounds = new google.maps.LatLngBounds();

    addMarkers();
  });

  /* Fonction chargée de lancer l'application */
  /* Elle appelle la fonction "downloadUrl" pour aller lire
   le fichier fichier-json.json en AJAX*/

  function addMarkers() {
    for (a = 0; a < youknow.length; a++) {
      youknow.forEach(element => {
        //console.log(adress)
        geocodeAddress(element);
      });
  }
};

  /* Fonction de géocodage d'une adresse */
  function geocodeAddress(addr) {
    geocoder.geocode({ 'address': addr }, function (results, status) {
      /* Si la géolocalisation réussit */
      console.log(addr, results, status);
      if (status == google.maps.GeocoderStatus.OK) {
        /* On récupère les coordonnées de l'adresse */
        coords = results[0].geometry.location;
        /* On étend les limites de la carte afin d'y inclure ce nouveau point */
        bounds.extend(coords);
        /* On déclare le marker associé */
        /* On l'ajoute à la carte */
        maps.forEach(map => {
          var marker = new google.maps.Marker({
            position: coords
          });

          marker.setMap(map);
          map.fitBounds(bounds);
        });

        /* On affiche la carte avec un zoom adapté afin de voir tous nos markers */

      }
    });
  }
}

//search.addEventListener('input' , () => searchNames(search.value));

var delay = (function () {
  var timer = 0;
  return function (callback, ms) {
    clearTimeout(timer);
    timer = setTimeout(callback, ms);
  };
})();

let execute = function () {
  searchNames(search.value);
};
delay(function (){
  initMap();
}, 1000);

search.addEventListener('keyup', function () {
  delay(function () {
    execute();
  }, 1000);
});
