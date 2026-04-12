<script lang="ts">
	import ReportCard from './ReportCard.svelte';

	import { reportStore } from '../../stores/reportStore';
</script>

<div class="flex flex-col gap-2 p-2 pb-20 max-h-full">
	{#each $reportStore.reports as report}
		<ReportCard {report} />
	{/each}
</div>
