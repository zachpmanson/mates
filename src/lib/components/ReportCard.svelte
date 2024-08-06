<script lang="ts">
	import { timeAgo } from '$lib/util/timesince';
	import Card from './ui/card/card.svelte';

	export let title: string;
	export let timestamp: Date;
	export let author: string;
</script>

<Card>
	<div class="flex flex-col gap-2 p-4">
		<div class="flex items-center gap-2">
			<span class="">{title}</span>
		</div>
		<div class="flex items-center gap-3">
			<span class="text-gray-400">
				{timeAgo(timestamp)} ago
			</span>

			<span class="text-gray-400 text-sm">
				reported by {author}
			</span>
		</div>
	</div>
</Card>
