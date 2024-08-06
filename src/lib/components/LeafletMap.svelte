<script lang="ts">
	import 'leaflet/dist/leaflet.css'; //Don't forget to declare leaflet css
	import { onMount, onDestroy } from 'svelte';
	import type { Map } from 'leaflet';

	let mapElement: HTMLDivElement;
	let map: Map;

	let MELB: [number, number] = [-37.81, 144.96];

	onMount(async () => {
		const L = await import('leaflet');

		map = L.map(mapElement).setView(MELB, 13);

		L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
			attribution:
				'© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
		}).addTo(map);

		L.marker(MELB)
			.addTo(map)
			.bindPopup('A pretty CSS3 popup.<br> Easily customizable.')
			.openPopup();

		let circle = L.circle(MELB, {
			color: 'red',
			fillColor: '#f03',
			fillOpacity: 0.5,
			radius: 500
		}).addTo(map);

		navigator.geolocation.getCurrentPosition(
			(pos) => {
				L.marker([pos.coords.latitude, pos.coords.longitude])
					.addTo(map)
					.bindPopup('A pretty CSS3 popup.<br> Easily customizable.')
					.openPopup();
			},
			() => {},
			{}
		);
	});

	onDestroy(async () => {
		if (map) {
			console.log('Unloading Leaflet map.');
			map.remove();
		}
	});
</script>

<div class="h-full w-full">
	<div bind:this={mapElement} id="map"></div>
</div>

<!-- <link
	rel="stylesheet"
	href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
	integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
	crossorigin=""
/> -->

<style>
	#map {
		height: 50vh;
	}
</style>
