// Draws the map for /geography. All the grouping, sizing and colouring happens
// in Go; MAP_DATA arrives keyed by "all" and by each year, so switching years is
// just picking a different slice and redrawing.
(function () {
	const MAP_DATA = JSON.parse(document.getElementById('map-data').textContent);

	const worldBounds = L.latLngBounds(L.latLng(-90, -180), L.latLng(90, 180));

	const map = L.map('map', {
		center: [0, 0],
		zoom: 3,
		minZoom: 2,
		maxBounds: worldBounds,
		maxBoundsViscosity: 0.8,
	});

	// Tiles wrap on purpose: pinning them to a single world leaves dead space
	// either side of any container wider than the map.
	L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
		attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
		maxZoom: 18,
	}).addTo(map);

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
