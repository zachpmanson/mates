<script>
	import Button from '$lib/components/ui/button/button.svelte';
	import logo from '$lib/images/logo.svg';
</script>

<header class="flex justify-between items-center gap-1 p-3 bg-mate-grey">
	<div class="h-8">
		<a
			href="/"
			class="flex gap-2 h-full items-center uppercase font-light text-white hover:text-mate-lime transition duration-75"
		>
			<img src={logo} alt="Where are our mates?" class="h-full" />

			Where Are Our Mates?
		</a>
	</div>

	<nav></nav>

	<Button variant={'action'}>Report Mate Sighting</Button>
</header>
