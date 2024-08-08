<script lang="ts">
	import { Map } from 'leaflet';
	import 'leaflet/dist/leaflet.css';
	import { onDestroy, onMount } from 'svelte';
	import { reportStore } from '../../stores/reportStore';

	let mapElement: HTMLDivElement;
	let map: Map;

	let MELB: [number, number] = [-37.81, 144.96];

	let markers: Record<string, L.CircleMarker> = {};
	let userLocation: [number, number] | null = null;

	onMount(async () => {
		const L = await import('leaflet');

		map = L.map(mapElement).setView(MELB, 13);

		L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
			attribution:
				'© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
		}).addTo(map);

		for (const report of $reportStore.reports) {
			let marker = L.circleMarker(report.coordinates, {
				color: 'red',
				radius: 4
			});
			markers[report.id.toString()] = marker;

			marker.addTo(map).on('click', () => {
				console.log(report);
				reportStore.setHighlightedReport(report.id);
				const el = document.querySelector(`#report-${report.id}`);
				if (el) {
					el.scrollIntoView({
						behavior: 'smooth'
					});
				}
			});
		}

		navigator.geolocation.getCurrentPosition(
			(pos) => {
				userLocation = [pos.coords.latitude, pos.coords.longitude];
				L.circleMarker(userLocation).addTo(map).bindPopup('You are here!').openPopup();
			},
			() => {},
			{}
		);
	});

	$: if ($reportStore.highlightedReport) {
		const report = $reportStore.reports.find((r) => r.id === $reportStore.highlightedReport);
		if (report) {
			map?.setView(report.coordinates, 13);
			for (const marker of Object.values(markers)) {
				marker.setStyle({
					color: 'red'
				});
			}
			markers[report.id.toString()].setStyle({
				color: 'purple'
			});
		}
	} else {
		map?.setView(userLocation ?? MELB, 13);
	}

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

<style>
	#map {
		height: 100%;
	}
</style>
