<script>
	import Button from '$lib/components/ui/button/button.svelte';
</script>

<header class="flex justify-between items-center p-3 bg-mate-grey">
	<div class="h-8">
		<a
			href="/"
			class="flex gap-2 h-full items-center uppercase font-light text-white hover:text-mate-lime transition duration-75"
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="stroke-mate-lime h-full"
				fill="none"
				viewBox="0 0 24 24"
				stroke-width="1.5"
				stroke="currentColor"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"
				/>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0Z"
				/>
			</svg>

			Where Are Our Mates?
		</a>
	</div>

	<nav></nav>

	<Button variant={'action'}>Report Mate Sighting</Button>
</header>
