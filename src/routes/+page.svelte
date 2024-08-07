<script lang="ts">
	import LeafletMap from '$lib/components/LeafletMap.svelte';
	import ReportList from '$lib/components/ReportList.svelte';

	import { reportStore } from '../stores/reportStore.js';

	export let data;

	$: reportStore.setReports(data.reports);
</script>

<svelte:head>
	<title>Mates?</title>
</svelte:head>

<main class="flex flex-col lg:flex-row overflow-hidden">
	<div class="flex-1">
		<LeafletMap />
	</div>
	<div class="flex-1 overflow-scroll h-full">
		<ReportList />
	</div>
</main>
