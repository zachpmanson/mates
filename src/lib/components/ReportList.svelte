<script lang="ts">
	import ReportCard from './ReportCard.svelte';
</script>

<div class="h-full flex flex-col gap-2 p-2">
	<ReportCard
		title="Mates on the number 5 heading towards the city."
		timestamp={new Date('2024-08-06T20:00:00.000Z')}
		author="John Doe"
	/>
	<ReportCard
		title="Mates waiting for the 1 or 6 city bound on Lygon st Brunswick east"
		timestamp={new Date('2024-08-06T07:30:00.000Z')}
		author="Holly McGrath"
	/>
	<ReportCard
		title="Mates all over the place on Sydney rd tram in Brunswick today"
		timestamp={new Date('2024-08-06T08:00:00.000Z')}
		author="Jane Doe"
	/>
</div>
