<script lang="ts">
	import type { MateReport } from '../../types/report';
	import ReportCard from './ReportCard.svelte';

	import { reportStore } from '../../stores/reportStore';
</script>

<div class="h-full flex flex-col gap-2 p-2">
	{#each $reportStore.reports as report}
		<ReportCard {report} />
	{/each}
</div>
