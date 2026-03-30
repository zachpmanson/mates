<script>
	import Header from './Header.svelte';
	import '../app.css';
</script>

<div class="app h-screen flex flex-col">
	<Header />

	<main class="grow">
		<slot />
	</main>
</div>
