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

<main class="h-full flex flex-col lg:flex-row overflow-hidden">
	<div class="flex-[2] lg:flex-1 lg:h-full">
		<LeafletMap />
	</div>
	<div class="flex-[3] lg:flex-1 overflow-scroll lg:h-full">
		<ReportList />
	</div>
</main>
