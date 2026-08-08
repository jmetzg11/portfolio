(function () {
	const PLACES = JSON.parse(document.getElementById('map-data').textContent);

	const map = L.map('map', { scrollWheelZoom: false });

	// Carto Positron: same OSM data, far less ink than the standard style, so the
	// pins read as the content rather than competing with the basemap. No key.
	L.tileLayer('https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png', {
		attribution:
			'&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors &copy; <a href="https://carto.com/attributions">CARTO</a>',
		subdomains: 'abcd',
		maxZoom: 20,
	}).addTo(map);

	const layer = L.layerGroup().addTo(map);
	const markers = new Map();

	for (const p of PLACES) {
		const marker = L.marker([p.lat, p.lng], { icon: pinFor(p), title: p.name }).bindPopup(popup(p));
		markers.set(p.id, { marker, place: p });
	}

	function pinFor(p) {
		const booking = p.booking ? ` book-${p.booking}` : '';
		return L.divIcon({
			className: 'cdmx-pin-wrap',
			html: `<span class="cdmx-pin${booking}">${p.icon}</span>`,
			iconSize: [36, 36],
			iconAnchor: [18, 18],
			popupAnchor: [0, -19],
		});
	}

	// Built as DOM rather than an HTML string so a stray angle bracket in a
	// place name stays text instead of becoming markup.
	function popup(p) {
		const el = document.createElement('div');
		el.className = 'cdmx-popup';

		const name = document.createElement('strong');
		name.textContent = p.name;
		el.append(name);

		if (p.note) {
			const note = document.createElement('em');
			note.textContent = p.note;
			el.append(note);
		}

		if (p.bookText) {
			const booking = document.createElement('span');
			booking.className = 'badge badge-' + p.booking;
			booking.textContent = p.bookText;
			el.append(booking);
		}

		const address = document.createElement('span');
		address.textContent = p.address;
		el.append(address);

		const links = document.createElement('div');
		links.className = 'links';

		if (p.slug) {
			links.append(link('/cdmx/' + p.slug, 'More details', false));
		}
		links.append(link(p.mapsUrl, 'Google Maps', true));
		el.append(links);

		return el;
	}

	function link(href, text, external) {
		const a = document.createElement('a');
		a.href = href;
		a.textContent = text;
		if (external) {
			a.target = '_blank';
			a.rel = 'noopener noreferrer';
		}
		return a;
	}

	let current = 'all';

	function render(category) {
		current = category;
		layer.clearLayers();

		const shownPoints = [];
		for (const { marker, place } of markers.values()) {
			// Home survives every filter, and counts towards the bounds below, so
			// each category can be read as a distance from the house.
			const shown = category === 'all' || place.category === category || place.category === 'home';
			if (shown) {
				marker.addTo(layer);
				shownPoints.push([place.lat, place.lng]);
			}
		}

		for (const section of document.querySelectorAll('section.category')) {
			const shown = category === 'all' || section.dataset.category === category;
			section.hidden = !shown;
		}

		fit(shownPoints);
	}

	// Frames whatever is currently on the map. maxZoom keeps a single-pin
	// category (or two pins on the same block) from slamming to street level.
	function fit(points) {
		if (!points.length) return;
		map.fitBounds(L.latLngBounds(points), { padding: [40, 40], maxZoom: 15 });
	}

	const picker = document.getElementById('category-picker');

	picker.addEventListener('click', (event) => {
		const button = event.target.closest('button[data-category]');
		if (!button) return;

		for (const b of picker.querySelectorAll('button')) {
			b.classList.toggle('is-selected', b === button);
		}
		render(button.dataset.category);
	});

	// Clicking a card shows it on the map. Links inside the card are left alone
	// so "Directions" still opens Google Maps.
	document.addEventListener('click', (event) => {
		if (event.target.closest('a')) return;

		const card = event.target.closest('li[data-id]');
		if (!card) return;

		const entry = markers.get(Number(card.dataset.id));
		if (!entry) return;

		map.setView([entry.place.lat, entry.place.lng], 16);
		entry.marker.openPopup();
		document.getElementById('map').scrollIntoView({ behavior: 'smooth', block: 'nearest' });
	});

	// A rotated phone or a resized window leaves Leaflet with stale dimensions,
	// so it has to be told to re-measure before the refit means anything.
	window.addEventListener('resize', () => {
		map.invalidateSize();
		render(current);
	});

	render('all');
})();
