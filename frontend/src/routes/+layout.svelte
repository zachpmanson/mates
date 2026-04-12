<script>
	import '../app.css';
	import Header from './Header.svelte';
</script>

<div class="app h-screen flex flex-col overflow-hidden">
	<Header />

	<main class="grow overflow-hidden">
		<slot />
	</main>
</div>
