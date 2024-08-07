<script>
	import Header from './Header.svelte';
	import '../app.css';
</script>

<div class="app max-h-screen">
	<Header />

	<main>
		<slot />
	</main>
</div>
