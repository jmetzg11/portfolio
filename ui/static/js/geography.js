// Draws the map for /geography. All the grouping, sizing and colouring happens
// in Go; MAP_DATA arrives keyed by "all" and by each year, so switching years is
// just picking a different slice and redrawing.
(function () {
	const MAP_DATA = JSON.parse(document.getElementById('map-data').textContent);

	const worldBounds = L.latLngBounds(L.latLng(-90, -180), L.latLng(90, 180));

	const dotBounds = L.latLngBounds((MAP_DATA.all || []).map((m) => [m.lat, m.lng]));

	const map = L.map('map', {
		center: [0, 0],
		zoom: 3,
		maxBounds: worldBounds,
		maxBoundsViscosity: 0.8,
	});

	// The dots span most of the globe, so a fixed zoom floor that keeps a desktop
	// full of map leaves a phone unable to see them all. Float the floor to
	// whatever fits them here, but never let it out past the old fixed 2.
	// getBoundsZoom clamps to the current minZoom, hence the reset first.
	function fitMinZoom() {
		if (!dotBounds.isValid()) return;
		map.setMinZoom(0);
		map.setMinZoom(Math.min(2, map.getBoundsZoom(dotBounds, false, [40, 40])));
	}

	fitMinZoom();
	map.on('resize', fitMinZoom);

	// Tiles wrap on purpose: pinning them to a single world leaves dead space
	// either side of any container wider than the map.
	// Esri's grey canvas over stock OSM: it stays flat grey at every zoom, where
	// OSM turns green and busy once you're past country level. It's HTTP/1.1 off
	// one host, so the browser only gets ~6 tiles in flight — worth keeping the
	// tile count down. Tiles stop at 16, so deeper zooms upscale.
	L.tileLayer(
		'https://server.arcgisonline.com/ArcGIS/rest/services/Canvas/World_Light_Gray_Base/MapServer/tile/{z}/{y}/{x}',
		{
			attribution: 'Tiles &copy; <a href="https://www.esri.com/">Esri</a>',
			maxZoom: 18,
			maxNativeZoom: 16,
			updateWhenZooming: false,
		},
	).addTo(map);

	// One group holding every marker, so switching years is a single clear.
	const layer = L.layerGroup().addTo(map);

	function render(key) {
		layer.clearLayers();

		for (const m of MAP_DATA[key] || []) {
			L.circleMarker([m.lat, m.lng], {
				radius: m.radius,
				color: m.color,
				fillColor: m.color,
				fillOpacity: 0.5,
			})
				.bindTooltip(tooltip(m), { direction: 'top' })
				.addTo(layer);
		}
	}

	// Built as DOM rather than an HTML string so a stray angle bracket in a
	// place name stays text instead of becoming markup.
	function tooltip(m) {
		const el = document.createElement('div');
		el.className = 'map-tooltip';

		const name = document.createElement('strong');
		name.textContent = m.location;

		const note = document.createElement('em');
		note.textContent = m.text;

		const years = document.createElement('span');
		years.textContent = 'Years: ' + m.years;

		el.append(name, note, years);
		return el;
	}

	const picker = document.getElementById('year-picker');

	picker.addEventListener('click', (event) => {
		const button = event.target.closest('button[data-year]');
		if (!button) return;

		for (const b of picker.querySelectorAll('button')) {
			b.classList.toggle('is-selected', b === button);
		}
		render(button.dataset.year);
	});

	render('all');
})();
