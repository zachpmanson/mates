<script>
	import LeafletMap from '$lib/components/LeafletMap.svelte';
	import ReportList from '$lib/components/ReportList.svelte';
</script>

<main class="max-h-screen flex sm:flex-col md:flex-row">
	<div class="flex-1">
		<LeafletMap />
	</div>
	<div class="flex-1">
		<ReportList />
	</div>
</main>

<style>
</style>
