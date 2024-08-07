<script>
	import Header from './Header.svelte';
	import '../app.css';

	export let data;
</script>

<div class="app max-h-screen">
	<Header {data} />

	<main>
		<slot />
	</main>
</div>
