<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import LeafletMap from '$lib/components/LeafletMap.svelte';
	import ReportList from '$lib/components/ReportList.svelte';
	import { reportStore } from '../stores/reportStore.js';
	import type { Feed } from './+page.js';

	export let data: { feeds: Feed[]; reports: import('../types/report').MateReport[] };

	$: reportStore.setReports(data.reports);

	$: currentFeedId = $page.url.searchParams.get('feed_id') ?? '';

	let query = '';
	$: filteredFeeds = query
		? data.feeds.filter((f) => f.Name.toLowerCase().includes(query.toLowerCase()))
		: data.feeds;

	function selectFeed(id: number | undefined) {
		query = '';
		goto(`/?feed_id=${id ?? ''}`);
	}
</script>

<svelte:head>
	<title>Mates?</title>
</svelte:head>

<main class="h-full flex flex-col lg:flex-row overflow-hidden">
	<div class="flex-[2] lg:flex-1 lg:h-full">
		<LeafletMap />
	</div>
	<div class="flex-[3] lg:flex-1 lg:h-full overflow-scroll flex flex-col justify-start">
		<div class="p-2 border-b">
			<input
				type="search"
				placeholder="Search feeds…"
				bind:value={query}
				class="w-full border rounded px-3 py-1.5 text-sm outline-none focus:ring-2 focus:ring-gray-300"
			/>
		</div>
		{#if query}
			<ul class="border-b text-sm max-h-48 overflow-y-auto">
				{#each filteredFeeds as feed}
					<li>
						<button
							class="w-full text-left px-3 py-2 hover:bg-gray-100"
							on:click={() => selectFeed(feed.ID)}
						>
							{feed.Name}
							{#if feed.Desc.Valid}
								<span class="text-gray-400 ml-1">— {feed.Desc.String}</span>
							{/if}
						</button>
					</li>
				{:else}
					<li class="px-3 py-2 text-gray-400">No feeds found</li>
				{/each}
			</ul>
		{/if}
		{#if currentFeedId}
			{@const feed = data.feeds.find((f) => String(f.ID) === currentFeedId)}
			{#if feed}
				<div class="px-3 py-1 text-md text-gray-500 border-b flex justify-between">
					<div>
						Showing: <strong>{feed.Name}</strong>
					</div>
					<button on:click={() => selectFeed(undefined)}>Close</button>
				</div>
			{/if}
		{/if}
		<ReportList />
	</div>
</main>
