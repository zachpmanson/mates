<script lang="ts">
	import { timeAgo } from '$lib/util/timesince';
	import Card from './ui/card/card.svelte';

	import { reportStore } from '../../stores/reportStore';
	import type { MateReport } from '../../types/report';
	import tram from '../images/tram.svg';
	import train from '../images/train.svg';

	export let report: MateReport;
</script>

<Card
	class={$reportStore.highlightedReport === report.id ? 'bg-gray-100' : ''}
	on:click={() => {
		if ($reportStore.highlightedReport !== report.id) reportStore.setHighlightedReport(report.id);
		else reportStore.setHighlightedReport(null);
	}}
>
	<div class="flex flex-col py-3 px-4">
		<div class="flex items-start justify-between gap-2">
			<span class="">{report.title}</span>
			{#if report.type === 'tram'}
				<img src={tram} alt="Tram icon" class="h-6 w-6" />
				<!-- {:else if report.type === 'bus'}
				<img src={bus} alt="Bus icon" class="h-6 w-6" /> -->
			{:else if report.type === 'train'}
				<img src={train} alt="Train icon" class="h-6 w-6" />
			{/if}
		</div>
		<div class="flex items-end justify-between gap-2">
			<span class="text-gray-400">
				{timeAgo(new Date(report.timestamp))} ago
			</span>

			<span class="text-gray-400 text-xs">
				Spotted by {report.author}
			</span>
		</div>
	</div>
</Card>
